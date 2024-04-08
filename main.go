package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
