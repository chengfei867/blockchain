package main

import (
	"blockchain/publicChain/part3-Basic-Prototype/BLC"
	"fmt"
)

func main() {
	genesisBlockchain := BLC.CreateBlockchainWithGenesisBlock()
	fmt.Println(genesisBlockchain)
}
