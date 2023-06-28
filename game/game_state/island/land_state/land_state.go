package land_state

import (
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/land_type"
)

type LandState struct {
	LandType land_type.LandType

	// NOTE: Spirit Island bookkeeping uses *damage taken*, NOT health, as a way to track death
	// Meaning, if Dahan temporarily have 5 health, one takes 4 damage, and it then gets moved
	// to a land where Dahan have 2 health, the Dahan will die

	ExplorerDamageTaken []int // Len is a count of the number of explorers
	ExplorerHealth      int

	TownDamageTaken []int // Len is a count of the number of towns
	TownHealth      int

	CityDamageTaken []int // Len is a count of the number of cities
	CityHealth      int

	DahanDamageTaken []int // Len is a count of the number of Dahan
	DahanHealth      int

	NumBlight int

	ExtraDefenseAdded int

	// TODO support multiple players
	NumPresence int
}
