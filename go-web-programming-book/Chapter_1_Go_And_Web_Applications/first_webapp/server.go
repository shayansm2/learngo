package main

import (
	"fmt"
	"net/http"
)

func handlerFunction(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

type strHandler string

func (this strHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, string(this))
}

func main() {
	http.HandleFunc("/", handlerFunction)
	var a strHandler = "my name is chicky"
	http.Handle("/dev/", a)
	http.ListenAndServe(":8080", nil)
}
