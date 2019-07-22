package main

import "fmt"

type greet struct {
	greeting string
	to       string
}

func main() {
	g := greet{greeting: "Hello", to: "世界"}
	fmt.Println(g.greeting + ", " + g.to)
}
