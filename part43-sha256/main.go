package main

import (
	"crypto/sha256"
	"fmt"
)

func main()  {

	// 256
	//7b28134d636423d716aaa45227099629783ae6d6012b331049466658b88bf3b5
	// 0111 1011
	hasher := sha256.New()

	hasher.Write([]byte("http://liyuechun.org"))

	bytes := hasher.Sum(nil)

	fmt.Printf("%x\n",bytes)
	//7b28134d636423d716aaa45227099629783ae6d6012b331049466658b88bf3b5

}
