package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type MyClient struct {
	client *http.Client
	url    string
}

func InitHttpClient(port int) MyClient {
	return MyClient{
		client: http.DefaultClient,
		url:    fmt.Sprintf("http://localhost:%d/request-info/", port),
	}
}

func (this *MyClient) SamplePost() {
	this.client.PostForm(this.url, url.Values{
		"title":  []string{"title"},
		"author": []string{"author"},
	})

	http.DefaultClient.PostForm("http://localhost:1234/request-info/", url.Values{
		"title":  []string{"title"},
		"author": []string{"author"},
	})
}

func (this *MyClient) Get(queryParams map[string]string) (string, error) {
	//http.DefaultClient.Get()
	params := []string{}
	for key, value := range queryParams {
		params = append(params, fmt.Sprintf("%v=%v", key, value))
	}
	url := this.url + "?" + strings.Join(params, "&")
	return this.call(http.MethodGet, url, nil)
}

func (this *MyClient) call(method, url string, body io.Reader) (string, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}

	response, err := this.client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bas status code: %d", response.StatusCode)
	}

	responseContent, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(responseContent), nil
}
