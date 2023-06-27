package player

import (
	"fmt"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/power"
	"strings"
)

type PlayerState struct {
	Energy int

	CardPlaysRemaining int

	// Represents the elements the player has as a result of their plays
	NumElements map[power.Element]int

	Hand []power.PowerCard

	Played []power.PowerCard

	Discard []power.PowerCard
}

func (state PlayerState) String() string {
	energy := fmt.Sprintf("⚡ %d", state.Energy)

	elementCountStrs := make([]string, len(power.ElementValues()))
	for i, element := range power.ElementValues() {
		elementSymbol := power.ElementSymbols[element]
		count := state.NumElements[element]
		elementCountStrs[i] = fmt.Sprintf("%s %d", elementSymbol, count)
	}
	elementCounts := strings.Join(elementCountStrs, "   ")

	talliesLine := energy + "   and   " + elementCounts

	lines := []string{
		talliesLine,
	}

	lines = append(lines, "-------- HAND ---------")
	for _, card := range state.Hand {
		lines = append(lines, card.String())
	}
	lines = append(lines, "-------- PLAYED ---------")
	for _, card := range state.Played {
		lines = append(lines, card.String())
	}

	return strings.Join(lines, "\n")
}
