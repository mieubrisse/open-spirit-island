package player

import (
	"fmt"
	"strings"
)

type PlayerState struct {
	Energy int

	// SpiritBoardState SpiritBoardState

	// TODO hand
	// TODO cards played
	// TODO discard
}

func (state PlayerState) String() string {
	lines := []string{
		fmt.Sprintf("Energy: %d", state.Energy),
	}
	return strings.Join(lines, "\n")
}
