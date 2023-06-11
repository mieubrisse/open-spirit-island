package island

import (
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/filter"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewCoastalLandsFilter(t *testing.T) {
	board := NewBoardA()

	actualCoastalLands := board.FilterLands(filter.NewCoastalLandsFilter())
	require.Equal(t, set.New(1, 2, 3), actualCoastalLands)
}
