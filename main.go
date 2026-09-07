package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", HandleHelloWorld)
	mux.HandleFunc("/", Handle404)

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		fmt.Println("Server error:", err)
	}

}

func HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
