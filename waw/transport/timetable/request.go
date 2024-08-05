package timetable

import (
	"net/url"

	"github.com/TarasJan/go-waw/waw/transport/stop"
)

type TimetableRequest struct {
	ResourceId string `json:"resource_id"`
	ApiKey     string `json:"apikey"`
	Name       string `json:"name"`
	BusStopId  string `json:"busstopId"`
	BusStopNr  string `json:"busstopNr"`
	Line       string `json:"line"`
}

func (tr *TimetableRequest) WithName(name string) *TimetableRequest {
	tr.Name = name
	return tr
}

func (tr *TimetableRequest) WithLine(line string) *TimetableRequest {
	tr.Line = line
	return tr
}

func (tr *TimetableRequest) WithStop(stop *stop.Stop) *TimetableRequest {
	tr.BusStopId = stop.BusStopId
	tr.BusStopNr = stop.BusStopNr
	return tr
}

func (tr *TimetableRequest) ToValues() url.Values {
	values := url.Values{}
	values.Set("apikey", tr.ApiKey)
	values.Set("id", tr.ResourceId)
	if tr.Name != "" {
		values.Set("name", tr.Name)
	}

	if tr.Line != "" {
		values.Set("line", tr.Line)
	}

	if tr.BusStopId != "" {
		values.Set("busstopId", tr.BusStopId)
	}

	if tr.BusStopNr != "" {
		values.Set("busstopNr", tr.BusStopNr)
	}

	return values
}
