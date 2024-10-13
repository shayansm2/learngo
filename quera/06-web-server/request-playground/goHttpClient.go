package main

import (
	"context"
	"fmt"
	gohttpclient "github.com/bozd4g/go-http-client"
	"strings"
)

type SampleClient struct {
	url string
}

func InitGoHttpClient(port int) SampleClient {
	return SampleClient{
		url: fmt.Sprintf("http://localhost:%d/request-info/", port),
	}
}

func (this *SampleClient) GetRequestInfo(queryParams map[string]string) {
	params := []string{}
	for key, value := range queryParams {
		params = append(params, fmt.Sprintf("%v=%v", key, value))
	}
	url := this.url + "?" + strings.Join(params, "&")
	client := gohttpclient.New(url)

	_, err := client.Get(context.Background(), "")
	if err != nil {
		fmt.Println(err)
		return
	}
}
