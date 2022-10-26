package main

import (
	"blockchain/publicChain/part15-cli/BLC"
	"flag"
	"fmt"
	"log"
	"os"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tinit  --初始化区块链")
	fmt.Println("\taddBlock -data DATA -交易数据")
	fmt.Println("\tprintChain  --输出区块信息")
}

func main() {

	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	flagAddBlockData := addBlockCmd.String("data", "ffg@xd.com", "交易数据...")
	flagCreateChainData := initCmd.String("data", "ffg@xd.com", "创世区块数据...")

	switch os.Args[1] {
	case "init":
		err := initCmd.Parse(os.Args[2:])
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
		printUsage()
		os.Exit(1)
	}

	if initCmd.Parsed() {
		if *flagCreateChainData == "" {
			printUsage()
			os.Exit(1)
		}
		BLC.CreateBlockchainWithGenesisBlock(*flagCreateChainData)
	}

	if addBlockCmd.Parsed() {
		blockchain := BLC.GetBlockChain()
		defer blockchain.Close()
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}
		blockchain.AddBlockToBlockchain(*flagAddBlockData)
	}
	if printChainCmd.Parsed() {
		blockchain := BLC.GetBlockChain()
		defer blockchain.Close()
		blockchain.PrintChain()
	}
}
