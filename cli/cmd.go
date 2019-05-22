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

var supplier = &cobra.Command{
	Use:   `supplier`,
	Short: `Displays basic data for a given supplier`,
	Run:   supplierHandler,
}

func Execute() error {
	dishes.Flags().String("supplier", "", `The ID of the supplier`)
	supplier.Flags().String("supplier", "", `The ID of the supplier`)

	root.AddCommand(offers)
	root.AddCommand(dishes)
	root.AddCommand(supplier)

	err := root.Execute()
	return err
}
