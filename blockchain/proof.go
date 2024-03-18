package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

const Difficulty = 18

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	var target *big.Int = big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty)) //left shift the target by 256 - Difficulty
	return &ProofOfWork{b, target}           //return the proof of work
}
func (pow *ProofOfWork) InitData(nonce int) []byte { //replacing the DeriveHash function
	var data []byte = bytes.Join([][]byte{pow.Block.PrevHash, pow.Block.Data, ToHex(int64(nonce)), ToHex(int64(Difficulty))}, []byte{})
	return data
}
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func (pow *ProofOfWork) Run() (int, []byte) { //for hashing the block data and checking if it meets the difficulty criteria
	var intHash big.Int
	var hash [32]byte

	nonce := 0
	for nonce < math.MaxInt64 { //very big number basically infinite loop
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash) //we will be printing the hash to the console to see the progress
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 { //if the hash is less than the target
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool { //for validating the proof of work
	var intHash big.Int
	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}
