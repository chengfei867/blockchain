package BLC

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Block struct {
	//1. 区块高度
	Height int64
	//2. 上一个区块HASH
	PrevBlockHash []byte
	//3. 交易数据
	Txs []*Transaction
	//4. 时间戳
	Timestamp int64
	//5. Hash
	Hash []byte
	//6、Nonce
	Nonce int64
}

// NewBlock 创建新的区块
func NewBlock(txs []*Transaction, height int64, preBlockHash []byte) *Block {
	//创建区块
	block := &Block{height, preBlockHash, txs, time.Now().Unix(), nil, 0}

	//调用工作量证明的方法并且返回有效的Hash和Nonce
	pow := NewProofOfWork(block)

	//000000
	hash, nonce := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// CreateGenesisBlock 生成创世区块
func CreateGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(txs, 1, nil)
}

//HashTransactions 将交易转换成字节数组
func (block *Block) HashTransactions() []byte {
	txsBytes, err := json.Marshal(block.Txs)
	if err != nil {
		log.Panic(err)
	}
	return txsBytes
}

// Serialize 序列化
func (block *Block) Serialize() ([]byte, error) {
	return json.Marshal(block)
}

// Deserialize 反序列化
func Deserialize(blockByte []byte) *Block {
	var block = new(Block)
	err := json.Unmarshal(blockByte, block)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return block
}

//19tp59gd68fitp9E8APBYf3JgUaHbzCvbR
//1FMCZiBKQDx8BowJN295WKffnJtR7qhAJL
