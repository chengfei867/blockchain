package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
}

func (cli *CLI) PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreateChain -address Address  --创建区块链")
	fmt.Println("\tsend -from FROM -to TO -amount AMOUNT -转账交易")
	fmt.Println("\tprintChain  --输出区块信息")
	os.Exit(1)
}

func (cli *CLI) Run() {
	createChainCmd := flag.NewFlagSet("createChain", flag.ExitOnError)
	sendBlockCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)

	flagCreateChainAddress := createChainCmd.String("address", "", "创建创世区块的地址")
	flagFrom := sendBlockCmd.String("from", "", "转账源地址。。。")
	flagTo := sendBlockCmd.String("to", "", "转账目的地址。。。")
	flagAmount := sendBlockCmd.String("amount", "", "转账金额。。。")

	if len(os.Args) == 1 {
		cli.PrintUsage()
	}
	switch os.Args[1] {
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
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.PrintUsage()
	}

	if createChainCmd.Parsed() {
		if *flagCreateChainAddress == "" {
			cli.PrintUsage()
		}
		CreateBlockchainWithGenesisBlock(*flagCreateChainAddress)
	}

	if sendBlockCmd.Parsed() {
		if !dbExists() {
			fmt.Println("区块链不存在!")
			os.Exit(1)
		}
		blockchain := GetBlockChain()
		defer blockchain.Close()
		if *flagFrom == "" {
			log.Panicln("转账源地址不能为空!")
			cli.PrintUsage()
		}
		if *flagTo == "" {
			log.Panicln("转账目的地址不能为空!")
			cli.PrintUsage()
		}
		if *flagAmount == "" {
			log.Panicln("转账金额不能为空!")
			cli.PrintUsage()
		}
		//blockchain.AddBlockToBlockchain([]*Transaction{})
		fmt.Println(*flagFrom)
		fmt.Println(*flagTo)
		fmt.Println(*flagAmount)
	}
	if printChainCmd.Parsed() {
		if !dbExists() {
			fmt.Println("区块链不存在!")
			os.Exit(1)
		}
		blockchain := GetBlockChain()
		defer blockchain.Close()
		blockchain.PrintChain()
	}
}
