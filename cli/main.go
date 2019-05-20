package main

import (
	"fmt"
	"os"
)

func main() {
	if err := Exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
