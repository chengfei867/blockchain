package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
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
func CreateBlockchainWithGenesisBlock(tx []*Transaction) {
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
	genesis := CreateGenesisBlock(tx)
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
func (blc *Blockchain) AddBlockToBlockchain(txs []*Transaction) {
	//获取当前区块高度
	height := blc.GetBlock(blc.Tip).Height + 1
	block := NewBlock(txs, height, blc.Tip)
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

// PrintChain  遍历区块链
func (blc *Blockchain) PrintChain() {
	hash := blc.Tip
	for hash != nil {
		block := blc.GetBlock(hash)
		fmt.Println("===============================================")
		fmt.Printf("Height：%d\n", block.Height)
		fmt.Printf("PrevBlockHash：%x\n", block.PrevBlockHash)
		fmt.Printf("Transactions：%v\n", block.Txs)
		fmt.Printf("Timestamp：%s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash：%x\n", block.Hash)
		fmt.Printf("Nonce：%d\n", block.Nonce)
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
