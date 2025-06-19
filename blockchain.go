package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	return &Block{
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
		previousHash: previousHash,
	}
}
func (b *Block) Print() {
	fmt.Printf("timestamp: %d\n", b.timestamp)
	fmt.Printf("nonce: %d\n", b.nonce)
	fmt.Printf("previousHash: %s\n", b.previousHash)
	fmt.Printf("transactions: %s\n", b.transactions)
}
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)

	return sha256.Sum256([]byte(m))
}

type BlockChain struct {
	transactionPool []string
	Blocks          []*Block
}

func NewBlockChain() *BlockChain {
	bc := new(BlockChain)
	bc.CreateBlock(0, "init hash")
	return bc
}
func (bc *BlockChain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.Blocks = append(bc.Blocks, b)
	return b
}
func (bc *BlockChain) Print() {
	for i, b := range bc.Blocks {
		color.Green(fmt.Sprintf("Block %d     ----------------------------------\n", i))
		b.Print()
		color.Red(" ----------------------------------")
	}
}

func main() {

	blockchain := NewBlockChain()
	blockchain.Print()
}
