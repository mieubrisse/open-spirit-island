package game_state

import (
	"github.com/mieubrisse/open-spirit-island/game_state/invaders"
	"github.com/mieubrisse/open-spirit-island/game_state/island"
	"github.com/mieubrisse/open-spirit-island/game_state/player"
	"github.com/mieubrisse/open-spirit-island/game_state/status"
	"strings"
)

type GameState struct {
	/*
		MajorPowersDeck    []power.Power
		MajorPowersDiscord []power.Power

		MinorPowersDeck    []power.Power
		MinorPowersDiscard []power.Power
	*/

	InvaderState invaders.InvaderBoardState

	// TODO multiple players
	PlayerState player.PlayerState

	// TODO multiple boards
	BoardState island.IslandBoardState
}

func (g GameState) GetStatus() status.GameStatus {
	numExplorers := 0
	numTowns := 0
	numCities := 0
	numPresence := 0
	for _, land := range g.BoardState.Lands {
		numExplorers += land.NumExplorers
		numTowns += land.NumTowns
		numCities += land.NumCities
		numPresence += land.NumPresence
	}

	// Check if the user has won
	terrorLevel := g.InvaderState.GetTerrorLevel()
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
	// TODO

	// Invader deck loss
	if len(g.InvaderState.RemainingInvaderDeck) == 0 {
		return status.Defeat
	}

	return status.Undecided
}

func (g GameState) String() string {
	lines := []string{
		"========================================== INVADERS ==========================================",
		g.InvaderState.String(),
		"=========================================== BOARD ============================================",
		g.BoardState.String(),
		// TODO multiple players
		"=========================================== PLAYER ===========================================",
		g.PlayerState.String(),
	}

	currentStatus := g.GetStatus()

	header := []string{}
	if currentStatus != status.Undecided {
		header = []string{
			"!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!",
			"                                            " + strings.ToUpper(currentStatus.String()),
			"!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!",
		}
	}
	lines = append(header, lines...)

	return strings.Join(lines, "\n")
}

// TODO get the score
