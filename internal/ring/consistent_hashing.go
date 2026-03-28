package ring

import (
	"fmt"
	"hash/fnv"
	"sort"
)

// Hashring struct
type HashRing struct {
	nodes  map[uint32]string
	keys   []uint32
	vnodes int
}

// constructor
func NewHashRing(vnodes int) *HashRing {
	return &HashRing{
		nodes:  make(map[uint32]string),
		keys:   []uint32{},
		vnodes: vnodes,
	}
}

// func: hashKey -> determine the hashed value of the string
func hashKey(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

// func: AddNode -> func to add a new node
func (r *HashRing) AddNode(node string) {
	for i := 0; i < r.vnodes; i++ {
		vnodeKey := fmt.Sprintf("%s - %d", node, i)
		hash := hashKey(vnodeKey)
		r.nodes[hash] = node
		r.keys = append(r.keys, hash)
	}

	sort.Slice(r.keys, func(i, j int) bool {
		return r.keys[i] < r.keys[j]
	})
}

// func: RemoveNode -> func to remove a node from the ring
func (r *HashRing) RemoveNode(node string) {
	for i := 0; i < r.vnodes; i++ {
		vnodeKey := fmt.Sprintf("%s-%d", node, i)
		hash := hashKey(vnodeKey)

		delete(r.nodes, hash)

		for j, k := range r.keys {
			if k == hash {
				r.keys = append(r.keys[:j], r.keys[j+1:]...)
				break
			}
		}
	}
}

// func : GetNode -> func to find a single node
func (r *HashRing) GetNode(key string) string {
	if len(r.keys) == 0 {
		return ""
	}
	hash := hashKey(key)

	idx := sort.Search(len(r.keys), func(i int) bool {
		return r.keys[i] >= hash
	})

	if idx == len(r.keys) {
		idx = 0
	}

	return r.nodes[r.keys[idx]]
}

// func : GetNodes -> in case of replication
func (r *HashRing) GetNodes(key string, n int) []string {
	if len(r.keys) == 0 {
		return nil
	}

	hash := hashKey(key)

	idx := sort.Search(len(r.keys), func(i int) bool {
		return r.keys[i] >= hash
	})

	result := []string{}
	seen := make(map[string]bool)

	for len(result) < n {
		if idx == len(r.keys) {
			idx = 0
		}

		node := r.nodes[r.keys[idx]]
		if !seen[node] {
			result = append(result, node)
			seen[node] = true
		}

		idx++
	}

	return result
}

// func main() {
// 	ring := newHashRing(5)

// 	ring.AddNode("NodeA")
// 	ring.AddNode("NodeB")
// 	ring.AddNode("NodeC")

// 	fmt.Println("Single node:", ring.GetNode("user1"))
// 	fmt.Println("Replicas:", ring.GetNodes("user1", 3))
// }
