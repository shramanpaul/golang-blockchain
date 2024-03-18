package main

import (
	"fmt"
	"github.com/shramanpaul/golang-blockchain/blockchain"

)

func main() {
	var chain blockchain.BlockChain = blockchain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, blocks := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", blocks.PrevHash)
		fmt.Printf("Data in Block: %s\n", blocks.Data)
		fmt.Printf("Hash: %x\n", blocks.Hash)
	}
}
