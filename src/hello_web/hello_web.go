package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	printInstructions()
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, 世界")
}

func printInstructions() {
	fmt.Println(`Visit http://localhost:8080 in the web browser to see "Hello, 世界"`)
	fmt.Println(`After that Press control+c to stop the server`)
}
