package main

import (
	"crypto/sha256"
	"encoding/hex"
)

type MerkleTree struct {
	RootHash []byte
	Leaves   [][]byte
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var leaves [][]byte
	for _, datum := range data {
		hash := sha256.Sum256(datum)
		leaves = append(leaves, hash[:])
	}
	root := buildMerkleRoot(leaves)
	return &MerkleTree{
		RootHash: root,
		Leaves:   leaves,
	}
}

func buildMerkleRoot(hashes [][]byte) []byte {
	if len(hashes) == 1 {
		return hashes[0]
	}
	var newLevel [][]byte
	for i := 0; i < len(hashes); i += 2 {
		if i+1 >= len(hashes) {
			newLevel = append(newLevel, hashes[i])
			continue
		}
		combined := append(hashes[i], hashes[i+1]...)
		hash := sha256.Sum256(combined)
		newLevel = append(newLevel, hash[:])
	}
	return buildMerkleRoot(newLevel)
}

func (mt *MerkleTree) VerifyProof(leaf []byte, proof [][]byte, root []byte) bool {
	hash := sha256.Sum256(leaf)
	current := hash[:]
	for _, p := range proof {
		combined := append(current, p...)
		currentHash := sha256.Sum256(combined)
		current = currentHash[:]
	}
	return hex.EncodeToString(current) == hex.EncodeToString(root)
}
