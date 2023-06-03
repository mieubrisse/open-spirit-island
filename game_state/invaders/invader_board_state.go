package invaders

import (
	"fmt"
	"github.com/mieubrisse/open-spirit-island/decks/fear"
	"github.com/mieubrisse/open-spirit-island/decks/invader"
	"math"
	"strings"
)

type MaybeInvaderCard struct {
	IsCardPresent bool
	MaybeCard     invader.InvaderCard
}

type InvaderBoardState struct {
	UnearnedFear int
	EarnedFear   int

	// The number of fear cards needed to earn Terror Level 1/2/3/4
	TerrorLevelThresholds []int

	UnearnedFearCards []fear.FearCard // New cards are popped from the first element
	EarnedFearCards   []fear.FearCard // Processed from left to right

	RemainingInvaderDeck []invader.InvaderCard
	BuildSlot            MaybeInvaderCard
	RavageSlot           MaybeInvaderCard
	InvaderDeckDiscard   []invader.InvaderCard
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

func (state InvaderBoardState) String() string {
	ravageLineContent := "<none>"
	if state.RavageSlot.IsCardPresent {
		ravageLineContent = state.RavageSlot.MaybeCard.String()
	}

	buildLineContent := "<none>"
	if state.BuildSlot.IsCardPresent {
		buildLineContent = state.BuildSlot.MaybeCard.String()
	}

	lines := []string{
		fmt.Sprintf("Terror Level: %d", state.GetTerrorLevel()),
		fmt.Sprintf("Fear Till Next Card: %d", state.UnearnedFear),
		fmt.Sprintf("Fear Cards Till Next Terror Level: %d", state.countFearCardsTillNextTerrorLevel()),
		fmt.Sprintf("Earned Fear Cards: %d", len(state.EarnedFearCards)),
		fmt.Sprintf(
			"Ravage(%s) <- Build(%s) <- Deck(%d)",
			ravageLineContent,
			buildLineContent,
			len(state.RemainingInvaderDeck),
		),
	}
	return strings.Join(lines, "\n")
}

func (state InvaderBoardState) countFearCardsTillNextTerrorLevel() int {
	numEarnedFearCards := len(state.EarnedFearCards)
	numFearCardsTillNextTerrorLevel := math.MaxInt
	for _, terrorLevelThreshold := range state.TerrorLevelThresholds {
		numCardsTillTerrorLevel := terrorLevelThreshold - numEarnedFearCards
		if numCardsTillTerrorLevel > 0 && numCardsTillTerrorLevel < numFearCardsTillNextTerrorLevel {
			numFearCardsTillNextTerrorLevel = numCardsTillTerrorLevel
		}
	}
	return numFearCardsTillNextTerrorLevel
}
