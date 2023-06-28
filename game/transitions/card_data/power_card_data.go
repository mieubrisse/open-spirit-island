package card_data

import (
	"fmt"
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game/game_state/player"
	"github.com/mieubrisse/open-spirit-island/game/static_assets"
	"github.com/mieubrisse/open-spirit-island/game/transitions/effects"
	"strings"
)

type PowerCardData struct {
	Title string

	Cost int

	Speed PowerCardSpeed

	// TODO add range

	Elements set.Of[player.Element]

	Effects []effects.LandTargetingEffect
}

// TODO maybe move somewhere else???
func (p PowerCardData) String() string {
	energy := fmt.Sprintf("%d⚡", p.Cost)

	elementSymbols := make([]string, 0, len(p.Elements))
	for _, element := range player.ElementValues() {
		if p.Elements.Has(element) {
			elementSymbols = append(elementSymbols, static_assets.ElementSymbols[element])
		}
	}
	elements := strings.Join(elementSymbols, "")

	lines := []string{
		p.Title,
		energy + " and " + elements,
	}
	for _, effect := range p.Effects {
		lines = append(lines, effect.ReadableStr)
	}
	return strings.Join(lines, "\n")
}
