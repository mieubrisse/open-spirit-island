package filter

import (
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/land_type"
)

func NewCoastalLandsFilter() IslandFilter {
	return IslandFilter{
		SourceFilter: LandFilter{
			LandTypes: set.New(land_type.Ocean),
		},
		MinRange: 1,
		MaxRange: 1,
		TargetFilter: LandFilter{
			LandTypes: land_type.NonOceanLandTypes,
		},
	}
}

func NewAdjacentNonOceanLandsFilter(sourceIdx int) IslandFilter {
	return IslandFilter{
		SourceNumbers: set.New(sourceIdx),
		MinRange:      1,
		MaxRange:      1,
		TargetFilter: LandFilter{
			LandTypes: land_type.NonOceanLandTypes,
		},
	}
}
