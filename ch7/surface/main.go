package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/surface", plot)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
