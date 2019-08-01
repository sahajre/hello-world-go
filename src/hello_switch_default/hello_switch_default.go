package main

import "fmt"

func main() {
	i := 42
	switch i {
	case 0:
		fmt.Println("case 0")
	case 1:
		fmt.Println("case 1")
	default:
		fmt.Println("Hello, 世界")
	}
}
