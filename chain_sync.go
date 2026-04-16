package main

import (
	"sync"
	"time"
)

type SyncManager struct {
	LocalChain  []Block
	PeerChains  map[string][]Block
	BestHeight  int
	BestChain   []Block
	mutex       sync.Mutex
	Syncing     bool
}

func NewSyncManager(chain []Block) *SyncManager {
	return &SyncManager{
		LocalChain: chain,
		PeerChains: make(map[string][]Block),
		BestHeight: len(chain),
		BestChain:  chain,
	}
}

func (s *SyncManager) ReceivePeerChain(nodeID string, chain []Block) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.PeerChains[nodeID] = chain
	if len(chain) > s.BestHeight {
		s.BestHeight = len(chain)
		s.BestChain = chain
	}
}

func (s *SyncManager) StartSync() {
	s.Syncing = true
	defer func() { s.Syncing = false }()
	time.Sleep(1 * time.Second)
	s.LocalChain = s.BestChain
}

func (s *SyncManager) VerifyChain(chain []Block) bool {
	for i := 1; i < len(chain); i++ {
		if chain[i].PrevHash != chain[i-1].Hash {
			return false
		}
	}
	return true
}
