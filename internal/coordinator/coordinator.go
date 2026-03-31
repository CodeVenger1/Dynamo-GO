package coordinator

import (
	"dynamo-go/internal/ring"
	"dynamo-go/internal/storage"
	"fmt"
)

type Coordinator struct {
	ring              *ring.HashRing
	stores            map[string]*storage.MemoryStore
	replicationFactor int
}

// constructor
func NewCoordinator(r *ring.HashRing, s map[string]*storage.MemoryStore, rf int) *Coordinator {
	return &Coordinator{
		ring:              r,
		stores:            s,
		replicationFactor: rf,
	}
}

// func Put: Put with replication
func (c *Coordinator) Put(key, value string) {
	fmt.Println("PUT key:", key)
	nodes := c.ring.GetNodes(key, c.replicationFactor)
	fmt.Println("Replicating to nodes:", nodes)
	for _, node := range nodes {
		c.stores[node].Put(key, value)
		fmt.Println("Writing to node:", node)
	}

}

// func Get: Get with fallback
func (c *Coordinator) Get(key string) (string, bool) {
	fmt.Println("GET key:", key)

	nodes := c.ring.GetNodes(key, c.replicationFactor)

	for _, node := range nodes {
		val, ok := c.stores[node].Get(key)
		fmt.Println("Checking node:", node, "Found:", ok)

		if ok {
			fmt.Println("Read from node:", node)
			return val, true
		}
	}
	return "", false
}
