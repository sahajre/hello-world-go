package main

import "fmt"

func printStr(s string) {
	fmt.Print(s)
}

func printHelloWorld() {
	defer printStr("界")
	defer printStr("世")
	defer printStr(", ")
	printStr("Hello")
}
func main() {
	printHelloWorld()
	printStr("\n")
}
