package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	gohttpclient "github.com/bozd4g/go-http-client"
)

const (
	port = "4001"
	path = "http://localhost:4001"
)

var serverSingleton *Server

func getServer() *Server {
	if serverSingleton == nil {
		serverSingleton = NewServer(port)
		go serverSingleton.Start()
		time.Sleep(1000 * time.Millisecond)
	}
	return serverSingleton
}

type responseForm struct {
	Result string `json:"Result"`
	Error  string `json:"Error"`
}

type testBook struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func addBook(title, author string, duplicate bool) (string, error) {
	data := url.Values{
		"title":  []string{title},
		"author": []string{author},
	}

	if title == "" && author == "" {
		data = nil
	}

	resp, err := http.DefaultClient.PostForm(path+"/book", data)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusBadRequest {
		return "", fmt.Errorf("status code not OK: %d", resp.StatusCode)
	}

	var rf responseForm
	err = json.NewDecoder(resp.Body).Decode(&rf)
	if err != nil {
		return "", err
	}

	if rf.Error != "" {
		return "", errors.New(rf.Error)
	}

	if rf.Error != "" && resp.StatusCode != http.StatusBadRequest {
		return "", fmt.Errorf("status code %d not OK for error %s", resp.StatusCode, rf.Error)
	}

	if rf.Result != fmt.Sprintf("added book %s by %s", strings.ToLower(title), strings.ToLower(author)) && !duplicate {
		return "", fmt.Errorf("result message is incorrect: %s", rf.Result)
	}

	if rf.Result == "this Book is already in the library" && duplicate {
		return "", errors.New(rf.Result)
	}

	return rf.Result, nil
}

func dummyAddBook(title, author string, duplicate bool) (string, error) {
	data := url.Values{
		"fake_key": []string{title},
		"author":   []string{author},
	}

	resp, err := http.DefaultClient.PostForm(path+"/book", data)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusBadRequest {
		return "", fmt.Errorf("status code not OK: %d", resp.StatusCode)
	}

	var rf responseForm
	err = json.NewDecoder(resp.Body).Decode(&rf)
	if err != nil {
		return "", err
	}

	if !isResponseCodeValid(rf, resp.StatusCode) {
		return "", fmt.Errorf("status code not OK: %d", resp.StatusCode)
	}

	if rf.Error != "" {
		return "", errors.New(rf.Error)
	}

	if rf.Result != fmt.Sprintf("added book %s by %s", strings.ToLower(title), strings.ToLower(author)) && !duplicate {
		return "", fmt.Errorf("result message is incorrect: %s", rf.Result)
	}

	if rf.Result == "this Book is already in the library" && duplicate {
		return "", errors.New(rf.Result)
	}

	return rf.Result, nil
}

func getBook(title, author string) (*testBook, error) {
	route := fmt.Sprintf("%s/book", path)
	if title != "" {
		route += fmt.Sprintf("?title=%s", url.QueryEscape(title))
	}
	if author != "" {
		symbol := "&"
		if title == "" {
			symbol = "?"
		}
		route += fmt.Sprintf("%sauthor=%s", symbol, url.QueryEscape(author))
	}

	resp, err := http.DefaultClient.Get(route)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusBadRequest {
		return nil, fmt.Errorf("invalid status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var successfulResult testBook
	var unsuccessfulResult responseForm
	err1 := json.Unmarshal(body, &successfulResult)
	err2 := json.Unmarshal(body, &unsuccessfulResult)
	if successfulResult.Title != "" {
		return &successfulResult, nil
	}

	if unsuccessfulResult.Error != "" {
		if !isResponseCodeValid(unsuccessfulResult, resp.StatusCode) {
			return nil, fmt.Errorf("status code %d not OK for error %s", resp.StatusCode, unsuccessfulResult.Error)
		}
		return nil, errors.New(unsuccessfulResult.Error)
	}

	return nil, errors.Join(err1, err2)
}

func reserveBook(title, author string, borrow ...bool) (string, error) {
	route := fmt.Sprintf("%s/book", path)
	if title != "" {
		route += fmt.Sprintf("?title=%s", url.QueryEscape(title))
	}
	if author != "" {
		symbol := "&"
		if title == "" {
			symbol = "?"
		}
		route += fmt.Sprintf("%sauthor=%s", symbol, url.QueryEscape(author))
	}

	var option gohttpclient.Option
	if len(borrow) == 1 {
		body := map[string]bool{
			"borrow": borrow[0],
		}

		jsonBody, err := json.Marshal(body)
		if err != nil {
			return "", errors.Join(err, errors.New("something is wrong with encoding request body"))
		}

		option = gohttpclient.WithBody(jsonBody)
	} else {
		option = gohttpclient.WithBody(nil)
	}

	client := gohttpclient.New(route)
	response, err := client.Put(context.Background(), "", option)
	if err != nil {
		return "", errors.Join(err, errors.New("something went wrong when sending a PUT request"))
	}

	return handleResponse(response)
}

func handleResponse(response *gohttpclient.Response) (string, error) {
	resp := response.Get()
	var rf responseForm

	err := json.Unmarshal(response.Body(), &rf)
	if err != nil {
		return "", errors.Join(err, errors.New("something wrong with decoding response body"))
	}

	if !isResponseCodeValid(rf, resp.StatusCode) {
		return "", fmt.Errorf("status code not OK: %d", resp.StatusCode)
	}

	if rf.Error != "" && resp.StatusCode != http.StatusBadRequest {
		return "", fmt.Errorf("status code %d not OK for error %s", resp.StatusCode, rf.Error)
	}

	if rf.Error != "" {
		return "", errors.New(rf.Error)
	}

	return rf.Result, nil
}

func delBook(title, author string) (string, error) {
	title = url.QueryEscape(title)
	author = url.QueryEscape(author)

	client := gohttpclient.New(fmt.Sprintf("%s/book?title=%s&author=%s", path, title, author))
	response, err := client.Delete(context.Background(), "")
	if err != nil {
		return "", err
	}

	return handleResponse(response)
}

func isResponseCodeValid(response responseForm, code int) bool {
	if code != http.StatusOK && code != http.StatusBadRequest {
		fmt.Println("status code is:", code)
		return false
	}

	if response.Error != "" && code == http.StatusBadRequest {
		return true
	}

	if response.Result != "" && code == http.StatusOK {
		return true
	}

	fmt.Println("resp is:", response)
	fmt.Println("status code is:", code)
	return false
}