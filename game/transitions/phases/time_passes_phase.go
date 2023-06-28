package phases

import (
	"github.com/mieubrisse/open-spirit-island/game/game_state"
)

func RunTimePassesPhase(state game_state.GameState) game_state.GameState {
	for _, land := range state.BoardState.Lands {
		// Heal everything
		// NOTE: I couldn't find a ruling on whether an expiring healh buff would kill objects
		// who've taken more-than-normal amount of damage
		// I'm going with the sensible thing, which is that if the object survives, it will heal
		for i := 0; i < len(land.CityDamageTaken); i++ {
			land.CityDamageTaken[i] = 0
		}
		for i := 0; i < len(land.TownDamageTaken); i++ {
			land.TownDamageTaken[i] = 0
		}
		for i := 0; i < len(land.ExplorerDamageTaken); i++ {
			land.ExplorerDamageTaken[i] = 0
		}
		for i := 0; i < len(land.DahanDamageTaken); i++ {
			land.DahanDamageTaken[i] = 0
		}

		// TODO remove defend tokens

		// TODO remove various skip tokens
	}

	// TODO Move all played cards to discord

	// TODO kill any Dahan or invaders whose health total has dropped

	state.Phase = game_state.SpiritGrow
	return state
}
