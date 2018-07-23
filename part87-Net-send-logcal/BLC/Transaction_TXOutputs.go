package BLC

import (
	"log"
	"bytes"
	"encoding/gob"
)

type TXOutputs struct {
	UTXOS []*UTXO
}

// 将区块序列化成字节数组
func (txOutputs *TXOutputs) Serialize() []byte {

	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(txOutputs)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
	//fmt.Println("$$$$$$$$$$")
	//jsonByte, err := json.Marshal(txOutputs)
	//if err != nil {
	//	//fmt.Println("序列化失败:",err)
	//	log.Panic(err)
	//}
	//return jsonByte
}

// 反序列化
func DeserializeTXOutputs(txOutputsBytes []byte) *TXOutputs {

	var txOutputs TXOutputs

	decoder := gob.NewDecoder(bytes.NewReader(txOutputsBytes))
	err := decoder.Decode(&txOutputs)
	if err != nil {
		log.Panic(err)
	}

	return &txOutputs
	//json.Unmarshal(txOutputsBytes, &txOutputs)
	//return &txOutputs
}
