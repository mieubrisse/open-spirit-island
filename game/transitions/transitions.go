package transitions

import (
	"github.com/mieubrisse/open-spirit-island/game/game_state"
)

type GameStateTransition struct {
	ReadableStr        string
	TransitionFunction func(state game_state.GameState) game_state.GameState
}

func (g GameStateTransition) String() string {
	return g.ReadableStr
}
