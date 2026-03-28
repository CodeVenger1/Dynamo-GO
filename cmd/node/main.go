package main

import (
	"dynamo-go/internal/ring"
	"fmt"
)

func main() {
	r := ring.NewHashRing(5)

	r.AddNode("NodeA")
	r.AddNode("NodeB")

	fmt.Println("Node:", r.GetNode("user123"))
}
