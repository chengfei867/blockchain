package BLC

import (
	"fmt"
	"os"
)

// PrintUsage 使用方法
func (cli *CLI) PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreateChain -address Address  --创建区块链")
	fmt.Println("\tsend -from FROM -to TO -amount AMOUNT -转账交易")
	fmt.Println("\tgetBalance -address Address  --获取账户余额")
	fmt.Println("\tprintChain  --输出区块信息")
	os.Exit(1)
}

// PrintChain 打印区块
func (cli *CLI) PrintChain() {
	if !dbExists() {
		fmt.Println("区块链不存在!")
		os.Exit(1)
	}
	blockchain := GetBlockChain()
	defer blockchain.Close()
	blockchain.PrintChain()
}

// GetBalance 获取余额
func (cli *CLI) GetBalance(address string) {
	if address == "" {
		fmt.Println("用户地址不能为空!")
		cli.PrintUsage()
	}
	blockchain := GetBlockChain()
	defer blockchain.Close()
	blockchain.GetBalance(address)
}

// Send 转账
func (cli *CLI) Send(fromStr string, toStr string, amountStr string) {
	if !dbExists() {
		fmt.Println("区块链不存在!")
		os.Exit(1)
	}
	if fromStr == "" {
		fmt.Println("转账源地址不能为空!")
		cli.PrintUsage()
	}
	if toStr == "" {
		fmt.Println("转账目的地址不能为空!")
		cli.PrintUsage()
	}
	if amountStr == "" {
		fmt.Println("转账金额不能为空!")
		cli.PrintUsage()
	}
	from := JsonToArray(fromStr)
	to := JsonToArray(toStr)
	amount := JsonToArray(amountStr)
	blockchain := GetBlockChain()
	defer blockchain.Close()
	blockchain.MineNewBlock(from, to, amount)
}

// CreateChain 创建创世区块
func (cli *CLI) CreateChain(address string) {
	if address == "" {
		cli.PrintUsage()
	}
	CreateBlockchainWithGenesisBlock(address)
}
