package gofer

import (
	"fmt"
	"net/http"
	"net/url"

	logger "github.com/sirupsen/logrus"
)

// 将query参数序列化为完整的url
func QueryParamDumps(urlString string, params map[string]string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Printf("Error parsing URL '%s': %v\n", urlString, err)
		return "", err
	}
	q := u.Query()

	// 将新的参数从 params map 添加到 url.Values 中
	for key, value := range params {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// 将query参数从url中提取
func QueryParamLoads(urlString string) (string, map[string]string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", nil, err
	}

	queryParams := make(map[string]string)
	values := u.Query()
	for key := range values {
		queryParams[key] = values.Get(key)
	}

	u.RawQuery = ""
	baseUrl := u.String()

	return baseUrl, queryParams, nil
}

// 发送http get请求
func Get(Url string, headers map[string]string, params map[string]string) (Response, error) {
	request, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		logger.Error("Failed to create request:", err)
		return Response{}, err
	}
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	query := request.URL.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	request.URL.RawQuery = query.Encode()
	logger.Info("Sending GET request to: ", request.URL.String())
	response, err := doRequest(request)
	return response, err
}
