// package main

// import (
// 	"fmt"
// )

// func say(s string, done chan int) {
// 	fmt.Println(s)
// 	done <- 1
// }

// func main() {
// 	done := make(chan int)
// 	go say("Hello, 世界", done)
// 	<-done
// }

package main

import (
	"fmt"
)

func greet(c chan string) {
	c <- "Hello, 世界"
}

func main() {
	c := make(chan string)
	go greet(c)
	msg := <-c
	fmt.Println(msg)
}
