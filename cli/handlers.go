package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

type F func(cmd *cobra.Command, args []string)

func Handler(cmd string) F {
	mappings := map[string]F{
		"root":   rootHandler,
		"offers": offersHandler,
	}
	return mappings[cmd]
}

var rootHandler = func(cmd *cobra.Command, args []string) {
	fmt.Println("FoodUnit CLI")
}

var offersHandler = func(cmd *cobra.Command, args []string) {
	api := StdApi
	_, err := api.req("/offers")

	if err != nil {
		fmt.Println(err)
	}
}
