package BLC

type TXInput struct {
	// 1. 交易的Hash
	TxHash      []byte
	// 2. 存储TXOutput在Vout里面的索引
	Vout      int
	// 3. 用户名
	ScriptSig string
}



// 判断当前的消费是谁的钱
func (txInput *TXInput) UnLockWithAddress(address string) bool {

	return txInput.ScriptSig == address
}