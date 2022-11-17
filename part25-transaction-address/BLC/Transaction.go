package BLC

import (
	"crypto/sha256"
	"encoding/hex"
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
	txInput := &TXInput{[]byte{}, -1, nil, []byte{}}
	txOutput := NewTXOutput(10, address)
	txCoinbase := &Transaction{[]byte{}, []*TXInput{txInput}, []*TXOutput{txOutput}}
	txCoinbase.SetHash()
	return txCoinbase
}

// NewTransaction 转账时产生的Transaction
func NewTransaction(from string, to string, amount int64, blc *Blockchain, txs []*Transaction) *Transaction {
	var txInputs []*TXInput
	var txOutputs []*TXOutput
	wallets, err := NewWallets()
	if err != nil {
		log.Panicln(err)
	}
	wallet := wallets.WalletsMap[from]
	//判断该用户
	//获取该用户utxo
	money, ableUTXO := blc.GetAbleUTXO(from, amount, txs)
	//
	for txHash, indexArray := range ableUTXO {
		txHashBytes, err := hex.DecodeString(txHash)
		if err != nil {
			log.Panicln(err)
		}
		for _, index := range indexArray {
			txInput := &TXInput{txHashBytes, index, nil, wallet.PublicKey}
			txInputs = append(txInputs, txInput)
		}
	}
	//转账
	txOutput := NewTXOutput(amount, to)
	txOutputs = append(txOutputs, txOutput)
	//找零
	txOutput = NewTXOutput(money-amount, from)
	txOutputs = append(txOutputs, txOutput)

	tx := &Transaction{[]byte{}, txInputs, txOutputs}
	tx.SetHash()

	return tx
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
