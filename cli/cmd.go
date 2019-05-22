package main

import "github.com/spf13/cobra"

var root = &cobra.Command{
	Use:   `foodunit`,
	Short: `FoodUnit CLI`,
	Run:   rootHandler,
}

var offers = &cobra.Command{
	Use:   `offers`,
	Short: `Lists all active offers`,
	Run:   offerHandler,
}

func Execute() error {
	root.AddCommand(offers)
	err := root.Execute()
	return err
}
