package main

import (
	"github.com/bwmarrin/snowflake"
	"fmt"
)

func main() {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a snowflake ID.
	m := make(map[string]int)
	for i := 0; i < 1000000; i++ {
		id := node.Generate()
		m[id.String()] = i
	}
	fmt.Println(len(m))

}
