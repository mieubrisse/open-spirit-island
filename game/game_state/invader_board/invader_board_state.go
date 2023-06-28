package invader_board

import (
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/blighted_island"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/fear"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/invader_deck"
)

type MaybeInvaderCard struct {
	IsCardPresent bool
	MaybeCard     invader_deck.InvaderCard
}

type InvaderBoardState struct {
	UnearnedFear int
	EarnedFear   int

	// The number of fear cards needed to earn Terror Level 1/2/3/4
	TerrorLevelThresholds []int

	UnearnedFearCards []fear.FearCard // New cards are popped from the first element
	EarnedFearCards   []fear.FearCard // Processed from left to right

	// TODO blight card & pool!!

	BlightedIslandCard blighted_island.BlightedIslandCard
	IsBlightedIsland   bool
	BlightPool         int

	RemainingInvaderDeck []invader_deck.InvaderCard
	BuildSlot            MaybeInvaderCard
	RavageSlot           MaybeInvaderCard
	InvaderDeckDiscard   []invader_deck.InvaderCard
}

func (state InvaderBoardState) AddFear(fear int) InvaderBoardState {
	for i := 0; i < fear; i++ {
		state.UnearnedFear -= 1
		state.EarnedFear += 1

		// Earned a fear card
		if state.UnearnedFear == 0 {
			state.UnearnedFear = state.EarnedFear
			state.EarnedFear = 0

			if len(state.UnearnedFearCards) > 0 {
				state.EarnedFearCards = append(state.EarnedFearCards, state.UnearnedFearCards[0])
				state.UnearnedFearCards = state.UnearnedFearCards[1:]
			}
		}
	}
	return state
}

// Terror level
func (state InvaderBoardState) GetTerrorLevel() int {
	numEarnedFearCards := len(state.EarnedFearCards)
	highestIdxReached := 0
	for idx, terrorLevelThreshold := range state.TerrorLevelThresholds {
		if numEarnedFearCards >= terrorLevelThreshold {
			highestIdxReached = idx
		}
	}
	return highestIdxReached + 1 // Because terror level is 1-indexed
}

func (state InvaderBoardState) AdvanceInvaderCards() InvaderBoardState {
	if len(state.RemainingInvaderDeck) == 0 {
		return state
	}

	state.RavageSlot = state.BuildSlot
	state.BuildSlot = MaybeInvaderCard{
		IsCardPresent: true,
		MaybeCard:     state.RemainingInvaderDeck[0],
	}
	state.RemainingInvaderDeck = state.RemainingInvaderDeck[1:]
	return state
}
