package main

import (
	"fmt"
	"log"
	"net/http"
)

type helloHandler struct{}

func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world\n")
}

func main() {
	fmt.Println("Listening on port :8080")
	err := http.ListenAndServe(":8080", helloHandler{})
	log.Fatal(err)
}
