package invader_deck

import "github.com/mieubrisse/open-spirit-island/game_state/island/filter"

// Represents a generic invader card
type InvaderCard struct {
	// TODO add the invasion phase??

	TargetedLandSelector filter.IslandFilter

	// True if it's an adversary action card
	IsAdversaryActionCard bool

	HumanReadableStr string
}

func (i InvaderCard) String() string {
	return i.HumanReadableStr
}
