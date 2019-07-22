package main

import "fmt"

func main() {
	m := make(chan string, 2)

	m <- "Hello, "
	m <- "世界"

	fmt.Print(<-m)
	fmt.Println(<-m)
}
