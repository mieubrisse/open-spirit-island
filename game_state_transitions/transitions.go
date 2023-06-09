package game_state_transitions

import (
	"fmt"
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game_state"
	"github.com/mieubrisse/open-spirit-island/game_state/decks/power"
	"github.com/mieubrisse/open-spirit-island/game_state/island/filter"
	"github.com/mieubrisse/open-spirit-island/game_state/island/land_type"
	"github.com/mieubrisse/open-spirit-island/input"
	"sort"
)

type GameStateTransition func(state game_state.GameState) game_state.GameState

var ReclaimAllCardsTransition = func(state game_state.GameState) game_state.GameState {
	newPlayerState := state.PlayerState

	newPlayerState.Hand.Add(newPlayerState.Discard.Slice()...)
	newPlayerState.Discard = set.New[power.PowerCardID]()

	state.PlayerState = newPlayerState
	return state
}

func NewNormalAddPresenceTransition(addRange int) GameStateTransition {
	return func(state game_state.GameState) game_state.GameState {
		landIdxOptionsSet := state.BoardState.FilterLands(filter.IslandFilter{
			SourceFilter: filter.LandFilter{
				PresenceMin: 1,
			},
			MinRange: 0,
			MaxRange: addRange,
			TargetFilter: filter.LandFilter{
				LandTypes: land_type.NonOceanLandTypes,
			},
		})
		landIdxOptions := landIdxOptionsSet.Slice()
		sort.Ints(landIdxOptions)

		options := make([]string, len(landIdxOptions))
		for idx, landIdx := range landIdxOptions {
			land := state.BoardState.Lands[landIdx]
			options[idx] = fmt.Sprintf("%s #%d (%d Presence)", land.LandType.String(), landIdx, land.NumPresence)
		}

		selection := input.GetSelectionFromOptions("Choose a land to add Presence to:", options)
		selectedLandIdx := landIdxOptions[selection]

		state.BoardState.Lands[selectedLandIdx].NumPresence++

		return state
	}
}
