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

// IsCoinbaseTransaction 是否是否是创世区块交易
func (tx *Transaction) IsCoinbaseTransaction() bool {
	return len(tx.TxHash) == 0
}

// NewCoinbaseTransaction 创世区块创建时的Transaction
func NewCoinbaseTransaction(address string) *Transaction {
	txInput := &TXInput{[]byte{}, -1, "Genesis Block"}
	txOutput := &TXOutput{10, address}
	txCoinbase := &Transaction{[]byte{}, []*TXInput{txInput}, []*TXOutput{txOutput}}
	txCoinbase.SetHash()
	return txCoinbase
}

// NewTransaction 转账时产生的Transaction
func NewTransaction(from string, to string, amount int64) *Transaction {

	//unSpentTx := UnSpentTransactionsWithAddress(from)
	//fmt.Println(unSpentTx)

	//var txInputs []*TXInput
	//var txOutputs []*TXOutput
	//
	////消费
	//decodeString, _ := hex.DecodeString("c13d03716f4d51c355479b9969e5a0453a026177f22b1ca5e8baa43cc32d8ff2")
	//txInput := &TXInput{decodeString, 0, from}
	//txInputs = append(txInputs, txInput)
	//
	////转账
	//txOutput := &TXOutput{amount, to}
	//txOutputs = append(txOutputs, txOutput)
	//
	////找零
	//txOutput = &TXOutput{10 - amount, from}
	//txOutputs = append(txOutputs, txOutput)
	//
	//tx := &Transaction{[]byte{}, txInputs, txOutputs}
	//tx.SetHash()
	return nil
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
