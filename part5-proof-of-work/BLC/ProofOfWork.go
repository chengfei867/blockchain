package BLC

type ProofOfWork struct {
	block *Block
}

// NewProofOfWork 创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {
	return &ProofOfWork{
		block: block,
	}
}

// Run 工作量证明方法
func (pow ProofOfWork) Run() ([]byte, int64) {
	return nil, 0
}
