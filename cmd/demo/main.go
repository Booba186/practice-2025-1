package main

import (
	"encoding/hex"
	"fmt"
	"github.com/Booba186/practice-2025-1/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain(2)

	bc.AddBlock("Send 1 BTC to Booba")
	bc.AddBlock("Send 10 BTC to Booba")

	for _, block := range bc.Blocks {
		fmt.Println("===============")
		fmt.Printf("Индекс: %d\n", block.Index)
		fmt.Printf("Дата: %d\n", block.Timestamp)
		fmt.Printf("Данные: %s\n", block.Data)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Хэш: %s\n", hex.EncodeToString(block.Hash))
	}
}
