package game_state

import (
	"github.com/mieubrisse/open-spirit-island/game_state/island/land_state"
	"github.com/mieubrisse/open-spirit-island/game_state/island/land_type"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGameState_RunInvaderPhase(t *testing.T) {
	game := NewTestGame()

	game = game.RunInvaderPhase()

	// At this point -- Ravage: none, Build: jungle
	require.False(t, game.InvaderState.RavageSlot.IsCardPresent)
	require.True(t, game.InvaderState.BuildSlot.IsCardPresent)
	require.Equal(
		t,
		[]land_state.LandState{
			// 0
			{
				LandType: land_type.Ocean,
			},
			// 1
			{
				LandType:    land_type.Mountain,
				NumPresence: 1,
			},
			// 2
			{
				LandType:  land_type.Wetlands,
				NumCities: 1,
				NumDahan:  1,
			},
			// 3
			{
				LandType:     land_type.Jungle,
				NumDahan:     2,
				NumExplorers: 1,
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
				LandType: land_type.Mountain,
				NumDahan: 1,
			},
			// 7
			{
				LandType: land_type.Desert,
				NumDahan: 2,
			},
			// 8
			{
				LandType:     land_type.Jungle,
				NumTowns:     1,
				NumExplorers: 1,
			},
		},
		game.BoardState.Lands,
	)

	game = game.RunInvaderPhase()

	// At this point -- Ravage: jungle, build: desert
	require.True(t, game.InvaderState.RavageSlot.IsCardPresent)
	require.True(t, game.InvaderState.BuildSlot.IsCardPresent)
	require.Equal(
		t,
		[]land_state.LandState{
			// 0
			{
				LandType: land_type.Ocean,
			},
			// 1
			{
				LandType:    land_type.Mountain,
				NumPresence: 1,
			},
			// 2
			{
				LandType:  land_type.Wetlands,
				NumCities: 1,
				NumDahan:  1,
			},
			// 3
			{
				LandType:     land_type.Jungle,
				NumDahan:     2,
				NumExplorers: 1,
				NumTowns:     1,
			},
			// 4
			{
				LandType:     land_type.Desert,
				NumBlight:    1,
				NumExplorers: 1,
			},
			// 5
			{
				LandType: land_type.Wetlands,
			},
			// 6
			{
				LandType: land_type.Mountain,
				NumDahan: 1,
			},
			// 7
			{
				LandType:     land_type.Desert,
				NumDahan:     2,
				NumExplorers: 1,
			},
			// 8
			{
				LandType:     land_type.Jungle,
				NumTowns:     1,
				NumExplorers: 1,
				NumCities:    1,
			},
		},
		game.BoardState.Lands,
	)

	game = game.RunInvaderPhase()

	// At this point -- ravage: desert, build: wetlands
	require.True(t, game.InvaderState.RavageSlot.IsCardPresent)
	require.True(t, game.InvaderState.BuildSlot.IsCardPresent)
	require.Equal(
		t,
		[]land_state.LandState{
			// 0
			{
				LandType: land_type.Ocean,
			},
			// 1
			{
				LandType:    land_type.Mountain,
				NumPresence: 1,
			},
			// 2
			{
				LandType:     land_type.Wetlands,
				NumCities:    1,
				NumDahan:     1,
				NumExplorers: 1,
			},
			// 3
			{
				LandType:     land_type.Jungle,
				NumDahan:     1,
				NumExplorers: 1,
				NumTowns:     1,
				NumBlight:    1,
			},
			// 4
			{
				LandType:     land_type.Desert,
				NumBlight:    1,
				NumExplorers: 1,
				NumTowns:     1,
			},
			// 5
			{
				LandType:     land_type.Wetlands,
				NumExplorers: 1,
			},
			// 6
			{
				LandType: land_type.Mountain,
				NumDahan: 1,
			},
			// 7
			{
				LandType:     land_type.Desert,
				NumDahan:     2,
				NumExplorers: 1,
				NumTowns:     1,
			},
			// 8
			{
				LandType:     land_type.Jungle,
				NumTowns:     1,
				NumExplorers: 1,
				NumCities:    1,
				NumBlight:    1,
			},
		},
		game.BoardState.Lands,
	)

	// Can't go further because we need user input for the Blight cascade that occurs on Desert 4
}
