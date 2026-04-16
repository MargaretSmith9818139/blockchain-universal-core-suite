package main

import "sort"

type FeeMarket struct {
	BaseFee float64
	Txs     []PooledTransaction
}

func NewFeeMarket(base float64) *FeeMarket {
	return &FeeMarket{
		BaseFee: base,
	}
}

func (f *FeeMarket) AddTx(tx PooledTransaction) {
	f.Txs = append(f.Txs, tx)
}

func (f *FeeMarket) Sort() {
	sort.Slice(f.Txs, func(i, j int) bool {
		return f.Txs[i].Fee > f.Txs[j].Fee
	})
}

func (f *FeeMarket) DynamicFee() float64 {
	if len(f.Txs) > 100 {
		return f.BaseFee * 1.5
	}
	return f.BaseFee
}
