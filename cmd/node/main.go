package main

import (
	"dynamo-go/internals/ring"
	"fmt"
)

func main() {
	r := ring.NewHashRing(5)

	r.AddNode("NodeA")
	r.AddNode("NodeB")

	fmt.Println("Node:", r.GetNode("user123"))
}
