package game_state

import (
	"fmt"
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game_state/invader_board"
	"github.com/mieubrisse/open-spirit-island/game_state/island"
	"github.com/mieubrisse/open-spirit-island/game_state/island/filter"
	"github.com/mieubrisse/open-spirit-island/game_state/island/land_type"
	"github.com/mieubrisse/open-spirit-island/game_state/player"
	"github.com/mieubrisse/open-spirit-island/game_state/status"
	"github.com/mieubrisse/open-spirit-island/input"
	"github.com/mieubrisse/open-spirit-island/utils"
	"sort"
	"strings"
)

type GameState struct {
	/*
		MajorPowersDeck    []power.Power
		MajorPowersDiscord []power.Power

		MinorPowersDeck    []power.Power
		MinorPowersDiscard []power.Power
	*/

	InvaderState invader_board.InvaderBoardState

	// TODO multiple players
	PlayerState player.PlayerState

	// TODO multiple boards
	BoardState island.IslandBoardState
}

func (state GameState) GetStatus() status.GameStatus {
	numExplorers := 0
	numTowns := 0
	numCities := 0
	numPresence := 0
	for _, land := range state.BoardState.Lands {
		numExplorers += len(land.ExplorerHealth)
		numTowns += len(land.TownHealth)
		numCities += len(land.CityHealth)
		numPresence += land.NumPresence
	}

	// Check if the user has won
	terrorLevel := state.InvaderState.GetTerrorLevel()
	victoryChecks := []bool{
		terrorLevel >= 4,
		terrorLevel == 3 && numCities == 0,
		terrorLevel == 2 && numCities == 0 && numTowns == 0,
		terrorLevel == 1 && numCities == 0 && numTowns == 0 && numExplorers == 0,
	}
	isFearVictory := false
	for _, check := range victoryChecks {
		isFearVictory = isFearVictory || check
	}
	if isFearVictory {
		return status.Victory
	}

	// Presence token loss
	if numPresence == 0 {
		return status.Defeat
	}

	// Blight loss
	if state.InvaderState.BlightPool == 0 {
		return status.Defeat
	}

	// Invader deck loss
	if len(state.InvaderState.RemainingInvaderDeck) == 0 {
		return status.Defeat
	}

	return status.Undecided
}

func (state GameState) String() string {
	lines := []string{
		state.InvaderState.String(),
		state.BoardState.String(),
		// TODO multiple players
		"PLAYER",
		state.PlayerState.String(),
	}

	currentStatus := state.GetStatus()

	header := []string{}
	if currentStatus != status.Undecided {
		header = []string{
			"!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!",
			"                                            " + strings.ToUpper(currentStatus.String()),
			"!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!",
		}
	}
	lines = append(header, lines...)

	return strings.Join(lines, "\n")
}

// TODO get the score

// Runs the invasion phase
func (state GameState) RunInvaderPhase() GameState {
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
				state = state.blightLandWithCascade(landIdx)
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

func (state GameState) blightLandWithCascade(landIdx int) GameState {
	var shouldBlightCascade bool
	state, shouldBlightCascade = state.blightSingleLand(landIdx)

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

		state, shouldBlightCascade = state.blightSingleLand(selectedLandIdx)

		sourceLandIdx = selectedLandIdx
	}
	return state
}

func (state GameState) blightSingleLand(landIdx int) (GameState, bool) {
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
