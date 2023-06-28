package effects

import (
	"fmt"
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game/game_state"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/filter"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/land_state"
	"github.com/mieubrisse/open-spirit-island/game/input"
	"github.com/mieubrisse/open-spirit-island/game/static_assets"
	"sort"
)

type gatherableObjectCoords struct {
	// Index of the land where the object currently resides
	currentLandIdx int

	// Position within the array of same-type objects
	currentObjectPositionIdx int
}

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
	var objectHpGetter func(land land_state.LandState) []int
	var objectHpSetter func(land *land_state.LandState, newList []int)
	// TODO Account for base health increasing!!
	var objectBaseHealth int
	switch objectType {
	case Dahan:
		targetFilter.DahanMin = 1
		objectEmoji = static_assets.DahanSymbol
		objectHpGetter = func(land land_state.LandState) []int {
			return land.DahanHealth
		}
		objectHpSetter = func(land *land_state.LandState, newList []int) {
			land.DahanHealth = newList
		}
		objectBaseHealth = island.DahanBaseHealth
	case Explorer:
		targetFilter.ExplorersMin = 1
		objectEmoji = static_assets.ExplorerSymbol
		objectHpGetter = func(land land_state.LandState) []int {
			return land.ExplorerHealth
		}
		objectHpSetter = func(land *land_state.LandState, newList []int) {
			land.ExplorerHealth = newList
		}
		objectBaseHealth = island.ExplorerBaseHealth
	case Town:
		targetFilter.TownsMin = 1
		objectEmoji = static_assets.TownSymbol
		objectHpGetter = func(land land_state.LandState) []int {
			return land.TownHealth
		}
		objectHpSetter = func(land *land_state.LandState, newList []int) {
			land.TownHealth = newList
		}
		objectBaseHealth = island.TownBaseHealth
	case City:
		targetFilter.CitiesMin = 1
		objectEmoji = static_assets.CitySymbol
		objectHpGetter = func(land land_state.LandState) []int {
			return land.CityHealth
		}
		objectHpSetter = func(land *land_state.LandState, newList []int) {
			land.CityHealth = newList
		}
		objectBaseHealth = island.CityBaseHealth
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

	applicator := func(gameState game_state.GameState, targetLandIdx int) game_state.GameState {
		// Get adjacent lands with the object type
		eligibleGatherSourcesSet := gameState.BoardState.FilterLands(filter.IslandFilter{
			SourceNumbers: set.New(targetLandIdx),
			MinRange:      1,
			MaxRange:      1,
			TargetFilter:  targetFilter,
		})

		if len(eligibleGatherSourcesSet) == 0 {
			return gameState
		}

		eligibleGatherSourcesList := eligibleGatherSourcesSet.Slice()
		sort.Ints(eligibleGatherSourcesList)

		// Find all the matching object types in the lands, and build:
		// - prompts
		// - references to where they are
		prompts := make([]string, 0, len(eligibleGatherSourcesList))
		objectCoords := make([]gatherableObjectCoords, 0, len(eligibleGatherSourcesList))
		for i, landIdx := range eligibleGatherSourcesList {
			land := gameState.BoardState.Lands[i]

			objectHps := objectHpGetter(land)

			prompts = append(
				prompts,
				fmt.Sprintf("%s %d (%d/%d)"),
			)
		}

		input.GetMultipleSelections(effectStr, min, max)

		// TODO log a "no X to gather" message
	}

	return Effect{
		ReadableStr: effectStr,
		Applicator:  applicator,
	}
}
