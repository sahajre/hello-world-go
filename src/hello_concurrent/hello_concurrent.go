package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello, 世界"
	}()

	fmt.Println(<-ch)
}
