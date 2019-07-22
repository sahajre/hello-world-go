package main

import "fmt"

func main() {
	var s = [9]rune{
		'H', 'e', 'l', 'l', 'o', ',', ' ', '世', '界',
	}

	for _, ch := range s {
		fmt.Printf("%c", ch)
	}
	fmt.Println()
}
