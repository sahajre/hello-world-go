package main

import (
	"fmt"
	"strconv"
)

func main() {
	m := make(map[string]rune)
	m["key1"] = 'H'
	m["key2"] = 'e'
	m["key3"] = 'l'
	m["key4"] = 'l'
	m["key5"] = 'o'
	m["key6"] = ','
	m["key7"] = ' '
	m["key8"] = '世'
	m["key9"] = '界'

	for _, k := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		key := "key" + strconv.Itoa(k)
		fmt.Printf("%c", m[key])
	}

	fmt.Println()
}
