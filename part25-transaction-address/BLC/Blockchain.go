package BLC

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
	"strconv"
	"time"
)

const dbName = "blockchain.db"
const bucketName = "blockchain"
const tipString = "Tip"

type Blockchain struct {
	Tip []byte //最新区块的哈希值
	DB  *bolt.DB
}

// 判断数据库是否存在
func dbExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}

// CreateBlockchainWithGenesisBlock 创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock(address string) {
	if dbExists() {
		fmt.Println("区块链已存在......")
		blockchain := GetBlockChain()
		defer blockchain.Close()
		blockchain.PrintChain()
		os.Exit(1)
	}
	//创建区块链数据库
	DB, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Panicln(err)
	}
	defer func(DB *bolt.DB) {
		err := DB.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(DB)
	//初始化区块链
	var blockchain = new(Blockchain)
	//创建创世区块
	genesis := CreateGenesisBlock([]*Transaction{NewCoinbaseTransaction(address)})
	genesisBytes, err := genesis.Serialize()
	if err != nil {
		log.Panicln(err)
	}
	//将创世区块加入区块链
	err = DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte(bucketName))
		if err != nil {
			log.Panicln(err)
		}
		err = bucket.Put([]byte("Tip"), genesis.Hash)
		if err != nil {
			log.Panicln(err)
		}
		err = bucket.Put(genesis.Hash, genesisBytes)
		if err != nil {
			log.Panicln(err)
		}
		return nil
	})
	//更新区块结构
	blockchain.Tip = genesis.Hash
	blockchain.DB = DB
}

// GetUtxos 返回address所有未花费交易
func (blc *Blockchain) GetUtxos(address string, txs []*Transaction) []*UTXO {
	var utxo []*UTXO
	spentTXOutputs := make(map[string][]int)
	//先遍历txs
	for _, tx := range txs {
		if !tx.IsCoinbaseTransaction() {
			for _, in := range tx.Vins {
				publicKeyHash := Base58Decode([]byte(address))
				ripemd160Hash := publicKeyHash[1 : len(publicKeyHash)-4]
				//是否能解锁
				if in.UnLockRipemd160Hash(ripemd160Hash) {
					key := hex.EncodeToString(in.TxHash)
					spentTXOutputs[key] = append(spentTXOutputs[key], in.Vout)
				}
			}
		}
	}
	for _, tx := range txs {
	work1:
		for index, out := range tx.Vouts {
			if out.UnLockScriptPubKeyWithAddress(address) {
				if len(spentTXOutputs) == 0 {
					utxo = append(utxo, &UTXO{tx.TxHash, index, out})
				} else {
					for hash, indexArray := range spentTXOutputs {
						if hash == hex.EncodeToString(tx.TxHash) {
							var isUnSpentUTXO bool
							for _, outIndex := range indexArray {
								if outIndex == index {
									isUnSpentUTXO = true
									continue work1
								}
								if isUnSpentUTXO == false {
									utxo = append(utxo, &UTXO{tx.TxHash, index, out})
								}
							}
						} else {
							utxo = append(utxo, &UTXO{tx.TxHash, index, out})
						}
					}
				}
			}
		}
	}
	//最后一个区块hash
	hash := blc.Tip
	for hash != nil {
		//根据哈希查找区块
		block := blc.GetBlock(hash)
		//遍历交易
		for i := len(block.Txs) - 1; i >= 0; i-- {
			tx := block.Txs[i]
			//输入
			if !tx.IsCoinbaseTransaction() {
				for _, in := range tx.Vins {
					//是否能解锁
					publicKeyHash := Base58Decode([]byte(address))
					ripemd160Hash := publicKeyHash[1 : len(publicKeyHash)-4]
					if in.UnLockRipemd160Hash(ripemd160Hash) {
						key := hex.EncodeToString(in.TxHash)
						spentTXOutputs[key] = append(spentTXOutputs[key], in.Vout)
					}
				}
			}
			//输出
		to:
			for index, out := range tx.Vouts {
				//是否能解锁
				if out.UnLockScriptPubKeyWithAddress(address) {
					if spentTXOutputs != nil {
						//遍历spentTXOutputs判断当前输出是否在之后的区块已被花费
						for txHash, indexArray := range spentTXOutputs {
							if txHash == hex.EncodeToString(tx.TxHash) {
								for _, i := range indexArray {
									if index == i {
										continue to
									}
								}
							}
						}
					}
					utxo = append(utxo, &UTXO{tx.TxHash, index, out})
				}

			}
		}
		//hash指向前一个区块
		hash = block.PrevBlockHash
	}
	return utxo
}

// GetAbleUTXO 得到可用utxo
func (blc *Blockchain) GetAbleUTXO(from string, amount int64, txs []*Transaction) (int64, map[string][]int) {
	ableUTXO := make(map[string][]int)
	//获取所有的UTXO
	utxos := blc.GetUtxos(from, txs)
	//遍历utxo
	money := int64(0)
	for _, utxo := range utxos {
		hashString := hex.EncodeToString(utxo.TxHash)
		ableUTXO[hashString] = append(ableUTXO[hashString], utxo.Index)
		money += utxo.Output.Value
		if money >= amount {
			break
		}
	}
	if money < amount {
		fmt.Printf("%s账户余额不足!", from)
		os.Exit(1)
	}
	return money, ableUTXO
}

// GetBalance 获取账户余额
func (blc *Blockchain) GetBalance(address string) {
	utxo := blc.GetUtxos(address, []*Transaction{})
	var amount int64
	for _, output := range utxo {
		amount += output.Output.Value
	}
	fmt.Printf("%s的余额是：%dBTC\n", address, amount)
}

// MineNewBlock 打包交易形成新区快
func (blc *Blockchain) MineNewBlock(from []string, to []string, amount []string) {
	//1.构建交易数组
	var txs []*Transaction
	for index, _ := range from {
		value, _ := strconv.Atoi(amount[index])
		tx := NewTransaction(from[index], to[index], int64(value), blc, txs)
		txs = append(txs, tx)
	}
	//2.创建新区快
	block := NewBlock(txs, blc.GetHeight()+1, blc.Tip)
	//3.将区块存储到数据库
	blc.AddBlockToBlockchain(block)
}

//GetBlockChain 通过数据库获取blockchain
func GetBlockChain() *Blockchain {
	//创建区块链数据库
	DB, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Panicln(err)
	}
	var Tip []byte
	err = DB.View(func(tx *bolt.Tx) error {
		//获取区块链表
		bucket := tx.Bucket([]byte(bucketName))
		if bucket != nil {
			//获取最后一个区块哈希
			Tip = bucket.Get([]byte(tipString))
		}
		return nil
	})
	if err != nil {
		log.Panicln(err)
	}
	return &Blockchain{
		Tip: Tip,
		DB:  DB,
	}
}

// AddBlockToBlockchain 添加区块到区块链
func (blc *Blockchain) AddBlockToBlockchain(block *Block) {
	//获取当前区块高度
	err := blc.DB.Update(func(tx *bolt.Tx) error {
		//定义全局异常变量
		var err error
		//打开“blockchain”bucket 若为nil则创建
		bucket := tx.Bucket([]byte(bucketName))

		//将区块哈希作为最新区块哈希插入Tip
		err = bucket.Put([]byte("Tip"), block.Hash)
		if err != nil {
			log.Panicln(err)
		}
		//将区块序列化
		blockBytes, err := block.Serialize()
		if err != nil {
			log.Panicln(err)
		}
		//以区块哈希作为键 区块字节数组作为值存入数据库
		err = bucket.Put(block.Hash, blockBytes)
		if err != nil {
			log.Panicln(err)
		}
		return nil
	})
	if err != nil {
		log.Panicln(err)
	}
	//更新区块链的Tip
	blc.Tip = block.Hash
}

// GetBlock 读取区块
func (blc *Blockchain) GetBlock(hash []byte) *Block {
	db := blc.DB
	block := new(Block)
	err := db.View(func(tx *bolt.Tx) error {
		blockchain := tx.Bucket([]byte(bucketName))
		bytes := blockchain.Get(hash)
		block = Deserialize(bytes)
		return nil
	})
	if err != nil {
		log.Panicln(err.Error())
	}
	return block
}

// GetHeight 获取区块高度
func (blc *Blockchain) GetHeight() int64 {
	return blc.GetBlock(blc.Tip).Height
}

// PrintChain  遍历区块链
func (blc *Blockchain) PrintChain() {
	hash := blc.Tip
	for hash != nil {
		block := blc.GetBlock(hash)
		fmt.Println("===============================================")
		fmt.Printf("Height：%d\n", block.Height)
		fmt.Printf("PrevBlockHash：%x\n", block.PrevBlockHash)
		fmt.Printf("Timestamp：%s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash：%x\n", block.Hash)
		fmt.Printf("Nonce：%d\n", block.Nonce)
		fmt.Println("Transactions:")
		for _, tx := range block.Txs {
			fmt.Printf("%x\n", tx.TxHash)
			fmt.Println("Vins:")
			for _, in := range tx.Vins {
				fmt.Printf("\tTxHash:%x\n", in.TxHash)
				fmt.Printf("\tindex:%d\n", in.Vout)
				fmt.Printf("\tSig:%s\n", in.Signature)
				fmt.Printf("\tPubKey:%s\n", in.PublicKey)
			}

			fmt.Println("Vouts:")
			for _, out := range tx.Vouts {
				fmt.Printf("\tValue:%d\n", out.Value)
				fmt.Printf("\tScriptPubKey:%s\n", out.Ripemd160Hash)
			}
		}
		fmt.Println("===============================================")
		hash = block.PrevBlockHash
	}
}

//Close 关闭区块链数据库连接
func (blc *Blockchain) Close() {
	err := blc.DB.Close()
	if err != nil {
		log.Panicln(err)
	}
}

//19tp59gd68fitp9E8APBYf3JgUaHbzCvbR
