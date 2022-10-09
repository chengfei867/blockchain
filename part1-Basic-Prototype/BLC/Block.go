package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
}

//SetHash 设置区块hash
func (block *Block) SetHash() {
	//1.将Height转为[]byte
	heightBytes := IntToHex(block.Height)
	//2.时间戳转化为字节数组
	timeString := strconv.FormatInt(block.Timestamp, 2)
	timeBytes := []byte(timeString)
	//3.拼接所有的属性
	blockBytes := bytes.Join([][]byte{heightBytes, block.PrevBlockHash, block.Data, timeBytes}, []byte{})
	//4.生成Hash
	hash := sha256.Sum256(blockBytes)

	block.Hash = hash[:]
}

// NewBlock 创建新的区块
func NewBlock(data string, height int64, preBlockHash []byte) *Block {

	block := &Block{height, preBlockHash, []byte(data), time.Now().Unix(), nil}
	block.SetHash()
	return block
}
