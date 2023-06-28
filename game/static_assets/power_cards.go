package static_assets

import (
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks"
	"github.com/mieubrisse/open-spirit-island/game/game_state/player"
)

type PowerCard struct {
	Title    string
	Elements set.Of[player.Element]
}

func (card PowerCard) String() {

}

// TODO unit test to ensure completeness
var PowerCards = map[decks.PowerCardID]PowerCard{
	decks.AYearOfPerfectStillness: {},
}
