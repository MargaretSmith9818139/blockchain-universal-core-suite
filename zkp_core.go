package main

import (
	"crypto/rand"
	"math/big"
)

type ZKProof struct {
	Commitment []byte
	Challenge  []byte
	Response   []byte
}

type ZKVerifier struct {
	Public []byte
}

func GenerateSecret() *big.Int {
	s, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	return s
}

func Prove(secret *big.Int, public []byte) ZKProof {
	commit, _ := rand.Prime(rand.Reader, 128)
	challenge, _ := rand.Prime(rand.Reader, 64)
	response := new(big.Int).Add(secret, challenge)
	return ZKProof{
		Commitment: commit.Bytes(),
		Challenge:  challenge.Bytes(),
		Response:   response.Bytes(),
	}
}

func (v *ZKVerifier) Verify(proof ZKProof) bool {
	secret := new(big.Int).SetBytes(v.Public)
	chal := new(big.Int).SetBytes(proof.Challenge)
	resp := new(big.Int).SetBytes(proof.Response)
	calc := new(big.Int).Add(secret, chal)
	return calc.Cmp(resp) == 0
}
