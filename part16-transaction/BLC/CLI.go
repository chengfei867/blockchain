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
	fmt.Println("\taddBlock -data DATA -交易数据")
	fmt.Println("\tprintChain  --输出区块信息")
	os.Exit(1)
}

func (cli *CLI) Run() {
	createChainCmd := flag.NewFlagSet("createChain", flag.ExitOnError)
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	flagCreateChainAddress := createChainCmd.String("address", "", "创建创世区块的地址")
	flagAddBlockData := addBlockCmd.String("addBlock", "ffg@xd.com", "添加区块")
	if len(os.Args) == 1 {
		cli.PrintUsage()
	}
	switch os.Args[1] {
	case "createChain":
		err := createChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
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

	if addBlockCmd.Parsed() {
		if !dbExists() {
			fmt.Println("区块链不存在!")
			os.Exit(1)
		}
		blockchain := GetBlockChain()
		defer blockchain.Close()
		if *flagAddBlockData == "" {
			cli.PrintUsage()
		}
		blockchain.AddBlockToBlockchain([]*Transaction{})
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
