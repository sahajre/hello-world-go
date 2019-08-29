package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

const input = `Hello
, 世界
`

func main() {
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		fmt.Print(s.Text())
	}
	fmt.Println()
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
