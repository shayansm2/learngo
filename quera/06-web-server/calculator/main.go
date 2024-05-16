package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{port}
}

func (s *Server) Start() {
	http.Handle("/add", responseHandler(addHandler))
	http.Handle("/sub", responseHandler(subHandler))
	err := http.ListenAndServe(fmt.Sprintf(":%v", s.port), nil)
	if err != nil {
		fmt.Println(err)
	}
}

type response struct {
	Result string `json:"result"`
	Err    string `json:"error"`
}

type responseHandler func(r *http.Request) (string, error)

func (f responseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resut, err := f(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{Err: fmt.Sprint(err)})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response{Result: resut})
	}
}

func addHandler(r *http.Request) (string, error) {
	var result int64
	numbers, err := getNumbersFromRequest(r)
	if err != nil {
		return "", err
	}
	for _, num := range numbers {
		if checkAddOverflow(result, int64(num)) {
			return "", errors.New("Overflow")
		}
		result += int64(num)
	}
	return fmt.Sprintf("The result of your query is: %d", result), nil
}

func subHandler(r *http.Request) (string, error) {
	numbers, err := getNumbersFromRequest(r)
	if err != nil {
		return "", err
	}
	result := int64(numbers[0])
	for i := 1; i < len(numbers); i++ {
		if checkAddOverflow(result, -int64(numbers[i])) {
			return "", errors.New("Overflow")
		}
		result -= int64(numbers[i])
	}
	return fmt.Sprintf("The result of your query is: %d", result), nil
}

func getNumbersFromRequest(r *http.Request) ([]int, error) {
	queryParams, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return nil, err
	}

	if _, found := queryParams["numbers"]; !found {
		return nil, errors.New("'numbers' parameter missing")
	}

	result := make([]int, 0)
	for _, val := range strings.Split(queryParams["numbers"][0], ",") {
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil
}

func checkAddOverflow(a, b int64) bool {
	sum := a + b
	if ((sum - b) != a) || ((sum < a) != (b < 0)) {
		return true // Overflow occurred
	}
	return false // No overflow
}

func main() {
	s := NewServer("8369")
	s.Start()
}
