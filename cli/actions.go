package main

import "github.com/spf13/cobra"

type F func(cmd *cobra.Command, args []string)

func Action(cmd string) F {
	handlers := map[string]F{
		"root": func(cmd *cobra.Command, args []string) {},
	}
	return handlers[cmd]
}
