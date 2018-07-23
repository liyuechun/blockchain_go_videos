package BLC

import (
	"bytes"
	"fmt"
	"encoding/hex"
)

type TXInput struct {
	// 1. 交易的Hash
	TxHash      []byte
	// 2. 存储TXOutput在Vout里面的索引
	Vout      int

	Signature []byte // 数字签名

	PublicKey    []byte // 公钥，钱包里面
}



func (txInput *TXInput)PrintInfo()  {
	fmt.Printf("txHash:%s\n",hex.EncodeToString(txInput.TxHash))
	fmt.Printf("Vout:%d\n",txInput.Vout)
	fmt.Printf("Signature:%v\n",txInput.Signature)
	fmt.Printf("PublicKey:%v\n",txInput.PublicKey)
}
// 判断当前的消费是谁的钱
func (txInput *TXInput) UnLockRipemd160Hash(ripemd160Hash []byte) bool {

	publicKey := Ripemd160Hash(txInput.PublicKey)

	return bytes.Compare(publicKey,ripemd160Hash) == 0
}