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
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
		previousHash: previousHash,
		transactions: transactions,
	}
}
func (b *Block) Print() {
	fmt.Printf("timestamp: %d\n", b.timestamp)
	fmt.Printf("nonce: %d\n", b.nonce)
	fmt.Printf("previousHash: %x\n", b.previousHash)
	fmt.Printf("transactions:\n")
	for _, t := range b.transactions {
		t.Print()
	}
}
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)

	return sha256.Sum256([]byte(m))
}
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previousHash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

type BlockChain struct {
	transactionPool []*Transaction
	Blocks          []*Block
}

func NewBlockChain() *BlockChain {
	b := &Block{}
	bc := new(BlockChain)
	bc.CreateBlock(0, b.Hash())
	return bc
}
func (bc *BlockChain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.Blocks = append(bc.Blocks, b)
	bc.transactionPool = []*Transaction{}
	return b
}
func (bc *BlockChain) Print() {
	for i, b := range bc.Blocks {
		color.Green(fmt.Sprintf("Block %d     ----------------------------------\n", i))
		b.Print()
		color.Red(" ----------------------------------")
	}
}
func (bc *BlockChain) LastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *BlockChain) AddTransaction(sender string, recipient string, amount float32) {
	bc.transactionPool = append(bc.transactionPool, NewTransaction(sender, recipient, amount))
}

type Transaction struct {
	Sender    string
	Recipient string
	Amount    float32
}

func NewTransaction(sender string, recipient string, amount float32) *Transaction {
	return &Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender"`
		Recipient string  `json:"recipient"`
		Amount    float32 `json:"amount"`
	}{
		Sender:    t.Sender,
		Recipient: t.Recipient,
		Amount:    t.Amount,
	})
}

func (t *Transaction) Print() {
	fmt.Printf("%s -> %s: %f\n", t.Sender, t.Recipient, t.Amount)
}

func main() {

	blockchain := NewBlockChain()
	blockchain.Print()

	blockchain.AddTransaction("A", "B", 1.0)
	previousHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(1, previousHash)
	blockchain.Print()
}
