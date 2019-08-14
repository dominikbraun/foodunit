// Package main provides the main function for starting the API server.
package main

import (
	"github.com/dominikbraun/foodunit/server"
)

// main delegates all listening/serving tasks to the server package.
func main() {
	server.New(nil, server.Resume).Start()
}
