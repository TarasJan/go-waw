package main

import (
	"fmt"
	"gowaw/gowaw"
	"os"
)

func main() {
	apikey := os.Getenv("WAW_KEY")
	client := gowaw.NewClient(apikey)
	vehicleLocations := client.Fetch()
	for _, location := range vehicleLocations {
		fmt.Printf("%+v\n", location)
	}
}
