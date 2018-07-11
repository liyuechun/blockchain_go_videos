package BLC

import (
	"fmt"
	"os"
	"flag"
	"log"
)

type CLI struct {
	BC *Blockchain
}



func printUsage()  {

	fmt.Println("Usage:")
	fmt.Println("\tcreateblockchainwithgenesis -data -- 交易数据.")
	fmt.Println("\taddblock -data DATA -- 交易数据.")
	fmt.Println("\tprintchain -- 输出区块信息.")

}

func isValidArgs()  {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string)  {
	cli.BC.AddBlockToBlockchain(data)
}

func (cli *CLI) printchain()  {
	cli.BC.Printchain()
}



func (cli *CLI) Run()  {
	isValidArgs()

	addBlockCmd := flag.NewFlagSet("addblock",flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain",flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data","http://liyuechun.org","交易数据......")



	switch os.Args[1] {
		case "addblock":
			err := addBlockCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "printchain":
			err := printChainCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		default:
			printUsage()
			os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}

		//fmt.Println(*flagAddBlockData)
		cli.addBlock(*flagAddBlockData)
	}

	if printChainCmd.Parsed() {

		//fmt.Println("输出所有区块的数据........")
		cli.printchain()
	}
}