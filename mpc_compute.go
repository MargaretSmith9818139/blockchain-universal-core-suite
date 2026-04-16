package main

import (
	"math/big"
	"math/rand"
	"time"
)

type MPCNode struct {
	ID    string
	Shares []*big.Int
}

func NewMPCNode(id string) *MPCNode {
	return &MPCNode{ID: id, Shares: []*big.Int{}}
}

func SplitSecret(secret int64, parts int) []*big.Int {
	rand.Seed(time.Now().UnixNano())
	shares := make([]*big.Int, parts)
	total := big.NewInt(0)
	for i := 0; i < parts-1; i++ {
		shares[i] = big.NewInt(rand.Int63n(1000))
		total.Add(total, shares[i])
	}
	last := new(big.Int).Sub(big.NewInt(secret), total)
	shares[parts-1] = last
	return shares
}

func CombineShares(shares []*big.Int) *big.Int {
	sum := big.NewInt(0)
	for _, s := range shares {
		sum.Add(sum, s)
	}
	return sum
}

func (n *MPCNode) AddShare(share *big.Int) {
	n.Shares = append(n.Shares, share)
}
