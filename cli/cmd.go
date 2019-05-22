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

var dishes = &cobra.Command{
	Use:   `dishes`,
	Short: `Lists all dishes for a given supplier`,
	Run:   dishesHandler,
}

func Execute() error {
	dishes.Flags().String("supplier", "", `The ID of the supplier.`)

	root.AddCommand(offers)
	root.AddCommand(dishes)

	err := root.Execute()
	return err
}
