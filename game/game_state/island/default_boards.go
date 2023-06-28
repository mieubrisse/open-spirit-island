package island

import (
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/land_state"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/land_type"
	"github.com/yourbasic/graph"
)

// Corresponds to the default board A from the vanilla game
func NewBoardA() IslandBoardState {
	landStates := []land_state.LandState{
		// 0
		{
			LandType: land_type.Ocean,
		},
		// 1
		{
			LandType: land_type.Mountain,
		},
		// 2
		{
			LandType:         land_type.Wetlands,
			CityDamageTaken:  []int{0},
			DahanDamageTaken: []int{0},
		},
		// 3
		{
			LandType:         land_type.Jungle,
			DahanDamageTaken: []int{0, 0},
		},
		// 4
		{
			LandType:  land_type.Desert,
			NumBlight: 1,
		},
		// 5
		{
			LandType: land_type.Wetlands,
		},
		// 6
		{
			LandType:         land_type.Mountain,
			DahanDamageTaken: []int{0},
		},
		// 7
		{
			LandType:         land_type.Desert,
			DahanDamageTaken: []int{0, 0},
		},
		// 8
		{
			LandType:        land_type.Jungle,
			TownDamageTaken: []int{0},
		},
	}

	adjacencies := [][]int{
		// Duplicates aren't an issue, but to keep this easy to reason about it's always (lower, higher)
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 2},
		{1, 4},
		{1, 5},
		{1, 6},
		{2, 3},
		{2, 4},
		{3, 4},
		{4, 5},
		{5, 6},
		{5, 7},
		{5, 8},
		{6, 8},
		{7, 8},
	}

	mutableGraph := graph.New(len(landStates))
	for _, pair := range adjacencies {
		// 1 signifies the range of 1 between each adjacent land
		mutableGraph.AddBothCost(pair[0], pair[1], 1)
	}
	immutableGraph := graph.Sort(mutableGraph)

	return IslandBoardState{
		Lands: landStates,
		Graph: immutableGraph,
	}
}
