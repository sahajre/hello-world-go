package main

import "fmt"

func init() {
	fmt.Print(hello)
}

func init() {
	fmt.Println(world)
}

var (
	hello = "Hello, "
	world = "世界"
)

func main() {

}
