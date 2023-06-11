package game

import (
	"fmt"
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game/game_state"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/filter"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/land_type"
	"github.com/mieubrisse/open-spirit-island/game/game_state/status"
	input "github.com/mieubrisse/open-spirit-island/game/input"
	"github.com/mieubrisse/open-spirit-island/game/transitions"
	"github.com/mieubrisse/open-spirit-island/game/utils"
	"sort"
	"strings"
)

func RunGameLoop() {
	gameState := game_state.NewTestGame()

	// TODO allow for other spirits; ths is Vital Strength of Earth
	highestNumberedJungle := -1
	highestNumberedMountain := -1
	for idx, land := range gameState.BoardState.Lands {
		if land.LandType == land_type.Jungle {
			highestNumberedJungle = utils.GetMaxInt(idx, highestNumberedJungle)
		}
		if land.LandType == land_type.Mountain {
			highestNumberedMountain = utils.GetMaxInt(idx, highestNumberedMountain)
		}
	}
	gameState.BoardState.Lands[highestNumberedMountain].NumPresence += 2
	gameState.BoardState.Lands[highestNumberedJungle].NumPresence += 1

	// First explore
	printSection("Initial State")
	gameState = RunInvaderPhase(gameState)
	fmt.Println(gameState.String())

	turnNumber := 0
	for gameState.GetStatus() == status.Undecided {
		printPhaseSection("Growth", turnNumber)
		gameState = RunGrowthPhase(gameState)
		fmt.Println(gameState.String())

		// TODO fast power phase

		printPhaseSection("Invader", turnNumber)
		gameState = RunInvaderPhase(gameState)
		fmt.Println(gameState.String())

		// TODO slow power phase

		printPhaseSection("Time Passes", turnNumber)
		// TODO time passes

		turnNumber++
	}
}

func RunGrowthPhase(state game_state.GameState) game_state.GameState {
	// Can't change the game state beyond victory or defeat
	if state.GetStatus() != status.Undecided {
		return state
	}

	// TODO other spirit choices
	growthChoices := [][]transitions.GameStateTransition{
		{
			transitions.ReclaimAllCardsTransition,
			transitions.NewNormalAddPresenceTransition(2),
		},
		{
			// TODO gain a power card
			transitions.NewNormalAddPresenceTransition(0),
		},
		{
			transitions.NewNormalAddPresenceTransition(1),
			transitions.NewGainEnergyTransition(2),
		},
	}

	growthChoicesStrs := make([]string, len(growthChoices))
	for i, choiceSubcomponents := range growthChoices {
		choiceComponentStrs := make([]string, len(choiceSubcomponents))
		for j, subcomponent := range choiceSubcomponents {
			choiceComponentStrs[j] = subcomponent.String()
		}
		growthChoicesStrs[i] = strings.Join(choiceComponentStrs, "   |   ")
	}

	selectionIdx := input.GetSelectionFromOptions("Select a Growth option:", growthChoicesStrs)

	growthSelection := growthChoices[selectionIdx]

	newGameState := state
	for _, transition := range growthSelection {
		newGameState = transition.TransitionFunction(newGameState)
	}

	// TODO elemental income

	// TODO play & pay power cards

	return newGameState
}

// Runs the invasion phase
func RunInvaderPhase(state game_state.GameState) game_state.GameState {
	// Can't change the game state beyond victory or defeat
	if state.GetStatus() != status.Undecided {
		return state
	}

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

			invaderDamage := island.ExplorerBaseDamage*len(ravageLand.ExplorerHealth) + island.TownBaseDamage*len(ravageLand.TownHealth) + island.CityBaseDamage*len(ravageLand.CityHealth)

			// Dahan are attacked
			// From the rules, Dahan must be destroyed as efficiently as possible
			ravageLand.DahanHealth = efficientlyDamageDahan(ravageLand.DahanHealth, invaderDamage)

			// Dahan strike back
			dahanDamage := island.DahanBaseDamage * len(ravageLand.DahanHealth)
			ravageLand.CityHealth, ravageLand.TownHealth, ravageLand.ExplorerHealth = input.DamageInvaders(
				"Dahan",
				ravageLand.CityHealth,
				ravageLand.TownHealth,
				ravageLand.ExplorerHealth,
				dahanDamage,
			)

			// TODO defend

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

			if len(buildLand.TownHealth) > len(buildLand.CityHealth) {
				buildLand.CityHealth = append(buildLand.CityHealth, island.CityBaseHealth)
			} else {
				buildLand.TownHealth = append(buildLand.TownHealth, island.TownBaseHealth)
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
		exploreLand := state.BoardState.Lands[landIdx]
		exploreLand.ExplorerHealth = append(exploreLand.ExplorerHealth, island.ExplorerBaseHealth)
		state.BoardState.Lands[landIdx] = exploreLand
	}

	// Advance invader slots
	state.InvaderState = state.InvaderState.AdvanceInvaderCards()

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

		selection := input.GetSelectionFromOptions(
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

func efficientlyDamageDahan(currentDahanHp []int, damageToDistribute int) []int {
	var workingCopy []int

	if len(currentDahanHp) == 0 {
		return workingCopy
	}

	workingCopy = make([]int, len(currentDahanHp))
	copy(workingCopy, currentDahanHp)

	// To efficiently distribute damage, we hunt the weakest Dahan first (always using the
	// least amount of damage to gain kills)
	dahanHealthToTarget := 1
	for damageToDistribute > 0 {
		for i, dahanHp := range workingCopy {
			if dahanHp == dahanHealthToTarget {
				damageConsumed := utils.GetMintInt(dahanHp, damageToDistribute)
				workingCopy[i] -= damageConsumed
				damageToDistribute -= damageConsumed
			}
		}
		dahanHealthToTarget++ // Now hunt the next-strongest Dahan
	}

	var result []int
	for _, hp := range workingCopy {
		if hp > 0 {
			result = append(result, hp)
		}
	}

	return result
}

func printSection(title string) {
	fmt.Println("\n\n")
	fmt.Println(fmt.Sprintf("===================================== %s ===============================================", title))
}

func printPhaseSection(phase string, turnNumber int) {
	title := fmt.Sprintf("%s Phase %d", phase, turnNumber)
	printSection(title)
}
