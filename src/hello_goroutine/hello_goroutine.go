package main

import (
	"fmt"
)

func greet(c chan string) {
	c <- "Hello, 世界"
}

func main() {
	c := make(chan string)
	go greet(c)
	msg := <-c
	fmt.Println(msg)
}
