package effects

import (
	"fmt"
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game/game_state"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/filter"
	"github.com/mieubrisse/open-spirit-island/game/static_assets"
	"sort"
)

// Useful for DRY gather/push mechanics
type ObjectType int

const (
	Dahan ObjectType = iota
	Explorer
	Town
	City
)

func NewGatherDahanEffect(limit int) Effect {
	applicator := func(state game_state.GameState, targetLandIdx int) game_state.GameState {
		// Get adjacent lands with Dahan
		state.BoardState.FilterLands(filter.IslandFilter{
			SourceNumbers: set.New(targetLandIdx),
			MinRange:      1,
			MaxRange:      1,
			TargetFilter: filter.LandFilter{
				DahanMin: 1,
			},
		})

	}

	return Effect{
		ReadableStr: fmt.Sprintf("Gather up to %d 🛖", limit),
		Applicator:  applicator,
	}
}

func NewGatherObjectEffect(min int, max int, objectType ObjectType) func(state game_state.GameState) game_state.GameState {
	if max < min {
		panic(fmt.Errorf("Cannot have a gather with max < min (got min %d, max %d)", min, max))
	}
	if max < 1 {
		panic(fmt.Errorf("Cannot have a gather with max < 1 (got %d)", max))
	}

	targetFilter := filter.LandFilter{
		DahanMin: 1,
	}

	var objectEmoji string
	switch objectType {
	case Dahan:
		targetFilter.DahanMin = 1
		objectEmoji = static_assets.DahanSymbol
	case Explorer:
		targetFilter.ExplorersMin = 1
		objectEmoji = static_assets.ExplorerSymbol
	case Town:
		targetFilter.TownsMin = 1
		objectEmoji = static_assets.TownSymbol
	case City:
		targetFilter.CitiesMin = 1
		objectEmoji = static_assets.CitySymbol
	default:
		panic(fmt.Errorf("Unrecognized object type: %d", objectType))
	}

	var effectStr string
	if min == 0 {
		effectStr = fmt.Sprintf("Gather up to %d %s", min, objectEmoji)
	} else if min == max {
		effectStr = fmt.Sprintf("Gather %d %s", min, objectEmoji)
	} else {
		effectStr = fmt.Sprintf("Gather %d to %d %s", min, max, objectEmoji)
	}

	applicator := func(state game_state.GameState, targetLandIdx int) game_state.GameState {
		// Get adjacent lands with the object type
		eligibleGatherSourcesSet := state.BoardState.FilterLands(filter.IslandFilter{
			SourceNumbers: set.New(targetLandIdx),
			MinRange:      1,
			MaxRange:      1,
			TargetFilter:  targetFilter,
		})

		if len(eligibleGatherSourcesSet) == 0 {
			return state
		}

		eligibleGatherSourcesList := eligibleGatherSourcesSet.Slice()
		sort.Ints(eligibleGatherSourcesList)

		objectIdx := make([]int, 0, len(eligibleGatherSourcesList))
		for i, landIdx := range eligibleGatherSourcesList {
			switch objectType {
			case Dahan:
				for
			case Explorer:
			case Town:
			case City:
			}
		}

		// TODO log a "no X to gather" message
	}
}
