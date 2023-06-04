package main

import (
	"fmt"
	"github.com/mieubrisse/open-spirit-island/decks/fear"
	"github.com/mieubrisse/open-spirit-island/decks/invader_deck"
	"github.com/mieubrisse/open-spirit-island/game_state"
	"github.com/mieubrisse/open-spirit-island/game_state/invader_board"
	"github.com/mieubrisse/open-spirit-island/game_state/island"
	"github.com/mieubrisse/open-spirit-island/game_state/player"
	"github.com/mieubrisse/open-spirit-island/game_state/status"
)

func main() {
	fearDeck := make([]fear.FearCard, 9)
	for i := 0; i < 9; i++ {
		fearDeck[i] = fear.NewDummyFearCard()
	}

	// TODO parameterize this based on adversary
	invaderDeck := []invader_deck.InvaderCard{
		// stage 1
		invader_deck.NewSingleTypeInvaderCard(island.Jungle),
		invader_deck.NewSingleTypeInvaderCard(island.Desert),
		invader_deck.NewSingleTypeInvaderCard(island.Wetlands),
		// stage 2
		invader_deck.NewSingleTypeAndAdversaryInvaderCard(island.Mountain),
		invader_deck.NewSingleTypeAndAdversaryInvaderCard(island.Wetlands),
		invader_deck.NewSingleTypeAndAdversaryInvaderCard(island.Desert),
		invader_deck.NewCoastalLandsInvaderCard(),
		// stage 3
		invader_deck.NewDoubleTypeInvaderCard(island.Mountain, island.Jungle),
		invader_deck.NewDoubleTypeInvaderCard(island.Jungle, island.Wetlands),
		invader_deck.NewDoubleTypeInvaderCard(island.Wetlands, island.Desert),
		invader_deck.NewDoubleTypeInvaderCard(island.Mountain, island.Desert),
		invader_deck.NewDoubleTypeInvaderCard(island.Jungle, island.Desert),
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
			MaybeCard:     invader_deck.InvaderCard{},
		},
		RavageSlot: invader_board.MaybeInvaderCard{
			IsCardPresent: false,
			MaybeCard:     invader_deck.InvaderCard{},
		},
		InvaderDeckDiscard: []invader_deck.InvaderCard{},
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

	// TODO presence-placing
	// TODO first explore

	fmt.Println(gameState.String())

	for gameState.GetStatus() == status.Undecided {
		fmt.Scanln()
		gameState = gameState.Advance()
		fmt.Println(gameState.String())
	}
}
