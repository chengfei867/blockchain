package main

import (
	"blockchain/publicChain/part22-wallet-address/BLC"
	"fmt"
)

func main() {

	wallet := BLC.NewWallet()

	address := wallet.GetAddress()

	fmt.Printf("address：%s\n", address)

	isValid := BLC.IsValidForAddress(address)

	fmt.Printf("%s 这个地址为 %v\n", address, isValid)

}
