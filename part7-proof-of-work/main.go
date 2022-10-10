package main

import (
	"blockchain/publicChain/part7-proof-of-work/BLC"
	"fmt"
)

func main() {
	//创世区块
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	//新区快
	blockchain.AddBlockToBlockchain("Send 100BTC To ffg")
	blockchain.AddBlockToBlockchain("Send 20BTC To ffg")
	fmt.Printf("区块0的nonce：%d,hash：%x\n", blockchain.Blocks[0].Nonce, blockchain.Blocks[0].Hash)
	fmt.Printf("区块1的nonce：%d,hash：%x\n", blockchain.Blocks[1].Nonce, blockchain.Blocks[1].Hash)
	fmt.Printf("区块2的nonce：%d,hash：%x\n", blockchain.Blocks[2].Nonce, blockchain.Blocks[2].Hash)
}
