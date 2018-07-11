package BLC

import (
	"time"
	"strconv"
	"fmt"
	"bytes"
	"crypto/sha256"
)

type Block struct {
	//1. 区块高度
	Height int64
	//2. 上一个区块HASH
	PrevBlockHash []byte
	//3. 交易数据
	Data []byte
	//4. 时间戳
	Timestamp int64
	//5. Hash
	Hash []byte
}

func (block *Block) SetHash()  {

	//	1. Height []byte

	heightBytes := IntToHex(block.Height)

	fmt.Println(heightBytes)

	//  2. 将时间戳转[]byte

	// 2 ~ 36
	// 1528097970
	timeString := strconv.FormatInt(block.Timestamp,2)

	fmt.Println(timeString)
	timeBytes := []byte(timeString)

	fmt.Println(timeBytes)

	//  3. 拼接所有属性
	blockBytes := bytes.Join([][]byte{heightBytes,block.PrevBlockHash,block.Data,timeBytes},[]byte{})

	// 4. 生成Hash
	hash := sha256.Sum256(blockBytes)

	block.Hash = hash[:]

}

//1. 创建新的区块
func NewBlock(data string,height int64,prevBlockHash []byte) *Block {

	//创建区块
	block := &Block{height,prevBlockHash,[]byte(data),time.Now().Unix(),nil}

	//设置Hash
	block.SetHash()

	return block

}

