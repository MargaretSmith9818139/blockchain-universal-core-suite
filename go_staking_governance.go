package main

import "time"

type StakingNode struct {
	Address string
	Amount  float64
	Since   time.Time
}

type Governance struct {
	Stakers map[string]StakingNode
	MinStake float64
}

func NewGovernance(min float64) *Governance {
	return &Governance{
		Stakers: make(map[string]StakingNode),
		MinStake: min,
	}
}

func (g *Governance) Stake(addr string, amount float64) bool {
	if amount < g.MinStake {
		return false
	}
	g.Stakers[addr] = StakingNode{
		Address: addr,
		Amount: amount,
		Since: time.Now(),
	}
	return true
}

func (g *Governance) Unstake(addr string) {
	delete(g.Stakers, addr)
}
