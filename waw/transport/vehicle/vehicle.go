package vehicle

import "time"

const ResourceId = "f2e5503e-927d-4ad3-9500-4ab9e55deb59"

type VehicleType int64

const (
	Bus  VehicleType = 1 // indexing starting from 1 as in the API
	Tram VehicleType = 2
)

type Vehicle struct {
	Line          string `json:"Lines"`
	Number        string `json:"VehicleNumber"`
	BrigadeNumber string `json:"Brigade"`
}

type VehicleLocation struct {
	Vehicle
	Time LocationTime `json:"Time"`
	Lat  float64      `json:"Lat"`
	Lon  float64      `json:"Lon"`
}

type LocationTime time.Time

func (lt LocationTime) Format(layout string) string {
	return time.Time(lt).Format(layout)
}

func (lt LocationTime) String() string {
	return lt.Format(time.DateTime)
}

func (lt *LocationTime) UnmarshalJSON(data []byte) error {
	data = data[len(`"`) : len(data)-len(`"`)]
	t, err := time.Parse(time.DateTime, string(data))
	if err != nil {
		return err
	}
	*lt = LocationTime(t)
	return nil
}
