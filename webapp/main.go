package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8000", nil)
}

func handlerFunc(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>Welcome to my site!</h1>")
}
