package api

import (
	"dynamo-go/internal/coordinator"
	"dynamo-go/internal/ring"
	"dynamo-go/internal/storage"
)

var hashRing *ring.HashRing
var stores map[string]*storage.MemoryStore
var coord *coordinator.Coordinator

func Init(r *ring.HashRing, s map[string]*storage.MemoryStore, rf, wq, rq int) {
	hashRing = r
	stores = s
	coord = coordinator.NewCoordinator(r, s, rf, wq, rq)
}
