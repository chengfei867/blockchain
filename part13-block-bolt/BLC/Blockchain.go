package BLC

type Blockchain struct {
	Blocks []*Block //存储有序的区块
}

// AddBlockToBlockchain 向区块链中添加区块
func (blc *Blockchain) AddBlockToBlockchain(data string) {
	//创建新区快
	newBlock := NewBlock(data, blc.Blocks[len(blc.Blocks)-1].Height+1, blc.Blocks[len(blc.Blocks)-1].Hash)
	//添加区块到区块链
	blc.Blocks = append(blc.Blocks, newBlock)
}

// CreateBlockchainWithGenesisBlock 创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {
	genesis := CreateGenesisBlock("创世区块!!!")
	return &Blockchain{[]*Block{genesis}}
}
