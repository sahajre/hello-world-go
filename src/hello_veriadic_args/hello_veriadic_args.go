package main

import "fmt"

func main() {
	greeting("Hello", ", ", "世", "界")
}

func greeting(h string, to ...string) {
	fmt.Print(h)
	for _, s := range to {
		fmt.Print(s)
	}
	fmt.Println()
}
