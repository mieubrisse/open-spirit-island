package power

import (
	"fmt"
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/power/transition_ids"
	"github.com/mieubrisse/open-spirit-island/game/game_state/player"
	"strings"
)

type PowerCard struct {
	Title string

	Cost int

	Speed PowerCardSpeed

	Elements set.Of[player.Element]

	// TODO I'd really like the Transitions themselves be the rendering
	FlavorText string

	TransitionsID transition_ids.PowerCardTransitionsID
}

// TODO maybe move somewhere else???
func (p PowerCard) String() string {
	energy := fmt.Sprintf("%d⚡", p.Cost)

	elementSymbols := make([]string, 0, len(p.Elements))
	for _, element := range player.ElementValues() {
		if p.Elements.Has(element) {
			elementSymbols = append(elementSymbols, player.ElementSymbols[element])
		}
	}
	elements := strings.Join(elementSymbols, "")

	lines := []string{
		p.Title,
		energy + " and " + elements,
		p.FlavorText,
	}
	return strings.Join(lines, "\n")
}
