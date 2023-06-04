package land_state

import (
	"fmt"
	"github.com/mieubrisse/open-spirit-island/game_state/island/land_type"
	"strings"
)

type LandState struct {
	LandType land_type.LandType

	NumExplorers int
	NumTowns     int
	NumCities    int

	NumDahan int

	NumBlight int

	ExtraDefenseAdded int

	// TODO support multiple players
	NumPresence int
}

func (state LandState) String() string {
	lines := []string{
		fmt.Sprintf("Type: %s", state.LandType.String()),
		fmt.Sprintf("Presence: %d", state.NumPresence),
		fmt.Sprintf(
			"Explorers(%d) + Towns(%d) + Cities(%d)",
			state.NumExplorers,
			state.NumTowns,
			state.NumCities,
		),
		fmt.Sprintf("Dahan: %d", state.NumDahan),
		fmt.Sprintf("Blight: %d", state.NumBlight),
		fmt.Sprintf("Extra Defense: %d", state.ExtraDefenseAdded),
	}
	return strings.Join(lines, "\n")
}
