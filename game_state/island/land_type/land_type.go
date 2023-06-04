package land_type

import "github.com/bobg/go-generics/v2/set"

//go:generate go run github.com/dmarkham/enumer -type=LandType
type LandType int

const (
	Jungle LandType = iota
	Wetlands
	Desert
	Mountain
	Ocean
)

var NonOceanLandTypes = set.New(Jungle, Wetlands, Desert, Mountain)
