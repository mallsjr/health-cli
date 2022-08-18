/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
)

var serviceEndpoint = map[string]string{
	"door":      "https://opsdoor-development.app.wtcdev1.paas.fedex.com/actuator/health",
	"equipment": "http://localhost:8080/equipment",
}

type Health struct {
	Status string `json:"status"`
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
		getServiceStatus(name)
	},
}

func getServiceStatus(name string) {
	fmt.Println("getServiceStatus called with name:", name)

	url := serviceEndpoint[name]

	responseBytes := getActuatorHealth(url)

	health := Health{}

	if err := json.Unmarshal(responseBytes, &health); err != nil {
		log.Panicf("Error unmarshalling response: %v", err)
	}

	fmt.Printf("%s service is reporting status of %s\n", name, health.Status)
}

func getActuatorHealth(url string) []byte {
	fmt.Println("getActuatorHealth called with url:", url)

	request, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		log.Printf("Error creating request: %v", err)
	}

	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Printf("Error getting response: %v", err)
	}

	responseBytes, err := io.ReadAll(response.Body)

	if err != nil {
		log.Printf("Error reading response: %v", err)
	}

	return responseBytes
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
