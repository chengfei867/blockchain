package main

import (
	"blockchain/publicChain/part5-proof-of-work/BLC"
)

func main() {
	//创世区块
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	//新区快
	blockchain.AddBlockToBlockchain("Send 100BTC To ffg")
	blockchain.AddBlockToBlockchain("Send 200BTC To ffg")
	blockchain.AddBlockToBlockchain("Send 300BTC To ffg")
	blockchain.AddBlockToBlockchain("Send 400BTC To ffg")

}
