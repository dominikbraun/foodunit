package main

import "github.com/spf13/cobra"

var root = &cobra.Command{
	Use:   "foodunit",
	Short: "FoodUnit CLI",
	Run:   Action("root"),
}

var suppliers = &cobra.Command{
	Use:   "suppliers",
	Short: "List all suppliers",
}

func init() {
	root.AddCommand(suppliers)
}

func Exec() error {
	return root.Execute()
}
