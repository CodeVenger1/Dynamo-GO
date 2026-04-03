package coordinator

import (
	"dynamo-go/internal/ring"
	"dynamo-go/internal/storage"
)

type Coordinator struct {
	ring   *ring.HashRing
	stores map[string]*storage.MemoryStore

	replicationFactor int
	Wq                int
	Rq                int
}

// constructor
func NewCoordinator(r *ring.HashRing, s map[string]*storage.MemoryStore, rf, wq, rq int) *Coordinator {
	return &Coordinator{
		ring:              r,
		stores:            s,
		replicationFactor: rf,
		Wq:                wq,
		Rq:                rq,
	}
}

// func Put: Put with replication(Write Quorum)
func (c *Coordinator) Put(key, value string) bool {
	nodes := c.ring.GetNodes(key, c.replicationFactor)

	success := 0

	for _, node := range nodes {
		c.stores[node].Put(key, value)
		success++

		if success >= c.Wq {
			return true
		}
	}

	return false
}

// func Get: Get with fallback(Read Quorum)
func (c *Coordinator) Get(key string) (string, bool) {
	nodes := c.ring.GetNodes(key, c.replicationFactor)

	count := 0
	var result string

	for _, node := range nodes {
		val, ok := c.stores[node].Get(key)

		if ok {
			result = val
			count++

			if count >= c.Rq {
				return result, true
			}
		}
	}

	return "", false
}
