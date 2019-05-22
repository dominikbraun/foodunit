package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type api struct {
	url     string
	lastReq *http.Request
	lastRes *http.Response
	client  *http.Client
}

func (a *api) request(rsrc string) (*http.Response, error) {
	reqUrl := fmt.Sprintf("%s%s", a.url, rsrc)
	var err error

	if a.lastReq, err = http.NewRequest(DefMethod, reqUrl, nil); err != nil {
		return nil, err
	}
	a.lastRes, err = a.client.Do(a.lastReq)

	return a.lastRes, err
}

func (a *api) lastResBytes() ([]byte, error) {
	body := a.lastRes.Body
	b, err := ioutil.ReadAll(body)

	return b, err
}

var StdApi *api

var rootHandler = func(cmd *cobra.Command, args []string) {}

var offerHandler = func(cmd *cobra.Command, args []string) {
	rsrc := "/offers"
	_, err := StdApi.request(rsrc)

	if err != nil {
		log.Println(err)
	}
	body, _ := StdApi.lastResBytes()
	var offers []Offer

	if err := json.Unmarshal(body, &offers); err != nil {
		log.Println(err)
	}
	for _, o := range offers {
		fmt.Println(o)
	}
}

var dishesHandler = func(cmd *cobra.Command, args []string) {

	supplier := flagOrExit(cmd, "supplier", `No supplier given.`)
	rsrc := "/dishes/" + supplier

	_, err := StdApi.request(rsrc)

	if err != nil {
		log.Println(err)
	}
	body, _ := StdApi.lastResBytes()
	var cats []Cat

	if err := json.Unmarshal(body, &cats); err != nil {
		log.Println(err)
	}
	for _, c := range cats {
		fmt.Println(c)
	}
}

var supplierHandler = func(cmd *cobra.Command, args []string) {

	supplier := flagOrExit(cmd, "supplier", `No supplier given.`)
	rsrc := "/supplier/" + supplier

	_, err := StdApi.request(rsrc)

	if err != nil {
		log.Println(err)
	}
	body, _ := StdApi.lastResBytes()
	var sup Supplier

	if err := json.Unmarshal(body, &sup); err != nil {
		log.Println(err)
	}
	fmt.Println(sup)
}

func flagOrExit(cmd *cobra.Command, name string, msg string) string {
	desired := cmd.Flag(name)
	val := desired.Value.String()

	if val == "" {
		log.Println(msg)
		os.Exit(1)
	}
	return val
}

func init() {
	StdApi = &api{
		url:     ApiUrl,
		lastReq: nil,
		lastRes: nil,
		client:  &http.Client{},
	}
}
