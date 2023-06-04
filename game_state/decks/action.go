package decks

import "github.com/mieubrisse/open-spirit-island/game_state"

type Action func(state game_state.GameState) game_state.GameState
