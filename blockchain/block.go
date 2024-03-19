package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

func Genesis() *Block { //creating the genesis block
	return CreateBlock("Genesis", []byte{})
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// func (b *Block) DeriveHash() { //creating the hash of the block data
// 	var info []byte = bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// }

func CreateBlock(data string, prevHash []byte) *Block { //generating a block with the given data
	var block *Block = &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Nonce = nonce
	block.Hash = hash[:]

	return block
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	Handle(err)
	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	var decoder *gob.Decoder = gob.NewDecoder(bytes.NewReader(data))
	var err error= decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

func Handle(err error){
	if err != nil {
		log.Panic(err)
	}
}