package island

import (
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/filter"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/land_state"
	"github.com/yourbasic/graph"
)

// TODO one day make these customizable?
const (
	DahanBaseHealth = 2
	DahanBaseDamage = 2

	CityBaseHealth = 3
	CityBaseDamage = 3

	TownBaseHealth = 2
	TownBaseDamage = 2

	ExplorerBaseHealth = 1
	ExplorerBaseDamage = 1
)

// TODO multiple boards
type IslandBoardState struct {
	Graph *graph.Immutable

	Lands []land_state.LandState

	/*
		// Mapping of land_index -> bordering_land_indexes
		Adjacencies [][]int

	*/
}

// TODO nice error-handling
func (state IslandBoardState) AddPresence(landIdx int) IslandBoardState {
	state.Lands[landIdx].NumPresence++
	return state
}

// Swiss army knife for selecting lands on the island
func (state IslandBoardState) FilterLands(filter filter.IslandFilter) set.Of[int] {
	sourcesIdx := set.New[int]()
	for idx, land := range state.Lands {
		if filter.SourceNumbers != nil {
			if !filter.SourceNumbers.Has(idx) {
				continue
			}
		}

		if filter.SourceFilter.Match(land) {
			sourcesIdx.Add(idx)
		}
	}

	result := set.New[int]()
	for sourceIdx := range sourcesIdx {
		_, distancesInt64 := graph.ShortestPaths(state.Graph, sourceIdx)
		distances := make([]int, len(distancesInt64))
		for i, value := range distancesInt64 {
			distances[i] = int(value)
		}

		for targetLandIdx, distance := range distances {
			targetLand := state.Lands[targetLandIdx]

			if distance < filter.MinRange {
				continue
			}

			if distance > filter.MaxRange {
				continue
			}

			if !filter.TargetFilter.Match(targetLand) {
				continue
			}

			result.Add(targetLandIdx)
		}
	}

	return result
}
