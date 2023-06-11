package land_state

import (
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/land_type"
)

type LandState struct {
	LandType land_type.LandType

	// TODO something to reflect that replaced invaders keep the damage they have
	ExplorerHealth []int
	TownHealth     []int
	CityHealth     []int

	DahanHealth []int

	NumBlight int

	ExtraDefenseAdded int

	// TODO support multiple players
	NumPresence int
}
