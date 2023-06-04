package invader_deck

import (
	"github.com/mieubrisse/open-spirit-island/game_state/island"
	"github.com/mieubrisse/open-spirit-island/game_state/island/land_type"
)

// Vanilla game Stage 1 card selector
func NewSingleTypeInvaderCard(landType land_type.LandType) InvaderCard {
	if landType == land_type.Ocean {
		panic("Cannot have an invader card for " + land_type.Ocean.String())
	}

	return InvaderCard{
		TargetedLandSelector: getSelectorMatchingTypes(map[land_type.LandType]bool{
			landType: true,
		}),
		isAdversaryActionCard: false,
		humanReadableStr:      landType.String(),
	}
}

// Vanilla game Stage 2 card selector
func NewSingleTypeAndAdversaryInvaderCard(landType land_type.LandType) InvaderCard {
	if landType == land_type.Ocean {
		panic("Cannot have an invader card for " + land_type.Ocean.String())
	}

	return InvaderCard{
		TargetedLandSelector: getSelectorMatchingTypes(map[land_type.LandType]bool{
			landType: true,
		}),
		isAdversaryActionCard: true,
		humanReadableStr:      landType.String() + "+Adversary",
	}
}

func NewCoastalLandsInvaderCard() InvaderCard {
	return InvaderCard{
		TargetedLandSelector: func(board island.IslandBoardState) []int {
			// 0 is always the ocean (little bit of forbidden knowledge though)
			return board.GetAdjacentLands(0)
		},
		isAdversaryActionCard: false,
		humanReadableStr:      "Coastal Lands",
	}
}

// Vanilla game Stage 3 land
func NewDoubleTypeInvaderCard(type1 land_type.LandType, type2 land_type.LandType) InvaderCard {
	return InvaderCard{
		TargetedLandSelector: getSelectorMatchingTypes(map[land_type.LandType]bool{
			type1: true,
			type2: true,
		}),
		isAdversaryActionCard: false,
		humanReadableStr:      type1.String() + "+" + type2.String(),
	}
}

// NOTE: more custom invader cards could easily be created here!

// ====================================================================================================
//                                   Private Helper Functions
// ====================================================================================================

func getSelectorMatchingTypes(targetedTypes map[land_type.LandType]bool) LandSelector {
	return func(board island.IslandBoardState) []int {
		result := make([]int, 0)
		for idx, land := range board.Lands {
			if _, found := targetedTypes[land.LandType]; found {
				result = append(result, idx)
			}
		}
		return result
	}
}
