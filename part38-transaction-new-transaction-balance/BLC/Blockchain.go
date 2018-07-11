package BLC

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
	"math/big"
	"time"
	"os"
	"strconv"
	"encoding/hex"
)

// 数据库名字
const dbName = "blockchain.db"

// 表的名字
const blockTableName = "blocks"

type Blockchain struct {
	Tip []byte //最新的区块的Hash
	DB  *bolt.DB
}

// 迭代器
func (blockchain *Blockchain) Iterator() *BlockchainIterator {

	return &BlockchainIterator{blockchain.Tip, blockchain.DB}
}

// 判断数据库是否存在
func DBExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}

	return true
}

// 遍历输出所有区块的信息
func (blc *Blockchain) Printchain() {

	fmt.Println("PrintchainPrintchainPrintchainPrintchain")
	blockchainIterator := blc.Iterator()

	for {
		block := blockchainIterator.Next()

		fmt.Printf("Height：%d\n", block.Height)
		fmt.Printf("PrevBlockHash：%x\n", block.PrevBlockHash)
		fmt.Printf("Timestamp：%s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash：%x\n", block.Hash)
		fmt.Printf("Nonce：%d\n", block.Nonce)
		fmt.Println("Txs:")
		for _,tx := range block.Txs {

			fmt.Printf("%x\n",tx.TxHash)
			fmt.Println("Vins:")
			for _,in := range tx.Vins  {
				fmt.Printf("%x\n",in.TxHash)
				fmt.Printf("%d\n",in.Vout)
				fmt.Printf("%s\n",in.ScriptSig)
			}

			fmt.Println("Vouts:")
			for _,out := range tx.Vouts  {
				fmt.Println(out.Value)
				fmt.Println(out.ScriptPubKey)
			}
		}

		fmt.Println("------------------------------")

		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)

		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y

		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break;
		}
	}

}

//// 增加区块到区块链里面
func (blc *Blockchain) AddBlockToBlockchain(txs []*Transaction) {

	err := blc.DB.Update(func(tx *bolt.Tx) error {

		//1. 获取表
		b := tx.Bucket([]byte(blockTableName))
		//2. 创建新区块
		if b != nil {

			// ⚠️，先获取最新区块
			blockBytes := b.Get(blc.Tip)
			// 反序列化
			block := DeserializeBlock(blockBytes)

			//3. 将区块序列化并且存储到数据库中
			newBlock := NewBlock(txs, block.Height+1, block.Hash)
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//4. 更新数据库里面"l"对应的hash
			err = b.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			//5. 更新blockchain的Tip
			blc.Tip = newBlock.Hash
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}







//1. 创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock(address string) *Blockchain {

	// 判断数据库是否存在
	if DBExists() {
		fmt.Println("创世区块已经存在.......")
		os.Exit(1)
	}


	fmt.Println("正在创建创世区块.......")

	// 创建或者打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}


	var genesisHash []byte

	// 关闭数据库
	err = db.Update(func(tx *bolt.Tx) error {

		// 创建数据库表
		b, err := tx.CreateBucket([]byte(blockTableName))

		if err != nil {
			log.Panic(err)
		}

		if b != nil {
			// 创建创世区块
			// 创建了一个coinbase Transaction
			txCoinbase := NewCoinbaseTransaction(address)

			genesisBlock := CreateGenesisBlock([]*Transaction{txCoinbase})
			// 将创世区块存储到表中
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}

			// 存储最新的区块的hash
			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}

			genesisHash = genesisBlock.Hash
		}

		return nil
	})


	return &Blockchain{genesisHash,db}

}


// 返回Blockchain对象
func BlockchainObject() *Blockchain {

	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var tip []byte

	err = db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			// 读取最新区块的Hash
			tip = b.Get([]byte("l"))

		}


		return nil
	})

	return &Blockchain{tip,db}
}


// 如果一个地址对应的TXOutput未花费，那么这个Transaction就应该添加到数组中返回
func (blockchain *Blockchain) UnUTXOs(address string) []*TXOutput {

	var unUTXOs []*TXOutput

	spentTXOutputs := make(map[string][]int)


	//{hash:[0]}

	blockIterator := blockchain.Iterator()


	for  {

		block := blockIterator.Next()

		fmt.Println(block)
		fmt.Println()

		for _,tx := range block.Txs {

			// txHash
			// Vins
			if tx.IsCoinbaseTransaction() == false {
				for _,in := range tx.Vins {
					//是否能够解锁
					if in.UnLockWithAddress(address) {

						key := hex.EncodeToString(in.TxHash)

						spentTXOutputs[key] = append(spentTXOutputs[key],in.Vout)
					}

				}
			}


			// Vouts
			for index,out := range tx.Vouts {

				if out.UnLockScriptPubKeyWithAddress(address) {

					fmt.Println(out)
					fmt.Println(spentTXOutputs)

					//&{2 zhangqiang}
					//map[]

					if spentTXOutputs != nil {

						//map[cea12d33b2e7083221bf3401764fb661fd6c34fab50f5460e77628c42ca0e92b:[0]]

						if len(spentTXOutputs) != 0 {
							for txHash,indexArray := range spentTXOutputs {


								for _,i := range  indexArray {
									if index == i && txHash == hex.EncodeToString(tx.TxHash){
										continue
									} else {
										unUTXOs = append(unUTXOs,out)
									}
								}


							}
						} else {
							unUTXOs = append(unUTXOs,out)
						}

					}
				}

			}

		}


		fmt.Println(spentTXOutputs)

		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)

		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break;
		}

	}


	return unUTXOs
}





// 挖掘新的区块
func (blockchain *Blockchain) MineNewBlock(from []string,to []string,amount []string) {

//	$ ./bc send -from '["juncheng"]' -to '["zhangqiang"]' -amount '["2"]'
//	[juncheng]
//	[zhangqiang]
//	[2]



	//1.建立一笔交易

	fmt.Println(from)
	fmt.Println(to)
	fmt.Println(amount)

	value ,_ := strconv.Atoi(amount[0])

	tx := NewSimpleTransaction(from[0],to[0],value)
	fmt.Println(tx)


	//1. 通过相关算法建立Transaction数组

	var txs []*Transaction
	txs = append(txs,tx)

	var block *Block

	blockchain.DB.View(func(tx *bolt.Tx) error{

		b := tx.Bucket([]byte(blockTableName))
		if b != nil {

			hash := b.Get([]byte("l"))

			blockBytes := b.Get(hash)

			block = DeserializeBlock(blockBytes)

		}

		return nil
	})


	//2. 建立新的区块
	block = NewBlock(txs,block.Height+1,block.Hash)

	//将新区块存储到数据库
	blockchain.DB.Update(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {

			b.Put(block.Hash,block.Serialize())

			b.Put([]byte("l"),block.Hash)

			blockchain.Tip = block.Hash

		}
		return nil
	})


}

// 查询余额
func (blockchain *Blockchain) GetBalance(address string) int64 {

	utxos := blockchain.UnUTXOs(address)

	var amount int64

	for _,out := range utxos {

		amount = amount + out.Value
	}

	return amount
}


