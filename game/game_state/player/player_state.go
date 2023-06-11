package player

import (
	"fmt"
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/power"
	"strings"
)

type PlayerState struct {
	Energy int

	// Represents the elements the player has as a result of their plays
	Elements map[power.Element]int

	Hand set.Of[power.PowerCardTransitionsID]

	Played set.Of[power.PowerCardTransitionsID]

	Discard set.Of[power.PowerCardTransitionsID]

	// TODO hand
	// TODO cards played
	// TODO discard
}

func (state PlayerState) String() string {
	lines := []string{
		fmt.Sprintf("⚡ %d", state.Energy),
	}
	return strings.Join(lines, "\n")
}
