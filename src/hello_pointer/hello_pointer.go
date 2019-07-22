package main

import "fmt"

func main() {
	s := ""
	populateHelloWorld(&s)
	fmt.Println(s)
}

func populateHelloWorld(sptr *string) {
	*sptr = "Hello, 世界"
}
