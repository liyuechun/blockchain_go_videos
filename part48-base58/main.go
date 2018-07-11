package main

import (
	"fmt"
	"crypto/sha256"
	"kongyixueyuan.com/publicChain/part48-base58/BLC"
)

func main()  {

	bytes := []byte("http://liyuechun.orghttp://liyuechun.orghttp://liyuechun.org")

	//1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa
	//17fHw3nUseFctpqUN6EQ5gsGTfEwwgxiK97WJBF3mcz8v

	hasher := sha256.New()
	hasher.Write(bytes)
	hash := hasher.Sum(nil)


	fmt.Printf("%x\n",hash)
	//
	bytes58 := BLC.Base58Encode(hash)



	fmt.Printf("%s\n",bytes58)



	bytesStr := BLC.Base58Decode(bytes58)

	fmt.Printf("%x\n",bytesStr[1:])


}
