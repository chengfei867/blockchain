package BLC

import (
	"flag"
	"log"
	"os"
)

type CLI struct {
}

func (cli *CLI) Run() {
	addressListCmd := flag.NewFlagSet("addressList", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createWallet", flag.ExitOnError)
	createChainCmd := flag.NewFlagSet("createChain", flag.ExitOnError)
	sendBlockCmd := flag.NewFlagSet("send", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getBalance", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)

	flagCreateChainAddress := createChainCmd.String("address", "", "创建创世区块的地址")
	flagFrom := sendBlockCmd.String("from", "", "转账源地址。。。")
	flagTo := sendBlockCmd.String("to", "", "转账目的地址。。。")
	flagAmount := sendBlockCmd.String("amount", "", "转账金额。。。")
	flagAddress := getBalanceCmd.String("address", "", "用户地址")

	if len(os.Args) == 1 {
		cli.PrintUsage()
	}
	switch os.Args[1] {
	case "createWallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "addressList":
		err := addressListCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createChain":
		err := createChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getBalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.PrintUsage()
	}

	if createWalletCmd.Parsed() {
		cli.CreateWallet()
	}
	if addressListCmd.Parsed() {
		cli.AddressList()
	}
	if createChainCmd.Parsed() {
		cli.CreateChain(*flagCreateChainAddress)
	}
	if sendBlockCmd.Parsed() {
		cli.Send(*flagFrom, *flagTo, *flagAmount)
	}
	if getBalanceCmd.Parsed() {
		cli.GetBalance(*flagAddress)
	}
	if printChainCmd.Parsed() {
		cli.PrintChain()
	}
}
