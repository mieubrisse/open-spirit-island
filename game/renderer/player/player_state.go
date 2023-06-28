package player

import (
	"fmt"
	"github.com/mieubrisse/open-spirit-island/game/game_state/player"
	"github.com/mieubrisse/open-spirit-island/game/static_assets"
	"strings"
)

func RenderPlayerState(state player.PlayerState) string {
	energy := fmt.Sprintf("⚡ %d", state.Energy)

	elementCountStrs := make([]string, len(player.ElementValues()))
	for i, element := range player.ElementValues() {
		elementSymbol := static_assets.ElementSymbols[element]
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
