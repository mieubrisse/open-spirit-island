package default_power_cards

import (
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/power"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/power/transition_ids"
	"github.com/mieubrisse/open-spirit-island/game/game_state/player"
)

var DrawOfTheFruitfulEarth = power.PowerCard{
	Title:         "Draw Of The Fruitful Earth",
	Cost:          1,
	Speed:         power.Slow,
	Elements:      set.New(player.Earth, player.Plant, player.Animal),
	FlavorText:    "Gather up to 2 🧍\nGather up to 2 🛖",
	TransitionsID: transition_ids.DrawOfTheFruitfulEarth,
}
