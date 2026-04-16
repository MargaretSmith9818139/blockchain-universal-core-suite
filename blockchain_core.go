package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	PrevHash     string
	Hash         string
	Nonce        int
}

type Transaction struct {
	TxID      string
	From      string
	To        string
	Amount    float64
	Timestamp string
}

type Blockchain struct {
	Chain  []Block
	Pool   []Transaction
	NodeID string
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.PrevHash + block.TxID
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (bc *Blockchain) CreateBlock(nonce int, prevHash string) Block {
	newBlock := Block{
		Index:        len(bc.Chain) + 1,
		Timestamp:    time.Now().String(),
		Transactions: bc.Pool,
		PrevHash:     prevHash,
		Nonce:        nonce,
	}
	newBlock.Hash = calculateHash(newBlock)
	bc.Pool = []Transaction{}
	bc.Chain = append(bc.Chain, newBlock)
	return newBlock
}

func (bc *Blockchain) AddTransaction(tx Transaction) {
	tx.Timestamp = time.Now().String()
	tx.TxID = hex.EncodeToString(sha256.New().Sum([]byte(tx.From + tx.To + string(tx.Amount) + tx.Timestamp)))
	bc.Pool = append(bc.Pool, tx)
}

func (bc *Blockchain) IsChainValid() bool {
	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		prevBlock := bc.Chain[i-1]
		if currentBlock.Hash != calculateHash(currentBlock) {
			return false
		}
		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}
	}
	return true
}
