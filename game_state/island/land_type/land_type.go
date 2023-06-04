package land_type

//go:generate go run github.com/dmarkham/enumer -type=LandType
type LandType int

const (
	// This should always be the 0th value, so we can throw an error if someone accidentally forgets to set this
	Invalid LandType = iota
	Jungle
	Wetlands
	Desert
	Mountain
	Ocean
)
