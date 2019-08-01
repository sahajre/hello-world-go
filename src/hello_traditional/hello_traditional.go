package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}

// To run: go run hello_traditional.go

/*
Every Go program starts with package clause. The package main tells the compiler
to create an executable instead of shared library.

func main() is entry level function for code execution.

The file defining func main() must have package main declaration.

Use import for using in-built or third party packages. In this case, in-built
text formatting package is used for printing string "Hello, world" followed by a
new line.

Go source code is UTF-8 encoded file. i.e. You can use unicode string without
any extra support for encoding, decoding. To demonstrate this, I will keep using
text "world" translated in Chinese as '世界". In further examples, this will
also help to understand difference between string literals vs rune type.
*/
