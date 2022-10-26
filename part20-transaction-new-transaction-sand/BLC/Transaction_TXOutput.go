package BLC

type TXOutput struct {
	Value        int64
	ScriptPubKey string
}

// UnLockWithAddress 判断输出是否属于当前地址
func (txOutput *TXOutput) UnLockWithAddress(address string) bool {
	return txOutput.ScriptPubKey == address
}
