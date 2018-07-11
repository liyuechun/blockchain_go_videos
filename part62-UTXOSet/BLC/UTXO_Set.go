package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

// 1.有一个方法，功能：

// 遍历整个数据库，读取所有的未花费的UTXO，然后将所有的UTXO存储到数据库
// reset
// 去遍历数据库时，
// []*TXOutputs

//[string]*TXOutputs
//
//txHash,TXOutputs := range txOutputsMap {
//
//
//
//
//}

const utxoTableName  = "utxoTableName"

type UTXOSet struct {
	Blockchain *Blockchain
}

// 重置数据库表
func (utxoSet *UTXOSet) ResetUTXOSet()  {

	err := utxoSet.Blockchain.DB.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(utxoTableName))

		if b != nil {
			tx.DeleteBucket([]byte(utxoTableName))

			b ,_ := tx.CreateBucket([]byte(utxoTableName))
			if b != nil {

				//[string]*TXOutputs
				//txOutputsMap := utxoSet.Blockchain.FindUTXOMap()


			}
		}


		return nil
	})

	if err != nil {
		log.Panic(err)
	}




}












