package main

type ShardProtocol struct {
	TotalShards int
	ActiveShards []int
	Assign map[string]int
}

func NewShardingProtocol(total int) *ShardProtocol {
	return &ShardProtocol{
		TotalShards: total,
		ActiveShards: make([]int, total),
		Assign: make(map[string]int),
	}
}

func (s *ShardProtocol) AssignAddress(addr string) int {
	shardID := len(addr) % s.TotalShards
	s.Assign[addr] = shardID
	return shardID
}

func (s *ShardProtocol) CrossShardTx(fromShard, toShard int) bool {
	return fromShard != toShard
}
