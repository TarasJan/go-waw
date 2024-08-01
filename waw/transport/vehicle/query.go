package vehicle

type VehicleQuery struct {
	Type          VehicleType `json:"type"`
	Line          string      `json:"line"`    // string due to some lines starting with letters N or L
	BrigadeNumber string      `json:"brigade"` // string due to some brigades starting with 0
}

type VehicleQueryOption func(*VehicleQuery)

func WithLine(line string) VehicleQueryOption {
	return func(query *VehicleQuery) {
		query.Line = line
	}
}

func WithBrigade(brigade string) VehicleQueryOption {
	return func(query *VehicleQuery) {
		query.BrigadeNumber = brigade
	}
}

func WithType(vt VehicleType) VehicleQueryOption {
	return func(query *VehicleQuery) {
		query.Type = vt
	}
}

func NewBusQuery(options ...VehicleQueryOption) *VehicleQuery {
	busOptions := append(options, WithType(Bus))

	return newVehicleQuery(
		busOptions...,
	)
}

func NewTramQuery(options ...VehicleQueryOption) *VehicleQuery {
	tramOptions := append(options, WithType(Tram))

	return newVehicleQuery(
		tramOptions...,
	)
}

func newVehicleQuery(options ...VehicleQueryOption) *VehicleQuery {
	query := &VehicleQuery{}
	for _, option := range options {
		option(query)
	}

	return query
}
