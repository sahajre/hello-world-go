package main

import "fmt"

func getGreetingByteSeq() func() (byte, bool) {
	i := 0
	s := "Hello, 世界"

	return func() (byte, bool) {
		i++
		return s[i-1], i == len(s)
	}
}

func main() {
	getNextByte := getGreetingByteSeq()

	s := make([]byte, 0)
	b, last := getNextByte()
	for {
		if last {
			s = append(s, b)
			break
		} else {
			s = append(s, b)
			b, last = getNextByte()
		}

	}

	fmt.Println(string(s))
}
