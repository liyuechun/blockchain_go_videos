package BLC



//TXOutput{100,"zhangbozhi"}
//TXOutput{30,"xietingfeng"}
//TXOutput{40,"zhangbozhi"}


type TXOutput struct {
	Value int64
	ScriptPubKey string  //用户名
}

// 解锁
func (txOutput *TXOutput) UnLockScriptPubKeyWithAddress(address string) bool {

	return txOutput.ScriptPubKey == address
}


