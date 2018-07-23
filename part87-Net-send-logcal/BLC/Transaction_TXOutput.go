package BLC

import (
	"bytes"
	"fmt"
	"encoding/hex"
)


type TXOutput struct {
	Value int64
	Ripemd160Hash []byte  //用户名
}
func (txOutput *TXOutput)PrintInfo()  {
	fmt.Printf("Value:%s\n",txOutput.Value)
	fmt.Printf("Ripemd160Hash:%s\n",hex.EncodeToString(txOutput.Ripemd160Hash))

}
func (txOutput *TXOutput)  Lock(address string)  {

	publicKeyHash := Base58Decode([]byte(address))

	txOutput.Ripemd160Hash = publicKeyHash[1:len(publicKeyHash) - 4]
}


func NewTXOutput(value int64,address string) *TXOutput {

	txOutput := &TXOutput{value,nil}

	// 设置Ripemd160Hash
	txOutput.Lock(address)

	return txOutput
}


// 解锁
func (txOutput *TXOutput) UnLockScriptPubKeyWithAddress(address string) bool {

	publicKeyHash := Base58Decode([]byte(address))
	hash160 := publicKeyHash[1:len(publicKeyHash) - 4]

	return bytes.Compare(txOutput.Ripemd160Hash,hash160) == 0
}



