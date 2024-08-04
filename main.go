package main

import (
	"fmt"

	"github.com/TarasJan/go-waw/waw/transport"
)

func main() {
	client, err := transport.NewClient()
	if err != nil {
		panic(err)
	}
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
	// tramLocations, err := client.Vehicles.FetchTrams(vehicle.WithLine("4"))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, location := range tramLocations {
	// 	fmt.Printf("%+v\n", location)
	// }

	// // Dictionary query
	// dictionary, err := client.Dictionary.Get()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// for _, street := range dictionary.Streets {
	// 	fmt.Println(street)
	// }

	// for _, location := range dictionary.Locations {
	// 	fmt.Println(location)
	// }

	// timetable, err := client.Timetables.Get()
	// if err != nil {
	// 	panic(err)
	// }
	// for _, tableObject := range timetable {
	// 	fmt.Println(tableObject.Values)
	// }

	stops, err := client.Stops.Get()
	if err != nil {
		panic(err)
	}
	for _, stop := range stops {
		fmt.Println(stop.GroupName)
	}
}
