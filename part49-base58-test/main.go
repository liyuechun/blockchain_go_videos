package main

import (
	"fmt"
	"kongyixueyuan.com/publicChain/part48-base58/BLC"
)

func main()  {


	//1. 创建钱包
	//（1）私钥
	//（2）公钥

	//2.先将公钥进行一次256hash，再进行一次160hash
	// 20字节的[]byte

	//version {0} + hash160 -> pubkey

	//256hash pubkey 几次
	// 256 64
	// 最前面的四个字节取出来
	// version {0} + hash160 + 4个字节 -》 25字节

	//base58  编码

	hash := []byte{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}
	// 25字节

	fmt.Printf("%d\n",len(hash))

	bytes58 := BLC.Base58Encode(hash)

	//1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa

	//1QRus492mJL2Cum4E2TSqUmjdCBE5Ke9hi

	fmt.Printf("%s\n",bytes58)

}
