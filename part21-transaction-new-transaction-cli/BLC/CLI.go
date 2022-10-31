package BLC

import (
	"flag"
	"log"
	"os"
)

type CLI struct {
}

func (cli *CLI) Run() {
	createChainCmd := flag.NewFlagSet("createChain", flag.ExitOnError)
	sendBlockCmd := flag.NewFlagSet("send", flag.ExitOnError)
	getBlanceCmd := flag.NewFlagSet("getBalance", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)

	flagCreateChainAddress := createChainCmd.String("address", "", "创建创世区块的地址")
	flagFrom := sendBlockCmd.String("from", "", "转账源地址。。。")
	flagTo := sendBlockCmd.String("to", "", "转账目的地址。。。")
	flagAmount := sendBlockCmd.String("amount", "", "转账金额。。。")
	flagAddress := getBlanceCmd.String("address", "", "用户地址")

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
	case "getBalance":
		err := getBlanceCmd.Parse(os.Args[2:])
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
		cli.CreateChain(*flagCreateChainAddress)
	}
	if sendBlockCmd.Parsed() {
		cli.Send(*flagFrom, *flagTo, *flagAmount)
	}
	if getBlanceCmd.Parsed() {
		cli.GetBalance(*flagAddress)
	}
	if printChainCmd.Parsed() {
		cli.PrintChain()
	}
}
