package player

import (
	"fmt"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/power"
	"strings"
)

type PlayerState struct {
	Energy int

	// TODO card plays

	// Represents the elements the player has as a result of their plays
	Elements map[power.Element]int

	Hand []power.PowerCard

	/*
		Played set.Of[power.PowerCardTransitionsID]

		Discard set.Of[power.PowerCardTransitionsID]
	*/
}

func (state PlayerState) String() string {
	lines := []string{
		fmt.Sprintf("⚡ %d", state.Energy),
	}

	for _, element := range power.ElementValues() {

	}

	return strings.Join(lines, "\n")
}
