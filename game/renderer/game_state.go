package renderer

import (
	"github.com/mieubrisse/open-spirit-island/game/game_state"
	"github.com/mieubrisse/open-spirit-island/game/game_state/status"
	"github.com/mieubrisse/open-spirit-island/game/renderer/invader_board"
	"strings"
)

func RenderGameState(state game_state.GameState) string {
	lines := []string{
		invader_board.RenderInvaderBoardState(state.InvaderState),
		state.BoardState.String(),
		// TODO multiple players
		"                                  PLAYER",
		state.PlayerState.String(),
	}

	currentStatus := state.GetStatus()

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
