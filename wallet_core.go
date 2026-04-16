package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
	Address    string
}

func NewWallet() *Wallet {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pub := elliptic.Marshal(elliptic.P256(), priv.X, priv.Y)
	addr := generateAddress(pub)
	return &Wallet{
		PrivateKey: priv,
		PublicKey:  pub,
		Address:    addr,
	}
}

func generateAddress(pub []byte) string {
	hash := sha256.Sum256(pub)
	return hex.EncodeToString(hash[:])[:40]
}

func (w *Wallet) SignTransaction(tx *Transaction) error {
	hash := sha256.Sum256([]byte(tx.From + tx.To + string(tx.Amount)))
	r, s, err := ecdsa.Sign(rand.Reader, w.PrivateKey, hash[:])
	if err != nil {
		return err
	}
	tx.TxID = hex.EncodeToString(append(r.Bytes(), s.Bytes()...))
	return nil
}
