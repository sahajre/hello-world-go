package main

import "fmt"

func main() {
	i := 42
	switch i {
	case 1:
		fmt.Print("case 1")
		fallthrough
	case 42:
		fmt.Print("Hello, ")
		fallthrough
	case 84:
		fmt.Print("世")
		fallthrough
	case 404:
		fmt.Println("界")
	case 440:
		fmt.Println("case 440")
	default:
		fmt.Println("default")
	}
}
