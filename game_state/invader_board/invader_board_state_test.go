package invader_board

import (
	"github.com/mieubrisse/open-spirit-island/game_state/decks/fear"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInvaderBoardState_AddFear(t *testing.T) {
	fearDeck := make([]fear.FearCard, 9)
	for i := 0; i < 9; i++ {
		fearDeck[i] = fear.NewDummyFearCard()
	}
	state := InvaderBoardState{
		UnearnedFear:          4,
		EarnedFear:            0,
		TerrorLevelThresholds: []int{0, 3, 6, 9},
		UnearnedFearCards:     fearDeck,
		EarnedFearCards:       make([]fear.FearCard, 0),
		RemainingInvaderDeck:  nil,
		BuildSlot:             MaybeInvaderCard{},
		RavageSlot:            MaybeInvaderCard{},
		InvaderDeckDiscard:    nil,
	}

	require.Equal(t, 1, state.GetTerrorLevel())
	require.Equal(t, 3, state.countFearCardsTillNextTerrorLevel())

	state = state.AddFear(3)
	require.Equal(t, 3, state.EarnedFear)
	require.Equal(t, 1, state.UnearnedFear)
	require.Equal(t, 0, len(state.EarnedFearCards))
	require.Equal(t, 9, len(state.UnearnedFearCards))
	require.Equal(t, 3, state.countFearCardsTillNextTerrorLevel())
	require.Equal(t, 1, state.GetTerrorLevel())

	state = state.AddFear(2)
	require.Equal(t, 1, state.EarnedFear)
	require.Equal(t, 3, state.UnearnedFear)
	require.Equal(t, 1, len(state.EarnedFearCards))
	require.Equal(t, 8, len(state.UnearnedFearCards))
	require.Equal(t, 2, state.countFearCardsTillNextTerrorLevel())
	require.Equal(t, 1, state.GetTerrorLevel())

	state = state.AddFear(8)
	require.Equal(t, 1, state.EarnedFear)
	require.Equal(t, 3, state.UnearnedFear)
	require.Equal(t, 3, len(state.EarnedFearCards))
	require.Equal(t, 6, len(state.UnearnedFearCards))
	require.Equal(t, 3, state.countFearCardsTillNextTerrorLevel())
	require.Equal(t, 2, state.GetTerrorLevel())
}
