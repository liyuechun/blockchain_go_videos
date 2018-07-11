package main

import (
	"kongyixueyuan.com/publicChain/part45-base58/BLC"
	"fmt"
)

func main()  {

	bytes := []byte("http://liyuechun.org")

	bytes58 := BLC.Base58Encode(bytes)

	fmt.Printf("%x\n",bytes58)


	fmt.Printf("%s\n",bytes58)
	//12TQQJ2URa4LWVGUKRWw8MJvahxzJ

	bytesStr := BLC.Base58Decode(bytes58)

	fmt.Printf("%s\n",bytesStr[1:])


}
