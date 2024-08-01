package vehicle

const VehicleRequestResourcId = "f2e5503e-927d-4ad3-9500-4ab9e55deb59"

type VehicleType int64

const (
	Bus  VehicleType = 1
	Tram VehicleType = 2
)

type Vehicle struct {
	Line          string `json:"Lines"`
	Number        string `json:"VehicleNumber"`
	BrigadeNumber string `json:"Brigade"`
}

type VehicleLocation struct {
	Vehicle
	Time string  `json:"Time"`
	Lat  float64 `json:"Lat"`
	Lon  float64 `json:"Lon"`
}
