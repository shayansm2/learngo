package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Server struct {
	port int
}

func (s *Server) Start() {
	http.HandleFunc("/request-info/", requestInfoHandler)
	http.HandleFunc("/json-response/", jsonResponseHandler)
	http.ListenAndServe(fmt.Sprintf(":%v", s.port), nil)
}

func requestInfoHandler(w http.ResponseWriter, r *http.Request) {
	rawResponse := prepareResponse(r)
	fmt.Fprintf(w, strings.Join(rawResponse, "\n"))
	fmt.Print(strings.Join(rawResponse, "\n"))
}

func jsonResponseHandler(w http.ResponseWriter, r *http.Request) {

}

func prepareResponse(r *http.Request) []string {
	result := []string{}

	result = append(result, fmt.Sprintf("the request \"method\" is %v", r.Method))
	result = append(result, fmt.Sprintf("the request \"host\" is %v", r.Host))
	result = append(result, fmt.Sprintf("the request \"URL\" is %v", r.URL))

	queryParams, _ := url.ParseQuery(r.URL.RawQuery)
	result = append(result, fmt.Sprintf("\"query parameters\" are %v", queryParams))

	r.ParseForm()
	result = append(result, fmt.Sprintf("\"form\" is %v", r.Form))
	result = append(result, fmt.Sprintf("\"post form\" is %v", r.PostForm))
	result = append(result, fmt.Sprintf("\"multipart form values\" is %v", r.MultipartForm))

	stringBody, _ := io.ReadAll(r.Body)
	result = append(result, fmt.Sprintf("\"string body\" is %v", string(stringBody)))

	var jsonBody map[string]any
	json.NewDecoder(r.Body).Decode(&jsonBody)
	result = append(result, fmt.Sprintf("\"json body\" is %v", jsonBody))

	result = append(result, fmt.Sprintf("\"header\" is %v", r.Header))

	return result
}
