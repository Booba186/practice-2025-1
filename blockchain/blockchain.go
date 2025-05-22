package blockchain

import (
	"encoding/gob"
	"log"
	"os"
	"time"
)

type Blockchain struct {
	Blocks     []*Block
	Difficulty int
}

func CreateGenesis(diff int) *Block {
	gen := &Block{
		Index:        0,
		Timestamp:    time.Now().Unix(),
		Transactions: []Transaction{{From: "system", To: "miner", Amount: 50}},
		PrevHash:     []byte{},
	}
	gen.Mine(diff)
	return gen
}

func NewBlockchain(diff int) *Blockchain {
	return &Blockchain{
		Blocks:     []*Block{CreateGenesis(diff)},
		Difficulty: diff,
	}
}

func (bc *Blockchain) AddBlock(txs []Transaction) {
	prev := bc.Blocks[len(bc.Blocks)-1]
	blk := &Block{
		Index:        prev.Index + 1,
		Timestamp:    time.Now().Unix(),
		Transactions: txs,
		PrevHash:     prev.Hash,
	}
	blk.Mine(bc.Difficulty)
	bc.Blocks = append(bc.Blocks, blk)
}

func (bc *Blockchain) Save(path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	if err := enc.Encode(bc.Blocks); err != nil {
		log.Fatal(err)
	}
}

func Load(path string, diff int) *Blockchain {
	f, err := os.Open(path)
	if err != nil {
		return NewBlockchain(diff)
	}
	defer f.Close()

	var blocks []*Block
	dec := gob.NewDecoder(f)
	if err := dec.Decode(&blocks); err != nil {
		log.Fatal(err)
	}
	return &Blockchain{Blocks: blocks, Difficulty: diff}
}
