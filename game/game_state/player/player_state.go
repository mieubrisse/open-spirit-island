package player

import (
	"fmt"
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/power"
	"strings"
)

type PlayerState struct {
	Energy int

	Hand set.Of[power.PowerCardID]

	Played set.Of[power.PowerCardID]

	Discard set.Of[power.PowerCardID]

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
