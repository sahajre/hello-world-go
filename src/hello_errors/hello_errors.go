package main

import (
	"errors"
	"fmt"
)

func createError() error {
	return errors.New("Hello, 世界")
}

func main() {
	err := createError()
	if err != nil {
		fmt.Println(err)
	}
}
