package main

import "fmt"

func main() {
	var s = [9]rune{
		'H', 'e', 'l', 'l', 'o', ',', ' ', '世', '界',
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
	fmt.Println()
}
