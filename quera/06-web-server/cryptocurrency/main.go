package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ResponseBody struct {
	Status string                       `json:"status"`
	Stats  map[string]map[string]string `json:"stats"`
}

func GetExchangeRate(source, destination string) (string, error) {
	if source == "" {
		return "", nil
	}
	source = strings.ToLower(source)

	if destination == "" {
		destination = "rls"
	}
	destination = strings.ToLower(destination)

	url := fmt.Sprintf("http://localhost:4001/rates?srcCurrency=%s&dstCurrency=%s", source, destination)
	rsp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	var responseBody ResponseBody
	err = json.NewDecoder(rsp.Body).Decode(&responseBody)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	if responseBody.Status != "OK" {
		return "", nil
	}

	key := fmt.Sprintf("%s-%s", source, destination)
	if _, found := responseBody.Stats[key]; !found {
		return "", nil
	}

	return responseBody.Stats[key]["latest"], nil
}
