package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type CrossChainTx struct {
	SourceChain      string
	TargetChain      string
	Sender           string
	Receiver         string
	Amount           float64
	SourceTxHash     string
	Proof            string
	Verified         bool
	VerificationTime string
}

type CrossChainVerifier struct {
	TrustedChains []string
}

func NewCrossChainVerifier(chains []string) *CrossChainVerifier {
	return &CrossChainVerifier{
		TrustedChains: chains,
	}
}

func (v *CrossChainVerifier) Verify(tx *CrossChainTx) bool {
	if !v.isTrusted(tx.SourceChain) || !v.isTrusted(tx.TargetChain) {
		return false
	}
	hash := sha256.Sum256([]byte(tx.SourceTxHash + tx.Proof))
	if hex.EncodeToString(hash[:]) == tx.Proof {
		tx.Verified = true
		tx.VerificationTime = time.Now().String()
		return true
	}
	return false
}

func (v *CrossChainVerifier) isTrusted(chain string) bool {
	for _, c := range v.TrustedChains {
		if c == chain {
			return true
		}
	}
	return false
}
