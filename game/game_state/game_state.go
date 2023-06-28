package game_state

import (
	"github.com/mieubrisse/open-spirit-island/game/game_state/invader_board"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island"
	"github.com/mieubrisse/open-spirit-island/game/game_state/player"
	"github.com/mieubrisse/open-spirit-island/game/game_state/status"
)

type GameState struct {
	Phase GamePhase

	/*
		MajorPowersDeck    []power.Power
		MajorPowersDiscord []power.Power

		MinorPowersDeck    []power.Power
		MinorPowersDiscard []power.Power
	*/

	InvaderState invader_board.InvaderBoardState

	// TODO multiple players
	PlayerState player.PlayerState

	// TODO multiple boards
	BoardState island.IslandBoardState
}

func (state GameState) GetStatus() status.GameStatus {
	numExplorers := 0
	numTowns := 0
	numCities := 0
	numPresence := 0
	for _, land := range state.BoardState.Lands {
		numExplorers += len(land.ExplorerHealth)
		numTowns += len(land.TownHealth)
		numCities += len(land.CityHealth)
		numPresence += land.NumPresence
	}

	// Check if the user has won
	terrorLevel := state.InvaderState.GetTerrorLevel()
	victoryChecks := []bool{
		terrorLevel >= 4,
		terrorLevel == 3 && numCities == 0,
		terrorLevel == 2 && numCities == 0 && numTowns == 0,
		terrorLevel == 1 && numCities == 0 && numTowns == 0 && numExplorers == 0,
	}
	isFearVictory := false
	for _, check := range victoryChecks {
		isFearVictory = isFearVictory || check
	}
	if isFearVictory {
		return status.Victory
	}

	// Presence token loss
	if numPresence == 0 {
		return status.Defeat
	}

	// Blight loss
	if state.InvaderState.BlightPool == 0 {
		return status.Defeat
	}

	// Invader deck loss
	if len(state.InvaderState.RemainingInvaderDeck) == 0 {
		return status.Defeat
	}

	return status.Undecided
}
