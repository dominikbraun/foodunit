package main

import "net/http"

type Api struct {
	url    string
	client *http.Client
}

func (a *Api) req(endpoint string) *http.Response {
	return &http.Response{}
}
