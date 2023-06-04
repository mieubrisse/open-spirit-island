package game_state

import (
	"github.com/mieubrisse/open-spirit-island/game_state/decks/fear"
	invader_deck2 "github.com/mieubrisse/open-spirit-island/game_state/decks/invader_deck"
	"github.com/mieubrisse/open-spirit-island/game_state/invader_board"
	"github.com/mieubrisse/open-spirit-island/game_state/island"
	"github.com/mieubrisse/open-spirit-island/game_state/island/land_state"
	"github.com/mieubrisse/open-spirit-island/game_state/island/land_type"
	"github.com/mieubrisse/open-spirit-island/game_state/player"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGameState_RunInvaderPhase(t *testing.T) {
	game := getTestGame()

	game = game.RunInvaderPhase()
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
				LandType:     land_type.Wetlands,
				NumCities:    1,
				NumDahan:     1,
				NumExplorers: 1,
			},
			// 3
			{
				LandType:     land_type.Jungle,
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

	game = game.RunInvaderPhase()
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
				LandType:     land_type.Wetlands,
				NumCities:    1,
				NumDahan:     1,
				NumExplorers: 1,
			},
			// 3
			{
				LandType:     land_type.Jungle,
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
}

func getTestGame() GameState {
	fearDeck := make([]fear.FearCard, 9)
	for i := 0; i < 9; i++ {
		fearDeck[i] = fear.NewDummyFearCard()
	}

	// TODO parameterize this based on adversary
	invaderDeck := []invader_deck2.InvaderCard{
		// stage 1
		invader_deck2.NewSingleTypeInvaderCard(land_type.Jungle),
		invader_deck2.NewSingleTypeInvaderCard(land_type.Desert),
		invader_deck2.NewSingleTypeInvaderCard(land_type.Wetlands),
		// stage 2
		invader_deck2.NewSingleTypeAndAdversaryInvaderCard(land_type.Mountain),
		invader_deck2.NewSingleTypeAndAdversaryInvaderCard(land_type.Wetlands),
		invader_deck2.NewSingleTypeAndAdversaryInvaderCard(land_type.Desert),
		invader_deck2.NewCoastalLandsInvaderCard(),
		// stage 3
		invader_deck2.NewDoubleTypeInvaderCard(land_type.Mountain, land_type.Jungle),
		invader_deck2.NewDoubleTypeInvaderCard(land_type.Jungle, land_type.Wetlands),
		invader_deck2.NewDoubleTypeInvaderCard(land_type.Wetlands, land_type.Desert),
		invader_deck2.NewDoubleTypeInvaderCard(land_type.Mountain, land_type.Desert),
		invader_deck2.NewDoubleTypeInvaderCard(land_type.Jungle, land_type.Desert),
		// invader_deck.NewDoubleTypeInvaderCard(island.Mountain, island.Wetlands),
	}

	invaderBoardState := invader_board.InvaderBoardState{
		UnearnedFear:          4,
		EarnedFear:            0,
		TerrorLevelThresholds: []int{3, 3, 3},
		UnearnedFearCards:     fearDeck,
		EarnedFearCards:       make([]fear.FearCard, 0),
		RemainingInvaderDeck:  invaderDeck,
		BuildSlot: invader_board.MaybeInvaderCard{
			IsCardPresent: false,
			MaybeCard:     invader_deck2.InvaderCard{},
		},
		RavageSlot: invader_board.MaybeInvaderCard{
			IsCardPresent: false,
			MaybeCard:     invader_deck2.InvaderCard{},
		},
		InvaderDeckDiscard: []invader_deck2.InvaderCard{},
	}

	playerState := player.PlayerState{
		Energy: 0,
	}

	boardState := island.NewBoardA()
	boardState = boardState.AddPresence(1) // Necessary so the game isn't immediately over

	return GameState{
		InvaderState: invaderBoardState,
		PlayerState:  playerState,
		BoardState:   boardState,
	}
}
