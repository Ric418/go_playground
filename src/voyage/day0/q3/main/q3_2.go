package main

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request){
	h(w,r)
}

func taimer() http.Handler {
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Je t'aime.")
	})
}

func main() {
	http.Handle("/", taimer())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

