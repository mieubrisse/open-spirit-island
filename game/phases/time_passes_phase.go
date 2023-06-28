package phases

import (
	"github.com/mieubrisse/open-spirit-island/game/game_state"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island"
)

func RunTimePassesPhase(state game_state.GameState) game_state.GameState {
	for _, land := range state.BoardState.Lands {
		// Heal everything
		for i := 0; i < len(land.CityHealth); i++ {
			land.CityHealth[i] = island.CityBaseHealth
		}
		for i := 0; i < len(land.TownHealth); i++ {
			land.TownHealth[i] = island.TownBaseHealth
		}
		for i := 0; i < len(land.ExplorerHealth); i++ {
			land.ExplorerHealth[i] = island.ExplorerBaseHealth
		}
		for i := 0; i < len(land.DahanHealth); i++ {
			land.DahanHealth[i] = island.DahanBaseHealth
		}

		// TODO remove defend tokens

		// TODO remove various skip tokens
	}

	// Move all played cards to discord
	state.PlayerState.Played

	state.Phase = game_state.SpiritGrow
	return state
}
