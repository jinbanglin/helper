package helper

import "sync"

const shardCount = uint64(32)

type ConcurrentMapShared struct {
	items map[uint64]interface{}
	sync.RWMutex
}

type ConcurrentMap []*ConcurrentMapShared

func New() ConcurrentMap {
	m := make([]*ConcurrentMapShared, shardCount)
	for i := uint64(0); i < shardCount; i++ {
		m[i] = &ConcurrentMapShared{items: make(map[uint64]interface{})}
	}
	return m
}

func (c ConcurrentMap) Set(key uint64, value interface{}) {
	shard := c[key%shardCount]
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

func (c ConcurrentMap) Get(key uint64) (interface{}, bool) {
	shard := c[key%shardCount]
	shard.RLock()
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

func (c ConcurrentMap) Remove(key uint64) {
	shard := c[key%shardCount]
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}
