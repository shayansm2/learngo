package main

import "net/http"

type errorHandler func(http.ResponseWriter, *http.Request) error

func (f errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := f(w, r)
	if err == nil {
		return
	}

	http.NotFound(w, r)
	return
}
