package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to this awsome site</h1>")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server...")
	http.ListenAndServe(":3000", mux)
}
