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
	// Maximum amount of bytes to read from a response.
	maxResponseBytes int64 = 4096
)

// Retrieve the content of an URL using GET method.
func HttpGet(url string) (*http.Response, error) {
	httpClient := &http.Client{Timeout: time.Duration(reqTimeout) * time.Second}
	response, err := httpClient.Get(url)
	if err != nil {
		return response, err
	}
	Log.Debug("Requested URL: " + response.Request.URL.String())
	return response, nil
}

// Read a response up to `maxResponseBytes` and return its JSON representation unmashalled into a struct.
func ResponseToJsonStruct(response *http.Response, target interface{}) error {
	defer response.Body.Close()
	body, err := io.ReadAll(io.LimitReader(response.Body, maxResponseBytes))
	if err != nil {
		return err
	}
	Log.Debug("Received response: " + string(body))
	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}
	Log.Debug("Unmarshalled JSON response: " + fmt.Sprint(target))
	return nil
}

// Read a response up to `maxResponseBytes` and return its string representation.
func ResponseToString(response *http.Response) (string, error) {
	defer response.Body.Close()
	body, err := io.ReadAll(io.LimitReader(response.Body, maxResponseBytes))
	if err != nil {
		return "", err
	}
	Log.Debug("Received response: " + string(body))
	return string(body), nil
}
