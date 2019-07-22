package main

import "fmt"

type greet struct {
	msg string
}

type gopher struct {
	greet
}

func (g greet) sayHelloWorld() {
	fmt.Println("Hello, 世界")
}

func main() {
	g := gopher{greet: greet{msg: "hello"}}
	g.sayHelloWorld()
}
