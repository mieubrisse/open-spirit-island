package player

import (
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks"
)

type PlayerState struct {
	Energy int

	CardPlaysRemaining int

	// Represents the elements the player has as a result of their plays
	NumElements map[Element]int

	Hand []decks.PowerCardID

	Played []decks.PowerCardID

	Discard []decks.PowerCardID
}
