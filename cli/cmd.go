package main

import "github.com/spf13/cobra"

// The root command serves as a base for all subcommands.
var root = &cobra.Command{
	Use:   `foodunit`,
	Short: `FoodUnit CLI`,
	Run:   rootHandler,
}

// `offers` runs an API request for the active offers, fetches the data
// and prints the result.
var offers = &cobra.Command{
	Use:   `offers`,
	Short: `Lists all active offers`,
	Run:   offerHandler,
}

// `menu` displays the menu for a given supplier. The supplier ID may
// be obtained by running `offers` before for example.
var menu = &cobra.Command{
	Use:   `menu`,
	Short: `Lists all menu for a given supplier`,
	Run:   menuHandler,
}

// `supplier` fetches all important data for a given supplier. Its ID
// can be retrieved by running `offers` for example.
var supplier = &cobra.Command{
	Use:   `supplier`,
	Short: `Displays basic data for a given supplier`,
	Run:   supplierHandler,
}

var orders = &cobra.Command{
	Use:   `orders`,
	Short: `Lists all orders associated with a given offer`,
	Run:   ordersHandler,
}

// Execute creates all necessary flags for each command and then attaches
// the commands to the root command. The root command is then executed.
func Execute() error {
	menu.Flags().String("supplier", "", `The ID of the supplier`)
	supplier.Flags().String("supplier", "", `The ID of the supplier`)
	orders.Flags().String("offer", "", `The ID of the offer`)

	root.AddCommand(offers)
	root.AddCommand(menu)
	root.AddCommand(supplier)
	root.AddCommand(orders)

	err := root.Execute()
	return err
}
