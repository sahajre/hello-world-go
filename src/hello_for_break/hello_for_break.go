package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i == 42 {
			fmt.Println("Hello, 世界")
			break
		}
	}
}
