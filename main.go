package main

import (
	"log"

	"github.com/dominikbraun/foodunit/server"
)

func main() {
	log.Fatal(server.Start())
}
