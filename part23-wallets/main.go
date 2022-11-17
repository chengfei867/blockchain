package main

import (
	"blockchain/publicChain/part23-wallets/BLC"
	"fmt"
)

func main() {
	wallets := BLC.NewWallets()
	fmt.Println(wallets.Wallets)
	wallets.CreateNewWallet()
	wallets.CreateNewWallet()
	fmt.Println(wallets.Wallets)
}
