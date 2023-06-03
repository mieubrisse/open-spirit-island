package main

import (
	"fmt"
	"github.com/mieubrisse/open-spirit-island/decks/fear"
	"github.com/mieubrisse/open-spirit-island/decks/invader"
	"github.com/mieubrisse/open-spirit-island/game_state"
	"github.com/mieubrisse/open-spirit-island/game_state/invaders"
	"github.com/mieubrisse/open-spirit-island/game_state/island"
	"github.com/mieubrisse/open-spirit-island/game_state/player"
)

func main() {
	fearDeck := make([]fear.FearCard, 9)
	for i := 0; i < 9; i++ {
		fearDeck[i] = fear.NewDummyFearCard()
	}

	invaderDeck := make([]invader.InvaderCard, 12)
	for i := 0; i < 12; i++ {
		invaderDeck[i] = invader.NewDummyInvaderCard()
	}

	invaderBoardState := invaders.InvaderBoardState{
		UnearnedFear:          4,
		EarnedFear:            0,
		TerrorLevelThresholds: []int{3, 3, 3},
		UnearnedFearCards:     fearDeck,
		EarnedFearCards:       make([]fear.FearCard, 0),
		RemainingInvaderDeck:  invaderDeck,
		BuildSlot: invaders.MaybeInvaderCard{
			IsCardPresent: false,
			MaybeCard:     invader.InvaderCard{},
		},
		RavageSlot: invaders.MaybeInvaderCard{
			IsCardPresent: false,
			MaybeCard:     invader.InvaderCard{},
		},
		InvaderDeckDiscard: []invader.InvaderCard{},
	}

	// TODO multiple players
	/*
		spiritBoardState := player.SpiritBoardState{
			SpiritPhaseOptions: [][]action.Action{
				{},
			},
			TopTrack: []action.Action{
				action.NewDummyAction(),
			},
			TopTrackRevealed: 1,
			BottomTrack: []action.Action{
				action.NewDummyAction(),
			},
			BottomTrackRevealed: 1,
		}

	*/
	playerState := player.PlayerState{
		Energy: 0,
		// SpiritBoardState: spiritBoardState,
	}

	// TODO multiple boards
	boardState := island.NewBoardA()
	boardState.AddPresence(1)

	gameState := game_state.GameState{
		InvaderState: invaderBoardState,
		PlayerState:  playerState,
		BoardState:   boardState,
	}

	fmt.Println(gameState.String())
}
