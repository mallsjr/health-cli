/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

var serviceEndpoint = map[string]string{
	"door":      "http://localhost:8080/door",
	"equipment": "http://localhost:8080/equipment",
}

//type Joke struct {
//ID     string `json:"id"`
//Joke   string `json:"joke"`
//Status int    `json:"status"`
//}

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

	//joke := Joke{}

	//if err := json.Unmarshal(responseBytes, &joke); err != nil {
	//log.Panicf("Error unmarshalling response: %v", err)
	//}

	//fmt.Println(joke.Joke)
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

	responseBytes, err := ioutil.ReadAll(response.Body)

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
