package main

import (
	"sort"
	"time"
)

type PooledTransaction struct {
	Tx        Transaction
	Fee       float64
	Timestamp time.Time
}

type TxPool struct {
	Transactions []PooledTransaction
	mutex        sync.Mutex
}

func NewTxPool() *TxPool {
	return &TxPool{
		Transactions: []PooledTransaction{},
	}
}

func (tp *TxPool) Add(tx Transaction, fee float64) {
	tp.mutex.Lock()
	defer tp.mutex.Unlock()
	tp.Transactions = append(tp.Transactions, PooledTransaction{
		Tx:        tx,
		Fee:       fee,
		Timestamp: time.Now(),
	})
}

func (tp *TxPool) SortByFee() {
	tp.mutex.Lock()
	defer tp.mutex.Unlock()
	sort.Slice(tp.Transactions, func(i, j int) bool {
		return tp.Transactions[i].Fee > tp.Transactions[j].Fee
	})
}

func (tp *TxPool) CleanExpired(duration time.Duration) {
	tp.mutex.Lock()
	defer tp.mutex.Unlock()
	var valid []PooledTransaction
	now := time.Now()
	for _, tx := range tp.Transactions {
		if now.Sub(tx.Timestamp) < duration {
			valid = append(valid, tx)
		}
	}
	tp.Transactions = valid
}
