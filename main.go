package main

import (
	"io"
	"log"
	"net/http"
)

func function(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/function/", function)

	log.Fatal(http.ListenAndServe(":80", nil))
}
