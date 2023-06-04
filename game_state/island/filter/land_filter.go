package filter

import (
	"github.com/mieubrisse/open-spirit-island/game_state/island/land_state"
	"github.com/mieubrisse/open-spirit-island/game_state/island/land_type"
)

type LandFilter struct {
	// Only lands matching these types will be considered as sources
	// If this value is nil, then ALL land types will be allowed!
	LandTypes map[land_type.LandType]bool

	InvadersMin int

	// Only lands that match all these minimums will be considered as sources
	ExplorersMin int
	TownsMin     int
	CitiesMin    int

	BlightMin int

	DahanMin int

	PresenceMin int
}

func (filter LandFilter) Match(land land_state.LandState) bool {
	if filter.LandTypes != nil {
		if _, found := filter.LandTypes[land.LandType]; !found {
			return false
		}
	}

	if land.NumExplorers+land.NumTowns+land.NumCities < filter.InvadersMin {
		return false
	}

	if land.NumExplorers < filter.ExplorersMin {
		return false
	}
	if land.NumTowns < filter.TownsMin {
		return false
	}
	if land.NumCities < filter.CitiesMin {
		return false
	}

	if land.NumBlight < filter.BlightMin {
		return false
	}

	if land.NumDahan < filter.DahanMin {
		return false
	}
	if land.NumPresence < filter.PresenceMin {
		return false
	}

	return true
}

/*
// This is a selector, for use with the board to grab lands based on certain conditions
// It works PURELY by filtering lands - it will not add lands back after excluding them!
// To add lands back, union the result of this filter with another
type LandSelector struct {
	// Only lands matching these types will be considered as sources
	SourceLandTypes map[island.LandType]bool

	// Only cities that match all these minimums will be considered as sources
	SourceExplorersMin int
	SourceTownsMin     int
	SourceCitiesMin    int
	// TODO Dahan as source??

	SourcePresenceMin int

	MinRange int
	MaxRange int

	TargetLandTypes map[island.LandType]bool

	TargetExplorersMin int
	TargetTownsMin     int
	TargetCitiesMin    int

	TargetDahanMin int
}
*/
