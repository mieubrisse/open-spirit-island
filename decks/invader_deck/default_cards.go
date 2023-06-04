package invader_deck

import (
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game_state/island/filter"
	"github.com/mieubrisse/open-spirit-island/game_state/island/land_type"
	"math"
)

// Vanilla game Stage 1 card selector
func NewSingleTypeInvaderCard(landType land_type.LandType) InvaderCard {
	if landType == land_type.Ocean {
		panic("Cannot have an invader card for " + land_type.Ocean.String())
	}

	return InvaderCard{
		TargetedLandSelector:  getInvaderFilterMatchingTypes(landType),
		IsAdversaryActionCard: false,
		HumanReadableStr:      landType.String(),
	}
}

// Vanilla game Stage 2 card selector
func NewSingleTypeAndAdversaryInvaderCard(landType land_type.LandType) InvaderCard {
	if landType == land_type.Ocean {
		panic("Cannot have an invader card for " + land_type.Ocean.String())
	}

	return InvaderCard{
		TargetedLandSelector:  getInvaderFilterMatchingTypes(landType),
		IsAdversaryActionCard: true,
		HumanReadableStr:      landType.String() + "+Adversary",
	}
}

func NewCoastalLandsInvaderCard() InvaderCard {
	return InvaderCard{
		TargetedLandSelector: filter.IslandFilter{
			SourceFilter: filter.LandFilter{
				LandTypes: set.New(land_type.Ocean),
			},
			MinRange: 1,
			MaxRange: 1,
		},
		IsAdversaryActionCard: false,
		HumanReadableStr:      "Coastal Lands",
	}
}

// Vanilla game Stage 3 land
func NewDoubleTypeInvaderCard(type1 land_type.LandType, type2 land_type.LandType) InvaderCard {
	return InvaderCard{
		TargetedLandSelector:  getInvaderFilterMatchingTypes(type1, type2),
		IsAdversaryActionCard: false,
		HumanReadableStr:      type1.String() + "+" + type2.String(),
	}
}

// NOTE: more custom invader cards could easily be created here!

// ====================================================================================================
//                                   Private Helper Functions
// ====================================================================================================

func getInvaderFilterMatchingTypes(targetedTypes ...land_type.LandType) filter.IslandFilter {
	return filter.IslandFilter{
		SourceFilter: filter.LandFilter{
			// Technically not necessary, but means that we only need to calculate distance from one land, rather than all
			LandTypes: set.New(land_type.Ocean),
		},
		MinRange: 1,
		MaxRange: math.MaxInt,
		TargetFilter: filter.LandFilter{
			LandTypes: set.New(targetedTypes...),
		},
	}
}
