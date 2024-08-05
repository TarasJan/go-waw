package stop

import (
	"strconv"
	"time"

	"github.com/TarasJan/go-waw/waw"
)

type Stop struct {
	BusStopId  string // Id of the group of stops with the same name
	BusStopNr  string // Number of the stop in its group
	Name       string // Name of the bus stop group
	StreetId   string
	Lat        float64
	Lon        float64
	Direction  string
	ActiveFrom time.Time
}

func NewStopFrom(valueObject waw.WawValue) *Stop {
	values := valueObject.ToMap()
	lat, _ := strconv.ParseFloat(values["szer_geo"], 64)
	lon, _ := strconv.ParseFloat(values["dlug_geo"], 64)
	time, _ := time.Parse("2006-01-02 15:04:05", values["obowiazuje_od"])
	return &Stop{
		BusStopId:  values["zespol"],
		BusStopNr:  values["slupek"],
		Name:       values["nazwa_zespolu"],
		StreetId:   values["id_ulicy"],
		Lat:        lat,
		Lon:        lon,
		Direction:  values["kierunek"],
		ActiveFrom: time,
	}
}
