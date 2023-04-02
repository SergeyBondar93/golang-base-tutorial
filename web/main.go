package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-custom-header-kek", "INDEX PAGE")
	fmt.Fprintf(w, "<h1>INDEX PAGE</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-custom-header-kek", "ABOUT PAGE")
	fmt.Fprintf(w, "<h1>ABOUT PAGE</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Server Starting on ", "http://localhost:3000")

	http.ListenAndServe(":3000", nil)
}
