package gofer

import (
	"io"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

type Response struct {
	Code int    `json:"code"`
	Body string `json:"body"`
}

func doRequest(request *http.Request) (Response, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logger.Error("Failed to send request:", err)
		return Response{}, err
	}
	defer response.Body.Close()
	bodyResponse, err := io.ReadAll(response.Body)

	return Response{
		Code: response.StatusCode,
		Body: string(bodyResponse),
	}, err
}
