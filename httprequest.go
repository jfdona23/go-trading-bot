package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// Default timeot for requests in seconds.
	reqTimeout int = 10
	// Maximum amount of bytes to read from a response when converted into a string
	maxResponseBytes int64 = 4096
)

// Retrieve the content of an URL using GET method.
func httpGet(url string) (*http.Response, error) {
	httpClient := &http.Client{Timeout: time.Duration(reqTimeout) * time.Second}
	response, err := httpClient.Get(url)
	if err != nil {
		return response, err
	}
	Log.Debug("Request URL: " + response.Request.URL.String())
	return response, nil
}

// Unmarshall a http response into a `target` structure.
func responseToJsonStruct(response *http.Response, target interface{}) error {
	defer response.Body.Close()
	err := json.NewDecoder(response.Body).Decode(target)
	if err != nil {
		return err
	}
	return nil
}

// Read a response up to `maxResponseBytes` and return its string representation.
func responseToString(response *http.Response) (string, error) {
	defer response.Body.Close()
	body, err := io.ReadAll(io.LimitReader(response.Body, maxResponseBytes))
	if err != nil {
		return "", err
	}
	Log.Debug(fmt.Sprintf("Plain text response: %+v", string(body)))
	return string(body), nil
}
