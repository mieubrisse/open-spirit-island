package effects

import (
	"github.com/mieubrisse/open-spirit-island/game/game_state"
)

// An effect corresponds to a line of instruction - be it on a Power card,
// on a Fear card, etc.
type Effect struct {
	ReadableStr string

	Applicator func(state game_state.GameState, targetLandIdx int) game_state.GameState
}

func (effect Effect) Apply(state game_state.GameState) game_state.GameState {
	return effect.Applicator(state)
}
