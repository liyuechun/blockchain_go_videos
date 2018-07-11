package BLC

import (
	"github.com/boltdb/bolt"
	"log"
	"encoding/hex"
)



const utxoTableName  = "utxoTableName"

type UTXOSet struct {
	Blockchain *Blockchain
}

// 重置数据库表
func (utxoSet *UTXOSet) ResetUTXOSet()  {

	err := utxoSet.Blockchain.DB.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(utxoTableName))

		if b != nil {


			err := tx.DeleteBucket([]byte(utxoTableName))

			if err!= nil {
				log.Panic(err)
			}

		}

		b ,_ = tx.CreateBucket([]byte(utxoTableName))
		if b != nil {

			//[string]*TXOutputs
			txOutputsMap := utxoSet.Blockchain.FindUTXOMap()


			for keyHash,outs := range txOutputsMap {

				txHash,_ := hex.DecodeString(keyHash)

				b.Put(txHash,outs.Serialize())

			}
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

}












