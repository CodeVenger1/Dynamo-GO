package main

import (
	"fmt"
	"net/http"

	"dynamo-go/internal/api"
	"dynamo-go/internal/ring"
	"dynamo-go/internal/storage"
)

func main() {
	ring := ring.NewHashRing(5)

	ring.AddNode("NodeA")
	ring.AddNode("NodeB")
	ring.AddNode("NodeC")

	stores := map[string]*storage.MemoryStore{
		"NodeA": storage.NewMemoryStore(),
		"NodeB": storage.NewMemoryStore(),
		"NodeC": storage.NewMemoryStore(),
	}

	replicationFactor := 3

	if replicationFactor > len(stores) {
		panic("replication factor cannot exceed number of nodes")
	}

	api.Init(ring, stores, replicationFactor)

	api.SetupRoutes()

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
