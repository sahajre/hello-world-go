package main

import "fmt"

func main() {
	s := make([]rune, 3)
	s[0] = 'H'
	s[1] = 'e'
	s[2] = 'l'

	s = append(s, 'l')
	s = append(s, 'l', 'o', ',', ' ', '世', '界')
	fmt.Println(string(s))
}
