package main

import (
	"fmt"
)
type Blockchain struct {
	blocks []*Block
}

type Block struct {
	Hash 	 []byte
	Data 	 []byte
	PrevHash []byte
}

func Genesis() *Block {
	return &Block{[]byte{}, []byte("Genesis"), []byte{}}
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockchain()

	for _, block := range chain.blocks {
		fmt.Printf("\n")
		fmt.Printf("Prev Hash : %v\n", block.PrevHash)
		fmt.Printf("Data : %v\n", block.Data)
		fmt.Printf("Hash : %v\n", block.Hash)
	}
}