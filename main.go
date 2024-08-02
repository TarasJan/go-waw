package main

import (
	"fmt"
	"os"

	"github.com/TarasJan/go-waw/waw/transport"
	"github.com/TarasJan/go-waw/waw/transport/vehicle"
)

func main() {
	apikey := os.Getenv("GOWAW_KEY")
	client := transport.NewClient(apikey)
	// vehicleLocations, err := client.FetchTrams()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, location := range vehicleLocations {
	// 	fmt.Printf("%+v\n", location)
	// }

	// Specific query - bus 411
	// specificLineLocations, err := client.Vehicles.FetchBuses(vehicle.WithLine("411"))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, location := range specificLineLocations {
	// 	fmt.Printf("%+v\n", location)
	// }

	// Specific query - tram 4
	tramLocations, err := client.Vehicles.FetchTrams(vehicle.WithLine("4"))
	if err != nil {
		fmt.Println(err)
	}
	for _, location := range tramLocations {
		fmt.Printf("%+v\n", location)
	}

	// Dictionary query
	dictionary, err := client.Dictionary.Get()
	if err != nil {
		fmt.Println(err)
	}

	for _, street := range dictionary.Streets {
		fmt.Println(street)
	}

	// for _, location := range dictionary.Locations {
	// 	fmt.Println(location)
	// }
}
