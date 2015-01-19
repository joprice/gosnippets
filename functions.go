package gosnippets

import (
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", errorHandler(handler))
}

func errorHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("handling %q: %v", r.RequestURI, err)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) error {
	return nil
}
