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
	fmt.Println("\tcreateChain -data DATA  --创建区块链")
	fmt.Println("\taddBlock -data DATA -交易数据")
	fmt.Println("\tprintChain  --输出区块信息")
	os.Exit(1)
}

func (cli *CLI) Run() {
	createChainCmd := flag.NewFlagSet("createChain", flag.ExitOnError)
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	flagCreateChainData := createChainCmd.String("data", "Genesis Block!!!", "创世区块数据")
	flagAddBlockData := addBlockCmd.String("data", "ffg@xd.com", "交易数据...")

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
		if *flagCreateChainData == "" {
			cli.PrintUsage()
		}
		CreateBlockchainWithGenesisBlock([]*Transaction{})
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
