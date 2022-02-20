package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Simple Service v1.0")

	// http server on port 8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Simple Service v1.0"))
	})
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
