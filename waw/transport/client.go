package transport

import (
	"github.com/TarasJan/go-waw/waw"
	"github.com/TarasJan/go-waw/waw/transport/dictionary"
	"github.com/TarasJan/go-waw/waw/transport/route"
	"github.com/TarasJan/go-waw/waw/transport/stop"
	"github.com/TarasJan/go-waw/waw/transport/timetable"
	"github.com/TarasJan/go-waw/waw/transport/vehicle"
)

type Client struct {
	APIKey     string
	Vehicles   *vehicle.VehicleClient
	Dictionary *dictionary.DictionaryClient
	Timetables *timetable.TimetableClient
	Stops      *stop.StopClient
	Routes     *route.RouteClient
}

func NewClient(key ...string) (*Client, error) {
	APIKey, err := waw.GetAPIKey(key...)
	if err != nil {
		return nil, err
	}
	return &Client{
		APIKey:     APIKey,
		Vehicles:   vehicle.NewVehicleClient(APIKey),
		Dictionary: dictionary.NewDictionaryClient(APIKey),
		Timetables: timetable.NewTimetableClient(APIKey),
		Stops:      stop.NewStopClient(APIKey),
		Routes:     route.NewRouteClient(APIKey),
	}, nil
}
