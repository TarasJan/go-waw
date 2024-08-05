package timetable

import (
	"fmt"
	"time"

	"github.com/TarasJan/go-waw/waw"
)

type Timetable struct {
	Line    string            `json:"line"`
	Records []TimetableRecord `json:"records"`
}

type TimetableRecord struct {
	Route       string    `json:"trasa"`
	Destination string    `json:"kierunek"` // usually the terminus of the vehicle
	Brigade     string    `json:"brygada"`
	ArrivalTime time.Time `json:"czas"` // arrival time in local polish time (24h format)
}

func NewTimetableFrom(response waw.WawResponse) (*Timetable, error) {
	records := make([]TimetableRecord, 0)

	for _, values := range response.Result {
		timeRecordParams := values.ToMap()

		arrival, err := time.Parse("15:04:05", timeRecordParams["czas"])
		if err != nil {
			return nil, fmt.Errorf("unparseable time expression found in the response: %s", timeRecordParams["czas"])
		}

		records = append(records, TimetableRecord{
			Route:       timeRecordParams["trasa"],
			Destination: timeRecordParams["kierunek"],
			Brigade:     timeRecordParams["brygada"],
			ArrivalTime: arrival,
		})
	}

	return &Timetable{
		Records: records,
	}, nil

}

func (t *Timetable) WithLine(line string) *Timetable {
	t.Line = line
	return t
}
