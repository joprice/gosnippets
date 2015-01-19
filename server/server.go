package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	// pass directly to log in case error occurs
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
}
