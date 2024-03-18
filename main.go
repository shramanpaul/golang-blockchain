package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct { //for maintaining multiple blocks
	blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) { //for adding multiple blocks
	var prevBlock *Block = chain.blocks[len(chain.blocks)-1]
	var new *Block = CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {	//creating the genesis block
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() BlockChain {	//initializing the blockchain with genesis
	return BlockChain{blocks: []*Block{Genesis()}}
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {  //creating the hash of the block data
	var info []byte = bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block { //generating a block with the given data
	var block *Block = &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func main() {
	var chain BlockChain = InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, blocks := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", blocks.PrevHash)
		fmt.Printf("Data in Block: %s\n", blocks.Data)
		fmt.Printf("Hash: %x\n", blocks.Hash)
	}
}
