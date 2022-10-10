package BLC

type Blockchain struct {
	Blocks []*Block //存储有序的区块
}

// CreateBlockchainWithGenesisBlock 创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {
	genesis := CreateGenesisBlock("创世区块!!!")
	return &Blockchain{[]*Block{genesis}}
}
