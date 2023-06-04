package power

import "github.com/mieubrisse/open-spirit-island/game_state"

type Power interface {
	Play(state game_state.GameState) game_state.GameState
}
