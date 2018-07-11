package BLC

import (
	"bytes"
	"log"
	"encoding/gob"
	"crypto/sha256"
)

// UTXO
type Transaction struct {

	//1. 交易hash
	TxHash []byte

	//2. 输入
	Vins []*TXInput

	//3. 输出
	Vouts []*TXOutput
}

//[]byte{}

// 判断当前的交易是否是Coinbase交易
func (tx *Transaction) IsCoinbaseTransaction() bool {

	return len(tx.Vins[0].TxHash) == 0 && tx.Vins[0].Vout == -1
}



//1. Transaction 创建分两种情况
//1. 创世区块创建时的Transaction
func NewCoinbaseTransaction(address string) *Transaction {

	//代表消费
	txInput := &TXInput{[]byte{},-1,"Genesis Data"}

	txOutput := &TXOutput{10,address}

	txCoinbase := &Transaction{[]byte{},[]*TXInput{txInput},[]*TXOutput{txOutput}}

	//设置hash值
	txCoinbase.HashTransaction()


	return txCoinbase
}

func (tx *Transaction) HashTransaction()  {

	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	hash := sha256.Sum256(result.Bytes())

	tx.TxHash = hash[:]
}



//2. 转账时产生的Transaction

func NewSimpleTransaction(from string,to string,amount int) *Transaction {

	//$ ./bc send -from '["juncheng"]' -to '["zhangqiang"]' -amount '["2"]'
	//	[juncheng]
	//	[zhangqiang]
	//	[2]

	//1. 有一个函数，返回from这个人所有的未花费交易输出所对应的Transaction

	 //unSpentTx := UnSpentTransationsWithAdress(from)
	 //
	 //fmt.Println(unSpentTx)


	//// 通过一个函数，返回
	////money,dic :=
	////
	////	{hash1:[0,2],hash2:[1,4]}
	//
	//var txIntputs []*TXInput
	//var txOutputs []*TXOutput
	//
	////代表消费
	//bytes ,_ := hex.DecodeString("1b5032e0cf4851f84dd89b9154912c082e28d5aa7f141625a0641c8a74f61802")
	//txInput := &TXInput{bytes,0,from}
	//
	////fmt.Printf("s:%s\n",s)
	//
	//// 消费
	//txIntputs = append(txIntputs,txInput)
	//
	//
	//// 转账
	//txOutput := &TXOutput{int64(amount),to}
	//txOutputs = append(txOutputs,txOutput)
	//
	//// 找零
	//txOutput = &TXOutput{4 - int64(amount),from}
	//txOutputs = append(txOutputs,txOutput)
	//
	//tx := &Transaction{[]byte{},txIntputs,txOutputs}
	//
	////设置hash值
	//tx.HashTransaction()



	return nil

}




