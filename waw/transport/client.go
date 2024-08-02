package transport

import (
	"github.com/TarasJan/go-waw/waw/transport/dictionary"
	"github.com/TarasJan/go-waw/waw/transport/vehicle"
)

type Client struct {
	APIKey     string
	Vehicles   *vehicle.VehicleClient
	Dictionary *dictionary.DictionaryClient
}

func NewClient(key string) *Client {
	return &Client{
		APIKey:     key,
		Vehicles:   vehicle.NewVehicleClient(key),
		Dictionary: dictionary.NewDictionaryClient(key),
	}
}
