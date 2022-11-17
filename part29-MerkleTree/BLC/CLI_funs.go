package BLC

import (
	"fmt"
	"os"
)

// PrintUsage 使用方法
func (cli *CLI) PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreateWallet --创建钱包")
	fmt.Println("\taddressList --输出所有钱包地址")
	fmt.Println("\tcreateChain -address Address  --创建区块链")
	fmt.Println("\tsend -from FROM -to TO -amount AMOUNT -转账交易")
	fmt.Println("\tgetBalance -address Address  --获取账户余额")
	fmt.Println("\tprintChain  --输出区块信息")
	fmt.Println("\ttest -- 测试.")
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
	utxoSet := &UTXOSet{blockchain}
	amount := utxoSet.GetBalance(address)
	fmt.Printf("%s一共有%d个Token\n", address, amount)
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
	utxoSet := &UTXOSet{blockchain}
	utxoSet.Update()
}

// CreateChain 创建创世区块
func (cli *CLI) CreateChain(address string) {
	if address == "" {
		cli.PrintUsage()
	}
	blockchain := CreateBlockchainWithGenesisBlock(address)
	defer blockchain.Close()
	utxoSet := &UTXOSet{blockchain}
	utxoSet.ResetUTXOSet()
}

// CreateWallet 创建钱包
func (cli *CLI) CreateWallet() {
	wallets, _ := NewWallets()

	wallets.CreateNewWallet()

	fmt.Println(len(wallets.WalletsMap))
}

// AddressList 打印所有钱包地址
func (cli *CLI) AddressList() {
	fmt.Println("打印所有的钱包地址:")

	wallets, _ := NewWallets()

	for address, _ := range wallets.WalletsMap {
		fmt.Println(address)
	}
}

func (cli *CLI) TestMethod() {
	blockchain := GetBlockChain()
	defer blockchain.Close()
	utxoSet := &UTXOSet{blockchain}
	utxoSet.ResetUTXOSet()
	//fmt.Println(len(blockchain.FindUTXOMap()))
}
