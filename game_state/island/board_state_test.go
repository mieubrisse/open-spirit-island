package island

import (
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game_state/island/filter"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIslandBoardState_GetMatchingLands_AdjacentLands(t *testing.T) {
	board := NewBoardA()

	expectedAdjacenciesOfLands := []set.Of[int]{
		set.New(1, 2, 3),
		set.New(0, 2, 4, 5, 6),
		set.New(0, 1, 3, 4),
		set.New(0, 2, 4),
		set.New(1, 2, 3, 5),
		set.New(1, 4, 6, 7, 8),
		set.New(1, 5, 8),
		set.New(5, 8),
		set.New(5, 6, 7),
	}

	for idx, expectedAdjacenciesForLand := range expectedAdjacenciesOfLands {
		actualAdjacencies := board.FilterLands(filter.IslandFilter{
			SourceNumbers: set.New(idx),
			MinRange:      1,
			MaxRange:      1,
		})
		require.Equal(t, expectedAdjacenciesForLand, actualAdjacencies, "Incorrect adjacencies for land %d", idx)
	}
}

func TestIslandBoardState_GetMatchingLands_CoastalLands(t *testing.T) {
	board := NewBoardA()

	actualCoastalLands := board.FilterLands(filter.NewCoastalLandsFilter())
	require.Equal(t, set.New(1, 2, 3), actualCoastalLands)
}
