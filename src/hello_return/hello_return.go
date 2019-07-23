package main

import "fmt"

func getHelloWorld() string {
	return "Hello, 世界"
}

func main() {
	s := getHelloWorld()
	fmt.Println(s)
}
