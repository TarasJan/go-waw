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
		sanitized_time := normalizeZTMTime(timeRecordParams["czas"])

		arrival, err := time.Parse("15:04:05", sanitized_time)
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

// The public transport sometimes expresses midnight as 24 which breaks the time parsing
func normalizeZTMTime(time string) string {
	if time[0:2] == "24" {
		return "00" + time[2:]
	}

	return time
}
