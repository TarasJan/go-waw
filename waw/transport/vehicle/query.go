package vehicle

type VehicleQuery struct {
	Type          VehicleType `json:"type"`
	Line          int64       `json:"line"`
	BrigadeNumber int64       `json:"brigade"`
}

type VehicleQueryOption func(*VehicleQuery)

func WithLine(line int64) VehicleQueryOption {
	return func(query *VehicleQuery) {
		query.Line = line
	}
}

func WithBrigade(brigade int64) VehicleQueryOption {
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
