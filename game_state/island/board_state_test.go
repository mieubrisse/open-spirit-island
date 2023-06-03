package island

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIslandBoardState_GetAdjacentLands(t *testing.T) {
	board := NewBoardA()

	require.Equal(t, []int{1, 2, 3}, board.GetAdjacentLands(0))
	require.Equal(t, []int{0, 2, 4, 5, 6}, board.GetAdjacentLands(1))
	require.Equal(t, []int{0, 1, 3, 4}, board.GetAdjacentLands(2))
	require.Equal(t, []int{0, 2, 4}, board.GetAdjacentLands(3))
	require.Equal(t, []int{1, 2, 3, 5}, board.GetAdjacentLands(4))
	require.Equal(t, []int{1, 4, 6, 7, 8}, board.GetAdjacentLands(5))
	require.Equal(t, []int{1, 5, 8}, board.GetAdjacentLands(6))
	require.Equal(t, []int{5, 8}, board.GetAdjacentLands(7))
	require.Equal(t, []int{5, 6, 7}, board.GetAdjacentLands(8))
}
