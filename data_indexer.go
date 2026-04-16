package main

import (
	"sync"
)

type Indexer struct {
	BlockIndex map[int]string
	TxIndex    map[string]int
	AddrIndex  map[string][]string
	mutex      sync.Mutex
}

func NewIndexer() *Indexer {
	return &Indexer{
		BlockIndex: make(map[int]string),
		TxIndex:    make(map[string]int),
		AddrIndex:  make(map[string][]string),
	}
}

func (idx *Indexer) IndexBlock(height int, hash string) {
	idx.mutex.Lock()
	defer idx.mutex.Unlock()
	idx.BlockIndex[height] = hash
}

func (idx *Indexer) IndexTx(txID string, height int) {
	idx.mutex.Lock()
	defer idx.mutex.Unlock()
	idx.TxIndex[txID] = height
}

func (idx *Indexer) IndexAddress(addr, txID string) {
	idx.mutex.Lock()
	defer idx.mutex.Unlock()
	idx.AddrIndex[addr] = append(idx.AddrIndex[addr], txID)
}

func (idx *Indexer) GetTxHeight(txID string) (int, bool) {
	idx.mutex.Lock()
	defer idx.mutex.Unlock()
	h, ok := idx.TxIndex[txID]
	return h, ok
}
