package main

import "net/http"

func main() {
	http.Handle("/", errorHandler(homePage))
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) error {
	return nil
}
