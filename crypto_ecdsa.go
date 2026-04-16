package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

type KeyPair struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

func GenerateKeyPair() (KeyPair, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return KeyPair{}, err
	}
	return KeyPair{
		PrivateKey: privateKey,
		PublicKey:  &privateKey.PublicKey,
	}, nil
}

func SignMessage(privateKey *ecdsa.PrivateKey, message []byte) ([]byte, []byte, error) {
	hash := sha256.Sum256(message)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return nil, nil, err
	}
	return r.Bytes(), s.Bytes(), nil
}

func VerifySignature(publicKey *ecdsa.PublicKey, message []byte, rBytes, sBytes []byte) bool {
	hash := sha256.Sum256(message)
	var r, s big.Int
	r.SetBytes(rBytes)
	s.SetBytes(sBytes)
	return ecdsa.Verify(publicKey, hash[:], &r, &s)
}

func PublicKeyToBytes(publicKey *ecdsa.PublicKey) []byte {
	return elliptic.Marshal(elliptic.P256(), publicKey.X, publicKey.Y)
}
