package main

import (
	"crypto/sha256"
	"encoding/hex"
	"sync"
)

type Shard struct {
	ID       string
	Nodes    []string
	Data     map[string][]byte
	Replicas int
}

type DistributedStorage struct {
	Shards map[string]*Shard
	mutex  sync.Mutex
}

func NewDistributedStorage() *DistributedStorage {
	return &DistributedStorage{
		Shards: make(map[string]*Shard),
	}
}

func (ds *DistributedStorage) CreateShard(replicas int) string {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	id := hex.EncodeToString(sha256.New().Sum([]byte(string(replicas))))
	ds.Shards[id] = &Shard{
		ID:       id,
		Replicas: replicas,
		Data:     make(map[string][]byte),
	}
	return id
}

func (ds *DistributedStorage) Put(shardID string, key string, value []byte) bool {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	shard, ok := ds.Shards[shardID]
	if !ok {
		return false
	}
	shard.Data[key] = value
	return true
}

func (ds *DistributedStorage) Get(shardID string, key string) ([]byte, bool) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	shard, ok := ds.Shards[shardID]
	if !ok {
		return nil, false
	}
	val, ok := shard.Data[key]
	return val, ok
}
