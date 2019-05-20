package main

import "github.com/spf13/cobra"

var root = &cobra.Command{
	Use:   "foodunit",
	Short: "FoodUnit CLI",
	Run:   Handler("root"),
}

var offers = &cobra.Command{
	Use:   "offers",
	Short: "List all offers",
	Run:   Handler("offers"),
}

func init() {
	root.AddCommand(offers)
}

func Exec() error {
	return root.Execute()
}
