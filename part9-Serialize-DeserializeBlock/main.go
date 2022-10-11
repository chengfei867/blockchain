package main

import (
	"blockchain/publicChain/part9-Serialize-DeserializeBlock/BLC"
	"fmt"
)

func main() {
	//创世区块
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	//新区快
	fmt.Printf("区块1的nonce：%d,hash：%x\n", blockchain.Blocks[0].Nonce, blockchain.Blocks[0].Hash)

	bytes, _ := blockchain.Blocks[0].Serialize()
	block := BLC.Deserialize(bytes)
	fmt.Printf("区块1的nonce：%d,hash：%x\n", block.Nonce, block.Hash)
}
