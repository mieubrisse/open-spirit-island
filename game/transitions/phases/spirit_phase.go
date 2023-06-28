package phases

import (
	"github.com/mieubrisse/open-spirit-island/game/game_state"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks"
	"github.com/mieubrisse/open-spirit-island/game/transitions"
	input2 "github.com/mieubrisse/open-spirit-island/game/transitions/input"
	"strings"
)

func RunSpiritGrowPhase(state game_state.GameState) game_state.GameState {
	// TODO get the growth options from the player's board

	// Growth
	// TODO other spirit choices
	growthChoices := [][]transitions.GameStateTransition{
		{
			transitions.ReclaimAllCardsTransition,
			transitions.NewNormalAddPresenceTransition(2),
		},
		{
			// TODO gain a power card
			transitions.NewNormalAddPresenceTransition(0),
		},
		{
			transitions.NewNormalAddPresenceTransition(1),
			transitions.NewGainEnergyTransition(2),
		},
	}

	growthChoicesStrs := make([]string, len(growthChoices))
	for i, choiceSubcomponents := range growthChoices {
		choiceComponentStrs := make([]string, len(choiceSubcomponents))
		for j, subcomponent := range choiceSubcomponents {
			choiceComponentStrs[j] = subcomponent.String()
		}
		growthChoicesStrs[i] = strings.Join(choiceComponentStrs, "   |   ")
	}

	selectionIdx := input2.GetSingleSelection("Select a Growth option:", growthChoicesStrs)

	growthSelection := growthChoices[selectionIdx]

	for _, transition := range growthSelection {
		state = transition.TransitionFunction(state)
	}

	state.Phase = game_state.SpiritGainTrackBenefits

	return state
}

func RunSpiritGainTrackBenefitsPhase(state game_state.GameState) game_state.GameState {
	// TODO do things based on what's uncovered on the track!!

	// TODO elemental income

	// TODO card plays from the track
	state.PlayerState.CardPlaysRemaining = 3

	return state
}

func RunSpiritPlayPowerPhase(state game_state.GameState) game_state.GameState {
	if state.PlayerState.CardPlaysRemaining == 0 {
		state.Phase = game_state.FastPower
	}

	oldHand := state.PlayerState.Hand

	// TODO correctly calculate card plays
	selectedHandCardIdxs, totalEnergyCost := input2.PlayCards(oldHand, state.PlayerState.Energy, state.PlayerState.CardPlaysRemaining)
	if len(selectedHandCardIdxs) == 0 {
		state.Phase = game_state.FastPower
		return state
	}

	newPlayed := make([]decks.PowerCardID, 0, len(selectedHandCardIdxs))
	newHand := make([]decks.PowerCardID, 0, len(oldHand)-len(selectedHandCardIdxs))
	for oldHandIdx := range state.PlayerState.Hand {
		card := oldHand[oldHandIdx]
		if selectedHandCardIdxs.Has(oldHandIdx) {
			newPlayed = append(newPlayed, card)
		} else {
			newHand = append(newHand, card)
		}
	}
	state.PlayerState.Energy -= totalEnergyCost
	state.PlayerState.CardPlaysRemaining -= len(selectedHandCardIdxs)
	state.PlayerState.Hand = newHand
	state.PlayerState.Played = newPlayed

	return state
}
