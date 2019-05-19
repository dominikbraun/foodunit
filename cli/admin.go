package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type handler func(cmd *cobra.Command, args []string)

func main() {
	var root = &cobra.Command{
		Use:   "foodunit",
		Short: "FoodUnit CLI",
		Run:   exec("root"),
	}
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func exec(cmd string) handler {
	handlers := map[string]handler{
		"root": func(cmd *cobra.Command, args []string) {},
	}
	return handlers[cmd]
}
