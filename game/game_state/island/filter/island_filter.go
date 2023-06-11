package filter

import (
	"github.com/bobg/go-generics/v2/set"
	"math"
)

type IslandFilter struct {
	// If nil, all numbers will be allowed
	SourceNumbers set.Of[int]

	SourceFilter LandFilter

	MinRange int
	MaxRange int

	TargetFilter LandFilter
}

func New() IslandFilter {
	return IslandFilter{
		SourceFilter: LandFilter{},
		MinRange:     0,
		MaxRange:     math.MaxInt,
		TargetFilter: LandFilter{},
	}
}

/*
func New(opts ...IslandFilterOpt) IslandFilter {
	result := IslandFilter{
		// NOTE: default LandFilter is to exclude nothing
		SourceFilter: LandFilter{},
		MinRange:     0,
		MaxRange:     math.MaxInt,
		TargetFilter: LandFilter{},
	}

	for _, opt := range opts {
		opt(&result)
	}

	return result
}


*/
