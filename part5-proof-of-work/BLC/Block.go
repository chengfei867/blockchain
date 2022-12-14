package BLC

import (
	"time"
)

type Block struct {
	//1. 区块高度
	Height int64
	//2. 上一个区块HASH
	PrevBlockHash []byte
	//3. 交易数据
	Data []byte
	//4. 时间戳
	Timestamp int64
	//5. Hash
	Hash []byte
	//6、Nonce
	Nonce int64
}

// NewBlock 创建新的区块
func NewBlock(data string, height int64, preBlockHash []byte) *Block {
	//创建区块
	block := &Block{height, preBlockHash, []byte(data), time.Now().Unix(), nil, 0}

	//调用工作量证明的方法并且返回有效的Hash和Nonce
	pow := NewProofOfWork(block)

	//000000
	hash, nonce := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// CreateGenesisBlock 生成创世区块
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
