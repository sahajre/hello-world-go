package main

import (
	"fmt"
	"log"
)

func helloWorld(s string) (string, error) {
	return s, nil
}

func main() {
	s, err := helloWorld("Hello, 世界")
	if err != nil {
		log.Fatal("could not get hello greeting")
	}
	fmt.Println(s)
}
