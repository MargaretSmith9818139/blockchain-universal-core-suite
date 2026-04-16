package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type UTXO struct {
	TxID      string
	Index     int
	Address   string
	Amount    float64
	Spent     bool
	CreatedAt string
}

type UTXOSet struct {
	UTXOs map[string]UTXO
}

func NewUTXOSet() *UTXOSet {
	return &UTXOSet{
		UTXOs: make(map[string]UTXO),
	}
}

func (u *UTXOSet) AddUTXO(txID string, index int, address string, amount float64) {
	key := txID + string(index)
	u.UTXOs[key] = UTXO{
		TxID:      txID,
		Index:     index,
		Address:   address,
		Amount:    amount,
		Spent:     false,
		CreatedAt: time.Now().String(),
	}
}

func (u *UTXOSet) SpendUTXO(key string) {
	utxo := u.UTXOs[key]
	utxo.Spent = true
	u.UTXOs[key] = utxo
}

func (u *UTXOSet) GetBalance(address string) float64 {
	balance := 0.0
	for _, utxo := range u.UTXOs {
		if utxo.Address == address && !utxo.Spent {
			balance += utxo.Amount
		}
	}
	return balance
}

func GenerateTxID(from, to string, amount float64) string {
	hash := sha256.Sum256([]byte(from + to + string(amount) + time.Now().String()))
	return hex.EncodeToString(hash[:])
}
