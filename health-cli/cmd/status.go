/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serviceEndpoint = map[string]string{
	"door":      "http://localhost:8080/door",
	"equipment": "http://localhost:8080/equipment",
}

var name string

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status called with name:", name)
		fmt.Printf("%s :service, %s :endpoint\n", name, serviceEndpoint[name])
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	statusCmd.Flags().StringVarP(&name, "name", "n", "", "name of the service")
	statusCmd.MarkFlagRequired("name")
}
