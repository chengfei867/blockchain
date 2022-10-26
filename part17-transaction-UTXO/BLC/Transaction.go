package BLC

import (
	"crypto/sha256"
	"encoding/json"
	"log"
)

type Transaction struct {
	//1.交易Hash
	TxHash []byte
	//2.输入
	Vins []*TXInput
	//3.输出
	Vouts []*TXOutput
}

// NewCoinbaseTransaction 创世区块创建时的Transaction
func NewCoinbaseTransaction(address string) *Transaction {
	txInput := &TXInput{[]byte{}, -1, "Genesis Block"}
	txOutput := &TXOutput{10, address}
	txCoinbase := &Transaction{[]byte{}, []*TXInput{txInput}, []*TXOutput{txOutput}}
	txCoinbase.SetHash()
	return txCoinbase
}

// SetHash  设置交易哈希
func (tx *Transaction) SetHash() {
	txBytes, err := json.Marshal(tx)
	if err != nil {
		log.Panic(err)
	}
	txHash := sha256.Sum256(txBytes)
	tx.TxHash = txHash[:]
}
