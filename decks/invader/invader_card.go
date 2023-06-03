package invader

import "github.com/mieubrisse/open-spirit-island/game_state/island"

// Represents a generic invader card
type InvaderCard struct {
	// The unique ID representing this particular invader card
	id int

	// TODO add the invasion phase??

	// Function for selecting targeted lands given an island board state
	targetedLandSelector func(board island.IslandBoardState) map[int]bool

	// True if it's an adversary action card
	isAdversaryActionCard bool

	humanReadableStr string
}

func NewDummyInvaderCard() InvaderCard {
	return InvaderCard{
		id: 0,
		targetedLandSelector: func(board island.IslandBoardState) map[int]bool {
			result := map[int]bool{}
			for idx := range board.Lands {
				// INVADE ALL THE LANDS!!!
				result[idx] = true
			}
			return result
		},
		isAdversaryActionCard: false,
		humanReadableStr:      "EVERYTHING!",
	}
}

func (i InvaderCard) String() string {
	return i.humanReadableStr
}
