package utils

import (
	"bytes"
	"errors"
	"git.qutoutiao.net/todoList/app/components/logger"
	"io/ioutil"
	"net/http"
)

type RequestErr struct {
	Err   error       `json:"err"`
	Uri   string      `json:"uri"`
	Param interface{} `json:"param"`
	Data  interface{} `json:"data"`
}

//GET METHOD
func Get(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	// new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.AppLogger.Error("new http error : %v", err)
		return nil, errors.New("new request is fail ")
	}

	// 拼接Url请求参数
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	// 拼接header参数
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	client := &http.Client{}
	// add log
	logger.AppLogger.WithField("RequestURL", req.URL.String()).Infoln("Begin to GET Request URL.")
	return client.Do(req)
}

//POST  method
func Post(url string, requestBody string, params map[string]string, headers map[string]string) (*http.Response, error) {
	var (
		req      *http.Request
		err      error
	)

	req, err = http.NewRequest("POST", url, bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		logger.AppLogger.Error("new post request error: %v\n", err)
		return nil, errors.New("new request is fail: %v \n")
	}
	req.Header.Set("Content-type", "application/json")

	// 拼接Url请求参数
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	// 拼接header参数
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	client := &http.Client{}
	// add log
	logger.AppLogger.WithField("RequestURL", req.URL.String()).Infoln("Begin to POST Request URL.")
	return client.Do(req)
}

func GetResponseBody(res *http.Response) []byte {
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}
