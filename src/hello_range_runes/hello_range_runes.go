package main

import (
	"fmt"
)

func main() {
	out := "Hello, 世界"
	for _, ch := range out {
		fmt.Printf("%c", ch)
	}
	fmt.Println()
}
