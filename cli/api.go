package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type Api struct {
	url    string
	client *http.Client
}

func (a *Api) req(endpoint string) (*http.Response, error) {
	if a.client == nil {
		return nil, errors.New("API client not set")
	}
	reqUrl := fmt.Sprintf("%s%s", a.url, endpoint)
	req, err := http.NewRequest("GET", reqUrl, nil)

	if err != nil {
		return nil, err
	}
	return a.client.Do(req)
}

func (a *Api) reqBody(res *http.Response) io.ReadCloser {
	return res.Body
}
