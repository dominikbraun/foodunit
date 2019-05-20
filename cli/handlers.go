package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

type F func(cmd *cobra.Command, args []string)

func Handler(cmd string) F {
	mappings := map[string]F{
		"root":      rootHandler,
		"suppliers": supplierHandler,
	}
	return mappings[cmd]
}

var rootHandler = func(cmd *cobra.Command, args []string) {
	fmt.Println("FoodUnit CLI")
}

var supplierHandler = func(cmd *cobra.Command, args []string) {
	fmt.Println("All Suppliers")
}
