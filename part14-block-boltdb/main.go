package main

import (
	"kongyixueyuan.com/publicChain/part14-block-boltdb/BLC"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main()  {

	////data string,height int64,prevBlockHash []byte
	//block := BLC.NewBlock("Test",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	//fmt.Printf("%d\n",block.Nonce)
	//fmt.Printf("%x\n",block.Hash)


	// 创建或者打开数据库
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	//// 更新数据库
	//err = db.Update(func(tx *bolt.Tx) error {
	//
	//	// 取表对象
	//	b := tx.Bucket([]byte("blocks"))
	//
	//	if b == nil {
	//		b,err = tx.CreateBucket([]byte("blocks"))
	//		if err != nil {
	//			log.Panic("Blocks table create failed......")
	//		}
	//	}
	//
	//	err = b.Put([]byte("l"),block.Serialize())
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//
	//	return nil
	//})
	//
	//if  err != nil{
	//	log.Panic(err)
	//}


	err = db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("blocks"))
		if b != nil {
			blockData := b.Get([]byte("l"))
			//fmt.Println(blockData)
			//fmt.Printf("%s\n",blockData)
			block := BLC.DeserializeBlock(blockData)
			fmt.Printf("%v\n",block)
		}
		return nil
	})

	if  err != nil{
		log.Panic(err)
	}

}
