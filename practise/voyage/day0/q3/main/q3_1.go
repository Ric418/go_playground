package main

import (
	"fmt"
	"log"
	"net/http"
)

type simpleHandler struct{
	msg string
}

func (h *simpleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	h.msg = "Je go'aime."
	fmt.Fprintf(w, "msg is %s\n", h.msg)
}

func main() {
	http.Handle("/", &simpleHandler{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
