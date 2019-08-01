package main

import "fmt"

func main() {
	i := 42
	switch i {
	case 0:
		fmt.Print("case 0")
	case 42:
		fmt.Println("Hello, 世界")
	default:
		fmt.Println("default")
	}
}
