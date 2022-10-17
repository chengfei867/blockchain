package BLC

type TXInput struct {
	//1.交易ID
	Txid []byte
	//2.存储TXOutput在Vout里面的索引
	Vout int
	//3.用户名
	ScriptSig string
}
