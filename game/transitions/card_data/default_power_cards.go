package card_data

import (
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks"
	"github.com/mieubrisse/open-spirit-island/game/game_state/player"
	"github.com/mieubrisse/open-spirit-island/game/transitions/effects"
)

// TODO unit test to ensure that we have data for all the cards defined in PowerCardID
var DefaultPowerCards = map[decks.PowerCardID]PowerCardData{
	// Vital Strength of the Earth unique powers
	decks.DrawOfTheFruitfulEarth: {
		Title:    "Draw Of The Fruitful Earth",
		Cost:     1,
		Speed:    Slow,
		Elements: set.New(player.Earth, player.Plant, player.Animal),
		Effects: []effects.LandTargetingEffect{
			effects.NewGatherObjectEffect(0, 2, effects.Dahan),
			effects.NewGatherObjectEffect(0, 2, effects.Explorer),
		},
	},
}
