package BLC

import "math/big"

const targetBit = 16

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork 创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {
	//1、创建一个初始值为1的target
	target := big.NewInt(1)
	//2、左移256-targetBit
	target = target.Lsh(target, 256-targetBit)
	return &ProofOfWork{
		block:  block,
		target: target,
	}
}

// Run 工作量证明方法
func (pow ProofOfWork) Run() ([]byte, int64) {
	return nil, 0
}
