package game_state

import (
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/blighted_island"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/fear"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/invader_deck"
	"github.com/mieubrisse/open-spirit-island/game/game_state/invader_board"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/land_type"
	"github.com/mieubrisse/open-spirit-island/game/game_state/player"
)

func NewTestGame() GameState {
	fearDeck := make([]fear.FearCard, 9)
	for i := 0; i < 9; i++ {
		fearDeck[i] = fear.NewDummyFearCard()
	}

	// TODO parameterize this based on adversary
	invaderDeck := []invader_deck.InvaderCard{
		// stage 1
		invader_deck.NewSingleTypeInvaderCard(land_type.Jungle),
		invader_deck.NewSingleTypeInvaderCard(land_type.Desert),
		invader_deck.NewSingleTypeInvaderCard(land_type.Wetlands),
		// stage 2
		invader_deck.NewSingleTypeAndAdversaryInvaderCard(land_type.Mountain),
		invader_deck.NewSingleTypeAndAdversaryInvaderCard(land_type.Wetlands),
		invader_deck.NewSingleTypeAndAdversaryInvaderCard(land_type.Desert),
		invader_deck.NewCoastalLandsInvaderCard(),
		// stage 3
		invader_deck.NewDoubleTypeInvaderCard(land_type.Mountain, land_type.Jungle),
		invader_deck.NewDoubleTypeInvaderCard(land_type.Jungle, land_type.Wetlands),
		invader_deck.NewDoubleTypeInvaderCard(land_type.Wetlands, land_type.Desert),
		invader_deck.NewDoubleTypeInvaderCard(land_type.Mountain, land_type.Desert),
		invader_deck.NewDoubleTypeInvaderCard(land_type.Jungle, land_type.Desert),
		// invader_deck.NewDoubleTypeInvaderCard(island.Mountain, island.Wetlands),
	}

	// TODO increase by number of players
	invaderBoardState := invader_board.InvaderBoardState{
		UnearnedFear:          4,
		EarnedFear:            0,
		TerrorLevelThresholds: []int{3, 3, 3},
		UnearnedFearCards:     fearDeck,
		EarnedFearCards:       make([]fear.FearCard, 0),
		BlightedIslandCard:    blighted_island.NewBlightedIslandCardIDontRememberRightNow(),
		IsBlightedIsland:      false,
		BlightPool:            2,
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

	playerState := player.PlayerState{
		Energy: 0,
	}

	boardState := island.NewBoardA()

	return GameState{
		InvaderState: invaderBoardState,
		PlayerState:  playerState,
		BoardState:   boardState,
	}
}
