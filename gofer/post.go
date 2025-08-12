package gofer

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	logger "github.com/sirupsen/logrus"
)

/*
PostJson sends a POST request with JSON data to the specified URL.
*/
func PostJson(URL string, headers map[string]string, jsonData map[string]any) (Response, error) {

	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		logger.Error("Failed to marshal JSON data:", err)
		return Response{}, err
	}
	request, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		logger.Error("Failed to create request:", err)
		return Response{}, err
	}
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	logger.Info("Sending POST request to:", request.URL.String())
	response, err := doRequest(request)
	return response, err
}

/*
PostForm sends a POST request with form data to the specified URL.
*/
func PostForm(URL string, headers map[string]string, formData map[string]string) (Response, error) {

	values := url.Values{}
	for key, value := range formData {
		values.Set(key, value)
	}
	requestBody := bytes.NewBufferString(values.Encode())
	request, err := http.NewRequest("POST", URL, requestBody)
	if err != nil {
		logger.Error("Failed to create request:", err)
		return Response{}, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	logger.Info("Sending POST request to:", request.URL.String())
	response, err := doRequest(request)
	return response, err
}

// PostFile sends a POST request with a file to the specified URL.
func PostFile(URL string, headers map[string]string, filePath string) (Response, error) {

	byteArray, err := os.ReadFile(filePath)
	if err != nil {
		logger.Error("Failed to read file:", err)
		return Response{}, err
	}

	request, err := http.NewRequest("POST", URL, bytes.NewBuffer(byteArray))
	if err != nil {
		logger.Error("Failed to create request:", err)
		return Response{}, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}
	request.Header.Set("Content-Type", "multipart/form-data")

	logger.Info("Sending POST request to:", request.URL.String())
	response, err := doRequest(request)
	return response, err
}
