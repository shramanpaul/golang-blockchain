package main

import (
	"fmt"
	"strconv"

	"github.com/shramanpaul/golang-blockchain/blockchain"
)

func main() {
	var chain blockchain.BlockChain = blockchain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, blocks := range chain.Blocks { //match them with the pows generated
		fmt.Printf("Previous Hash: %x\n", blocks.PrevHash)
		fmt.Printf("Data in Block: %s\n", blocks.Data)
		fmt.Printf("Hash: %x\n", blocks.Hash)

		pow := blockchain.NewProof(blocks)                          //to show at the end if the proof of work is valid or not
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate())) //checks the output of pow.Validate()
		fmt.Println()
	}
}
