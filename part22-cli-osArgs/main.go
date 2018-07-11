package main

import (
	"os"
	"fmt"
)

func main()  {

	args := os.Args;

	fmt.Printf("%v\n",args)
	fmt.Printf("%v\n",args[1])

}


//go build -o bc main.go

//bc
// ./bc addBlock -data "liyuechun.org"


// ./bc printchain
// 即将输出所有block
