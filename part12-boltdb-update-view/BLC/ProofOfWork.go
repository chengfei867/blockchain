package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const targetBit = 16

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProofOfWork 创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {
	//1、创建一个初始值为1的target
	target := big.NewInt(1)
	//2、左移256-targetBit
	target = target.Lsh(target, 256-targetBit)
	return &ProofOfWork{
		block,
		target,
	}
}

func (pow *ProofOfWork) isValid() bool {
	//1.proofOfWork.Block.Hash
	//2.proofOfWork.Target
	var hashInt *big.Int
	hashInt.SetBytes(pow.Block.Hash)
	//若hash>target
	if pow.Target.Cmp(hashInt) == -1 {
		return false
	}
	//hash<target
	return true
}

// prepareData 数据拼接，返回字节数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
			IntToHex(pow.Block.Height),
		},
		[]byte{},
	)

	return data
}

// Run 工作量证明方法
func (pow *ProofOfWork) Run() ([]byte, int64) {
	//1.将block的属性破解成字节数组
	//2.生成hash
	//3.判断hash的有效性 满足调节跳出循环
	hashInt := big.NewInt(0)
	nonce := 0
	for {
		//尊卑数据
		dataBytes := pow.prepareData(nonce)
		//计算哈希
		hash := sha256.Sum256(dataBytes)
		fmt.Printf("hash:%x\n", hash[:])
		//将哈希转换为big.Int类型
		hashInt.SetBytes(hash[:])
		if pow.Target.Cmp(hashInt) == 1 {
			fmt.Printf("Find nonce:%d,hash:%x\n", nonce, hash[:])
			return hash[:], int64(nonce)
		}
		nonce++
	}

}
