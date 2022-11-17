package BLC

import "bytes"

type TXInput struct {
	//交易ID
	TxHash []byte
	//存储TXOutput在Vout里面的索引
	Vout int
	//数字签名
	Signature []byte
	//公钥
	PublicKey []byte
}

// UnLockRipemd160Hash 判断当前的消费是谁的钱
func (txInput *TXInput) UnLockRipemd160Hash(ripemd160Hash []byte) bool {

	publicKey := Ripemd160Hash(txInput.PublicKey)

	return bytes.Compare(publicKey, ripemd160Hash) == 0
}
