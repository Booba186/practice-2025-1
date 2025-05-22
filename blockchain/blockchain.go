package blockchain

import "time"

type Blockchain struct {
	Blocks     []*Block
	Difficulty int
}

func CreateGenesis(difficulty int) *Block {
	genesis := &Block{
		Index:     0,
		Timestamp: time.Now().Unix(),
		Data:      "Genesis Block",
		PrevHash:  []byte{},
	}
	genesis.Mine(difficulty)
	return genesis
}

func NewBlockchain(difficulty int) *Blockchain {
	genesis := CreateGenesis(difficulty)
	return &Blockchain{
		Blocks:     []*Block{genesis},
		Difficulty: difficulty,
	}
}

func (bc *Blockchain) AddBlock(data string) {
	prev := bc.Blocks[len(bc.Blocks)-1]
	newBlock := &Block{
		Index:     prev.Index + 1,
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prev.Hash,
	}
	newBlock.Mine(bc.Difficulty)
	bc.Blocks = append(bc.Blocks, newBlock)
}
