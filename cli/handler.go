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

// api provides basic methods for accessing any public API and sending
// requests to it. It also stores the last request that was sent and
// the corresponding response it received.
type api struct {
	url     string
	lastReq *http.Request
	lastRes *http.Response
	client  *http.Client
}

// request sends a HTTP request to any API endpoint. The resulting HTTP
// response is stored in the lastRes field. However, it is also returned
// by the method - next to an error if the request failed.
func (a *api) request(rsrc string) (*http.Response, error) {
	reqUrl := fmt.Sprintf("%s%s", a.url, rsrc)
	var err error

	if a.lastReq, err = http.NewRequest(DefMethod, reqUrl, nil); err != nil {
		return nil, err
	}
	a.lastRes, err = a.client.Do(a.lastReq)

	return a.lastRes, err
}

// lastResBytes returns the body of the last response that was received
// as a byte array. This is useful to unmarshal the body. Note that the
// last response might be nil if the preceding request failed.
func (a *api) lastResBytes() ([]byte, error) {
	body := a.lastRes.Body
	b, err := ioutil.ReadAll(body)

	return b, err
}

// StdApi provides a precasted api instance for the API URL that is
// configured in conf.go.
var StdApi *api

// Processes the root command.
var rootHandler = func(cmd *cobra.Command, args []string) {}

// Processes the offer command.
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

// Processes the menu command.
var menuHandler = func(cmd *cobra.Command, args []string) {

	supplier := flagOrExit(cmd, "supplier", `No supplier given.`)
	rsrc := "/menu/" + supplier

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

// Processes the supplier command.
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

// flagOrExit can be used to obtain a certain flag value of any command.
// If the flag was not provided, the program exits (all flags are set
// to "" by default).
func flagOrExit(cmd *cobra.Command, name string, msg string) string {
	desired := cmd.Flag(name)
	val := desired.Value.String()

	if val == "" {
		log.Println(msg)
		os.Exit(1)
	}
	return val
}

// Initializes the StdApi variable by reading the configured URL and
// creating a new instance of http.Client.
func init() {
	StdApi = &api{
		url:     ApiUrl,
		lastReq: nil,
		lastRes: nil,
		client:  &http.Client{},
	}
}
