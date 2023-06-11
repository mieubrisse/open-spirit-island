package power

import (
	"fmt"
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/power/transition_ids"
	"strings"
)

type PowerCard struct {
	Title string

	Cost int

	Speed PowerCardSpeed

	Elements set.Of[Element]

	// TODO I'd really like the Transitions themselves be the rendering
	FlavorText string

	TransitionsID transition_ids.PowerCardTransitionsID
}

// TODO maybe move somewhere else???
func (p PowerCard) String() string {
	fmt.Println(fmt.Sprintf("%d⚡ %s", p.Cost, p.Title))

	elementSymbols := make([]string, 0, len(p.Elements))
	for _, element := range ElementValues() {
		if p.Elements.Has(element) {
			elementSymbols = append(elementSymbols, ElementSymbols[element])
		}
	}
	fmt.Println(strings.Join(elementSymbols, ""))

	//TODO implement me
	panic("implement me")
}
