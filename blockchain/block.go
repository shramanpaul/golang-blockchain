package blockchain

type BlockChain struct { //for maintaining multiple Blocks
	Blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) { //for adding multiple Blocks
	var prevBlock *Block = chain.Blocks[len(chain.Blocks)-1]
	var new *Block = CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block { //creating the genesis block
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() BlockChain { //initializing the blockchain with genesis
	return BlockChain{Blocks: []*Block{Genesis()}}
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
