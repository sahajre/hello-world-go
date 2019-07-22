package main

import (
	"encoding/json"
	"fmt"
)

type greet struct {
	Output string
}

func main() {
	g := &greet{Output: "Hello, 世界"}
	b, err := json.Marshal(g)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
