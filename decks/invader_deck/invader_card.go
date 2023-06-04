package invader_deck

import "github.com/mieubrisse/open-spirit-island/game_state/island"

type LandSelector func(board island.IslandBoardState) []int

// Represents a generic invader card
type InvaderCard struct {
	// TODO add the invasion phase??

	// Function for selecting targeted lands given an island board state
	// Returns indices in sorted order
	TargetedLandSelector LandSelector

	// True if it's an adversary action card
	isAdversaryActionCard bool

	humanReadableStr string
}

func (i InvaderCard) String() string {
	return i.humanReadableStr
}
