package main

import "fmt"

type greeter interface {
	greet() string
}

type gopher struct{}

func (gopher) greet() string {
	return "Hello, 世界"
}

func main() {
	var g greeter
	g = gopher{}
	fmt.Println(g.greet())
}
