package main

import "fmt"

func init() {
	fmt.Print(hello)
}

var hello = "Hello, "

func main() {
	fmt.Println()
}

/*
To Run: go run hello_init.go world_init.go
*/
