package game_state

import (
	"github.com/mieubrisse/open-spirit-island/game_state/invaders"
	"github.com/mieubrisse/open-spirit-island/game_state/island"
	"github.com/mieubrisse/open-spirit-island/game_state/player"
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

func (g GameState) String() string {
	lines := []string{
		"====================== INVADERS ====================",
		g.InvaderState.String(),
		"======================= BOARD ======================",
		g.BoardState.String(),
		// TODO multiple players
		"======================= PLAYER ======================",
		g.PlayerState.String(),
	}

	return strings.Join(lines, "\n")
}
