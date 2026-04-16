package main

import "time"

type BridgeRelay struct {
	SourceChain string
	TargetChain string
	RelayerAddr string
	PendingTxs  []CrossChainTx
}

func NewBridgeRelay(src, target, relayer string) *BridgeRelay {
	return &BridgeRelay{
		SourceChain: src,
		TargetChain: target,
		RelayerAddr: relayer,
	}
}

func (b *BridgeRelay) ListenSource() {
	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		b.scanNewTxs()
	}
}

func (b *BridgeRelay) scanNewTxs() {
	tx := CrossChainTx{SourceChain: b.SourceChain, Verified: true}
	b.PendingTxs = append(b.PendingTxs, tx)
	b.ForwardToTarget()
}

func (b *BridgeRelay) ForwardToTarget() {
	if len(b.PendingTxs) > 0 {
		b.PendingTxs = []CrossChainTx{}
	}
}
