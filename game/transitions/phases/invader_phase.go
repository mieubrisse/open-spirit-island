package phases

import (
	"fmt"
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game/game_state"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/filter"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/land_type"
	"github.com/mieubrisse/open-spirit-island/game/game_state/status"
	input2 "github.com/mieubrisse/open-spirit-island/game/transitions/input"
	"github.com/mieubrisse/open-spirit-island/game/utils"
	"sort"
)

// Runs the invasion phase
func RunInvaderPhase(state game_state.GameState) game_state.GameState {
	liveInvaderLandsFilter := filter.IslandFilter{
		TargetFilter: filter.LandFilter{
			// This is needed because Ocean's Hungry Grasp can drown invaders
			LandTypes:   land_type.NonOceanLandTypes,
			InvadersMin: 1,
		},
	}
	invaderLandIdxs := state.BoardState.FilterLands(liveInvaderLandsFilter)

	// Ravage
	ravageSlot := state.InvaderState.RavageSlot
	if ravageSlot.IsCardPresent {
		desiredRavageIdxs := state.BoardState.FilterLands(ravageSlot.MaybeCard.TargetedLandSelector)
		actualRavageIdxs := set.Intersect(invaderLandIdxs, desiredRavageIdxs)

		for landIdx := range actualRavageIdxs {
			ravageLand := state.BoardState.Lands[landIdx]

			// TODO ravage skips

			invaderDamage := island.ExplorerBaseDamage*len(ravageLand.ExplorerDamageTaken) + island.TownBaseDamage*len(ravageLand.TownDamageTaken) + island.CityBaseDamage*len(ravageLand.CityDamageTaken)

			// Dahan are attacked
			// From the rules, Dahan must be destroyed as efficiently as possible
			ravageLand.DahanDamageTaken = efficientlyDamageDahan(ravageLand.DahanDamageTaken, ravageLand.DahanHealth, invaderDamage)

			// Dahan strike back
			dahanDamage := island.DahanBaseDamage * len(ravageLand.DahanDamageTaken)
			ravageLand = input2.DamageInvaders(
				"Dahan",
				ravageLand,
				dahanDamage,
			)

			// TODO defend (Dahan and Land)

			if invaderDamage >= 2 {
				state = blightLandWithCascade(state, landIdx)
			}

			state.BoardState.Lands[landIdx] = ravageLand
		}
	}

	// Loss condition check (possible due to Blight allocation)
	if state.GetStatus() != status.Undecided {
		return state
	}

	// Build
	buildSlot := state.InvaderState.BuildSlot
	if buildSlot.IsCardPresent {
		desiredBuildIdxs := state.BoardState.FilterLands(buildSlot.MaybeCard.TargetedLandSelector)
		actualBuildIdxs := set.Intersect(invaderLandIdxs, desiredBuildIdxs)

		for landIdx := range actualBuildIdxs {
			buildLand := state.BoardState.Lands[landIdx]
			// TODO build skips

			if len(buildLand.TownDamageTaken) > len(buildLand.CityDamageTaken) {
				buildLand.CityDamageTaken = append(buildLand.CityDamageTaken, 0)
			} else {
				buildLand.TownDamageTaken = append(buildLand.TownDamageTaken, 0)
			}
			state.BoardState.Lands[landIdx] = buildLand
		}
	}

	// Explore
	coastalLandIdxs := state.BoardState.FilterLands(filter.NewCoastalLandsFilter())
	townExplorableLandIdxs := state.BoardState.FilterLands(filter.IslandFilter{
		SourceFilter: filter.LandFilter{
			TownsMin: 1,
		},
		MinRange: 0,
		MaxRange: 1,
		TargetFilter: filter.LandFilter{
			LandTypes: land_type.NonOceanLandTypes,
		},
	})
	cityExplorableLandIdxs := state.BoardState.FilterLands(filter.IslandFilter{
		SourceFilter: filter.LandFilter{
			CitiesMin: 1,
		},
		MinRange: 0,
		MaxRange: 1,
		TargetFilter: filter.LandFilter{
			LandTypes: land_type.NonOceanLandTypes,
		},
	})
	explorableLandIdx := set.Union(coastalLandIdxs, townExplorableLandIdxs, cityExplorableLandIdxs)

	explorerCard := state.InvaderState.RemainingInvaderDeck[0]
	desiredExploreLandIdxs := state.BoardState.FilterLands(explorerCard.TargetedLandSelector)

	toExploreLandIdxs := set.Intersect(explorableLandIdx, desiredExploreLandIdxs)

	for landIdx := range toExploreLandIdxs {
		// TODO explore skip

		exploreLand := state.BoardState.Lands[landIdx]
		exploreLand.ExplorerDamageTaken = append(exploreLand.ExplorerDamageTaken, 0)
		state.BoardState.Lands[landIdx] = exploreLand
	}

	// Advance invader slots
	state.InvaderState = state.InvaderState.AdvanceInvaderCards()

	state.Phase = game_state.SlowPower

	return state
}

// ====================================================================================================
//                                   Private Helper Functions
// ====================================================================================================

func blightLandWithCascade(state game_state.GameState, landIdx int) game_state.GameState {
	var shouldBlightCascade bool
	state, shouldBlightCascade = blightSingleLand(state, landIdx)

	sourceLandIdx := landIdx
	for shouldBlightCascade {
		sourceLand := state.BoardState.Lands[sourceLandIdx]

		adjacentLandIdxsSet := state.BoardState.FilterLands(filter.IslandFilter{
			SourceNumbers: set.New(sourceLandIdx),
			MinRange:      1,
			MaxRange:      1,
			TargetFilter:  filter.LandFilter{LandTypes: land_type.NonOceanLandTypes},
		})
		adjacentLandIdxs := adjacentLandIdxsSet.Slice()
		sort.Ints(adjacentLandIdxs)

		optionStrs := make([]string, len(adjacentLandIdxs))
		for i, adjacentIdx := range adjacentLandIdxs {
			adjacentLand := state.BoardState.Lands[adjacentIdx]
			optionStrs[i] = fmt.Sprintf("%s #%d (%d Blight, %d Presence)", adjacentLand.LandType, adjacentIdx, adjacentLand.NumBlight, adjacentLand.NumPresence)
		}

		selection := input2.GetSingleSelection(
			fmt.Sprintf("%s #%d is suffering a Blight cascade; select an adjacent land to spread Blight to:", sourceLand.LandType, sourceLandIdx),
			optionStrs,
		)
		selectedLandIdx := adjacentLandIdxs[selection]

		state, shouldBlightCascade = blightSingleLand(state, selectedLandIdx)

		sourceLandIdx = selectedLandIdx
	}
	return state
}

func blightSingleLand(state game_state.GameState, landIdx int) (game_state.GameState, bool) {
	if state.InvaderState.BlightPool == 0 {
		return state, false
	}

	land := state.BoardState.Lands[landIdx]

	shouldBlightCascade := false
	if land.NumBlight > 0 {
		shouldBlightCascade = true
	}
	land.NumBlight++
	land.NumPresence = utils.GetMaxInt(land.NumPresence-1, 0)
	state.BoardState.Lands[landIdx] = land
	state.InvaderState.BlightPool--

	if !state.InvaderState.IsBlightedIsland {
		state.InvaderState.IsBlightedIsland = true
		state.InvaderState.BlightPool = state.InvaderState.BlightedIslandCard.BlightPerPlayer
	}

	// Optimization: if there's no more Blight after allocating, don't Blight cascade (it's not possible)
	if state.InvaderState.BlightPool == 0 {
		shouldBlightCascade = false
	}

	return state, shouldBlightCascade
}

// To efficiently distribute damage, we hunt the most-damaged Dahan first (i.e. always using the
// least amount of damage to gain kills)
func efficientlyDamageDahan(damageTaken []int, dahanHealth int, damageToDistribute int) []int {
	var damageTakenCopy []int

	if len(damageTaken) == 0 {
		return damageTaken
	}

	damageTakenCopy = make([]int, len(damageTaken))
	copy(damageTakenCopy, damageTaken)

	// Therefore, we need to sort the Dahan from most-damaged to least-damaged
	// (and we use i and j as tiebreakers)
	dahanKillOrder := make([]int, len(damageTaken))
	for i := 0; i < len(damageTaken); i++ {
		dahanKillOrder[i] = i
	}
	sort.Slice(dahanKillOrder, func(i, j int) bool {
		iDamage := damageTaken[i]
		jDamage := damageTakenCopy[j]
		if iDamage == jDamage {
			return i < j
		}
		return damageTaken[i] > damageTaken[j]
	})

	// Now, start killing the weakest Dahan first
	remainingDamageToDistribute := damageToDistribute
	for _, dahanIdx := range dahanKillOrder {
		if remainingDamageToDistribute == 0 {
			break
		}

		damageToKill := dahanHealth - damageTakenCopy[dahanIdx]
		damageSpent := utils.GetMintInt(damageToKill, remainingDamageToDistribute)

		damageTakenCopy[dahanIdx] += damageSpent
		remainingDamageToDistribute -= damageSpent
	}

	// Finally, cull the dead Dahan
	var result []int
	for _, damage := range damageTakenCopy {
		if damage < dahanHealth {
			result = append(result, damage)
		}
	}

	return result
}
