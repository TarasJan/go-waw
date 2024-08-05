package route

type RoutesResponse struct {
	Result LineRoutes `json:"result"`
}

type LineRoutes map[string]Routes // key being line number

type Routes map[string]Route // Key being the relation

type Route map[string]RouteStop // Key being the stop number on the route

type RouteStop struct {
	StreetId  string `json:"ulica_id"`
	BusStopId string `json:"nr_zespolu"`
	BusStopNr string `json:"nr_przystanku"`
	Type      string `json:"typ"`       // Bus stop type
	Distance  int64  `json:"odleglosc"` // Distance from route start point in meters
}
