package types

// map.go provides types for describing maps.

type MapType int

var StrToMapType = map[string]MapType{
	"plan":       PLAN,
	"satellite":  SATELLITE,
	"hybrid":     HYBRID,
	"descriptor": DESCRIPTOR,
}

const (
	PLAN MapType = iota
	SATELLITE
	HYBRID
	DESCRIPTOR
)
