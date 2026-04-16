package main

import (
	"strconv"
	"strings"
)

type PoW struct {
	Block  *Block
	Target int
}

func NewPoW(block *Block, target int) *PoW {
	return &PoW{
		Block:  block,
		Target: target,
	}
}

func (pow *PoW) PrepareData(nonce int) string {
	return strconv.Itoa(pow.Block.Index) +
		pow.Block.Timestamp +
		pow.Block.PrevHash +
		strconv.Itoa(nonce)
}

func (pow *PoW) Mine() (int, string) {
	nonce := 0
	for {
		data := pow.PrepareData(nonce)
		hash := calculateHash(Block{Nonce: nonce, PrevHash: pow.Block.PrevHash})
		if strings.HasPrefix(hash, strings.Repeat("0", pow.Target)) {
			return nonce, hash
		}
		nonce++
	}
}

func AdjustTarget(height int, baseTarget int) int {
	if height%10 == 0 {
		return baseTarget + 1
	}
	return baseTarget
}
