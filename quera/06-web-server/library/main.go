package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Server struct {
	port          string
	borrowedBooks map[Book]bool
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type response struct {
	Result string `json:"Result"`
	Error  string `json:"Error"`
}

func NewServer(port string) *Server {
	return &Server{
		port:          port,
		borrowedBooks: make(map[Book]bool),
	}
}

func (s *Server) Start() {
	http.HandleFunc("/book", s.bookHandler)
	http.ListenAndServe(fmt.Sprintf(":%v", s.port), nil)
}

func (s *Server) bookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		book, err := s.getBookHandler(r)
		getResponseHandler(w, book, err)
		return
	}

	var logicHandler func(r *http.Request) (string, error)

	switch r.Method {
	case http.MethodPost:
		logicHandler = s.PostBookHandler
	case http.MethodPut:
		logicHandler = s.PutBookHandler
	case http.MethodDelete:
		logicHandler = s.delBookHandler
	default:
	}

	result, err := logicHandler(r)
	defaultResponseHandler(w, result, err)
}

func defaultResponseHandler(w http.ResponseWriter, message string, err error) {
	if err == nil && message == "" {
		return
	}

	resp := response{}

	if message != "" {
		resp.Result = message
	}

	if err != nil {
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(resp)
}

func getResponseHandler(w http.ResponseWriter, book *Book, err error) {
	if book != nil {
		json.NewEncoder(w).Encode(*book)
		return
	}

	resp := response{
		Result: "",
		Error:  err.Error(),
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) getBookHandler(r *http.Request) (result *Book, err error) {
	result, err = s.getBookFromQueryParams(r)
	if err != nil {
		return nil, err
	}

	borrowed, found := s.borrowedBooks[*result]

	if !found {
		return nil, errors.New("this Book does not exist")
	}

	if borrowed {
		return nil, errors.New("this Book is borrowed")
	}

	return result, nil
}

func (s *Server) getBookFromQueryParams(r *http.Request) (*Book, error) {
	queryParameters := r.URL.Query()

	title := strings.ToLower(queryParameters.Get("title"))
	author := strings.ToLower(queryParameters.Get("author"))

	if title == "" || author == "" {
		return nil, errors.New("title or author cannot be empty")
	}

	return &Book{Title: title, Author: author}, nil
}

type postRequestBody struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (s *Server) PostBookHandler(r *http.Request) (result string, err error) {
	var body *postRequestBody
	switch r.Header.Get("Content-Type") {
	case "application/x-www-form-urlencoded":
		body, err = retrieveWWWFromBodyPost(r)
	case "application/json":
	default:
		body, err = retrieveJsonBodyPost(r)
	}

	if err != nil {
		return
	}

	body.Title = strings.ToLower(strings.TrimSpace(body.Title))
	body.Author = strings.ToLower(strings.TrimSpace(body.Author))

	if body.Title == "" || body.Author == "" {
		err = errors.New("title or author cannot be empty")
		return
	}

	book := &Book{
		Title:  body.Title,
		Author: body.Author,
	}

	if _, found := s.borrowedBooks[*book]; found {
		result = "this Book is already in the library"
		return
	}

	s.borrowedBooks[*book] = false

	result = fmt.Sprintf("added book %s by %s", book.Title, book.Author)
	return
}

func (s *Server) PutBookHandler(r *http.Request) (result string, err error) {
	book, err := s.getBookFromQueryParams(r)
	if err != nil {
		return "", err
	}

	if _, found := s.borrowedBooks[*book]; !found {
		err = errors.New("book not found")
		return
	}

	var borrow bool
	switch r.Header.Get("Content-Type") {
	case "application/x-www-form-urlencoded":
		borrow, err = retrieveWWWFromBodyPut(r)
	case "application/json":
	default:
		borrow, err = retrieveJsonBodyPut(r)
	}

	if err != nil {
		return "", errors.New("borrow value cannot be empty")
	}

	if s.borrowedBooks[*book] == borrow {
		var errMessage string
		if borrow {
			errMessage = "this Book is already borrowed"
		} else {
			errMessage = "this Book is already in the library"
		}
		return "", errors.New(errMessage)
	}
	s.borrowedBooks[*book] = borrow

	if borrow {
		result = "you have borrowed this Book successfully"
	} else {
		result = "thank you for returning this Book"
	}
	return result, nil
}

func (s *Server) delBookHandler(r *http.Request) (result string, err error) {
	book, err := s.getBookFromQueryParams(r)
	if err != nil {
		return "", err
	}

	if _, found := s.borrowedBooks[*book]; !found {
		return "", errors.New("this Book does not exist")
	}

	delete(s.borrowedBooks, *book)

	return "successfully deleted", nil
}

func retrieveWWWFromBodyPost(r *http.Request) (*postRequestBody, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}
	return &postRequestBody{
		Title:  r.Form.Get("title"),
		Author: r.Form.Get("author"),
	}, nil
}

func retrieveWWWFromBodyPut(r *http.Request) (bool, error) {
	err := r.ParseForm()
	if err != nil {
		return false, err
	}
	return r.Form.Get("borrow") == "true", nil
}

func retrieveJsonBodyPost(r *http.Request) (*postRequestBody, error) {
	var body postRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		err = errors.Join(err, errors.New("the request body is invalid"))
		return nil, err
	}
	return &body, nil
}

func retrieveJsonBodyPut(r *http.Request) (bool, error) {
	var body map[string]bool
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		err = errors.Join(err, errors.New("the request body is invalid"))
		return false, err
	}

	if _, found := body["borrow"]; !found {
		return false, errors.New("\"borrow\" key was not in the payload")
	}

	return body["borrow"], nil
}

//func main() {
//	s := NewServer("9642")
//	s.Start()
//}
