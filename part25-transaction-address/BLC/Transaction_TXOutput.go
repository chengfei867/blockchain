package BLC

import "bytes"

type TXOutput struct {
	Value int64
	//收款人公钥
	Ripemd160Hash []byte
}

func (txOutput *TXOutput) Lock(address string) {

	publicKeyHash := Base58Decode([]byte(address))

	txOutput.Ripemd160Hash = publicKeyHash[1 : len(publicKeyHash)-4]
}

// NewTXOutput +根据address创建TXOutput
func NewTXOutput(value int64, address string) *TXOutput {

	txOutput := &TXOutput{value, nil}

	// 设置Ripemd160Hash
	txOutput.Lock(address)

	return txOutput
}

// UnLockScriptPubKeyWithAddress 解锁
func (txOutput *TXOutput) UnLockScriptPubKeyWithAddress(address string) bool {

	publicKeyHash := Base58Decode([]byte(address))
	hash160 := publicKeyHash[1 : len(publicKeyHash)-4]

	return bytes.Compare(txOutput.Ripemd160Hash, hash160) == 0
}
