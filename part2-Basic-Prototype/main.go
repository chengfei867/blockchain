package main

import (
	"blockchain/publicChain/part2-Basic-Prototype/BLC"
	"fmt"
)

func main() {
	block := BLC.NewBlock("Genenis Block!!", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(block)
}
