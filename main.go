package main

import (
	"errors"
	"fmt"
	"gowaw/gowaw"
	"gowaw/gowaw/transport"
	"os"
)

func main() {
	apikey := os.Getenv("GOWAW_KEY")
	client := transport.NewClient(apikey)
	vehicleLocations, err := client.FetchTrams()
	if err != nil {
		if errors.Is(err, &gowaw.UnauthorizedAccessError{}) {
			fmt.Println("The API did not recognize the provided token, make sure to set up GOWAW_KEY environment variable")
		}
		fmt.Println(err)
	}
	for _, location := range vehicleLocations {
		fmt.Printf("%+v\n", location)
	}
}
