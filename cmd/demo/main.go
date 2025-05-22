package main

import (
	"encoding/hex"
	"flag"
	"fmt"

	bc "github.com/Booba186/practice-2025-1/blockchain"
)

func main() {
	difficulty := flag.Int("difficulty", 2, "кол-во ведущих нулей")
	flag.Parse()

	chain := bc.Load("chain.gob", *difficulty)


	tx1 := bc.Transaction{From: "Alex", To: "Booba", Amount: 15}
	tx2 := bc.Transaction{From: "Adolf", To: "Sergi", Amount: 34}
	chain.AddBlock([]bc.Transaction{tx1, tx2})

	for _, block := range chain.Blocks {
		fmt.Println("===============")
		fmt.Printf("Индекс: %d\n", block.Index)
		fmt.Printf("Время : %d\n", block.Timestamp)
		fmt.Printf("Nonce : %d\n", block.Nonce)
		fmt.Println("Транзакции:")
		for _, t := range block.Transactions {
			fmt.Printf("  • %s → %s : %d\n", t.From, t.To, t.Amount)
		}
		fmt.Printf("Хэш  : %s\n", hex.EncodeToString(block.Hash))
	}

	chain.Save("chain.gob")
}
