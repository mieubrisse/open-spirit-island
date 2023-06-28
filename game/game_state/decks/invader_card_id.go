package decks

type InvaderCardID int

const (
	// Phase 1
	Mountain InvaderCardID = iota
	Jungle
	Wetlands
	Desert

	// Phase 2
	MountainAndAdversary
	JungleAndAdversary
	WetlandsAndAdversary
	DesertAndAdversary
	CoastalLands

	// Phase 3
	MountainAndJungle
	MountainAndWetlands
	MountainAndDesert
	JungleAndWetlands
	JungleAndDesert
	WetlandsAndDesert
)
