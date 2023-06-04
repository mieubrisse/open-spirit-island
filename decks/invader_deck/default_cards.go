package invader_deck

import "github.com/mieubrisse/open-spirit-island/game_state/island"

// Vanilla game Stage 1 card selector
func NewSingleTypeInvaderCard(landType island.LandType) InvaderCard {
	if landType == island.Ocean {
		panic("Cannot have an invader card for " + island.Ocean.String())
	}

	return InvaderCard{
		TargetedLandSelector: getSelectorMatchingTypes(map[island.LandType]bool{
			landType: true,
		}),
		isAdversaryActionCard: false,
		humanReadableStr:      landType.String(),
	}
}

// Vanilla game Stage 2 card selector
func NewSingleTypeAndAdversaryInvaderCard(landType island.LandType) InvaderCard {
	if landType == island.Ocean {
		panic("Cannot have an invader card for " + island.Ocean.String())
	}

	return InvaderCard{
		TargetedLandSelector: getSelectorMatchingTypes(map[island.LandType]bool{
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
func NewDoubleTypeInvaderCard(type1 island.LandType, type2 island.LandType) InvaderCard {
	return InvaderCard{
		TargetedLandSelector: getSelectorMatchingTypes(map[island.LandType]bool{
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

func getSelectorMatchingTypes(targetedTypes map[island.LandType]bool) LandSelector {
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
