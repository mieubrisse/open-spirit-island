package island

// Corresponds to the default board A from the vanilla game
func NewBoardA() IslandBoardState {
	return IslandBoardState{
		Lands: []LandState{
			// 0
			{
				LandType: Ocean,
			},
			// 1
			{
				LandType: Mountain,
			},
			// 2
			{
				LandType:  Wetlands,
				NumCities: 1,
				NumDahan:  1,
			},
			// 3
			{
				LandType: Jungle,
				NumDahan: 2,
			},
			// 4
			{
				LandType:  Desert,
				NumBlight: 1,
			},
			// 5
			{
				LandType: Wetlands,
			},
			// 6
			{
				LandType: Mountain,
				NumDahan: 1,
			},
			// 7
			{
				LandType: Desert,
				NumDahan: 2,
			},
			// 8
			{
				LandType: Jungle,
				NumTowns: 1,
			},
		},
		Adjacencies: [][]int{
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
		},
	}
}
