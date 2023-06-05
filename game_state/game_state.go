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
		numExplorers += land.NumExplorers
		numTowns += land.NumTowns
		numCities += land.NumCities
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
	// TODO

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

			invaderDamage := ravageLand.NumExplorers + 2*ravageLand.NumTowns + 3*ravageLand.NumCities

			// Dahan are attacked
			dahanToRemove := invaderDamage / 2
			ravageLand.NumDahan = utils.GetMaxInt(ravageLand.NumDahan-dahanToRemove, 0)
			state.BoardState.Lands[landIdx] = ravageLand

			// TODO Dahan strike back

			// TODO defend

			if invaderDamage >= 2 {

				mustSpreadBlight := false
				if ravageLand.NumBlight > 0 {
					mustSpreadBlight = true
				}
				ravageLand.NumBlight++
				ravageLand.NumPresence = utils.GetMaxInt(ravageLand.NumPresence-1, 0)
				state.BoardState.Lands[landIdx] = ravageLand

				sourceLandIdx := landIdx
				for mustSpreadBlight {
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
						optionStrs[i] = fmt.Sprintf("%s #%d (%d Blight)", adjacentLand.LandType, adjacentIdx, adjacentLand.NumBlight)
					}

					selectedLandIdx := input.GetUserSelection(
						fmt.Sprintf("%s #%d is suffering a Blight cascade; select an adjacent land to spread Blight to:", sourceLand.LandType, sourceLandIdx),
						optionStrs,
					)

					selectedLand := state.BoardState.Lands[selectedLandIdx]

					if selectedLand.NumBlight == 0 {
						mustSpreadBlight = false
					}

					sourceLand.NumBlight++
					sourceLand.NumPresence = utils.GetMaxInt(sourceLand.NumPresence-1, 0)
					state.BoardState.Lands[selectedLandIdx] = sourceLand
					sourceLandIdx = selectedLandIdx
				}

			}

			// TODO presence token destroy mechanics
		}
	}

	// Build
	buildSlot := state.InvaderState.BuildSlot
	if buildSlot.IsCardPresent {
		desiredBuildIdxs := state.BoardState.FilterLands(buildSlot.MaybeCard.TargetedLandSelector)
		actualBuildIdxs := set.Intersect(invaderLandIdxs, desiredBuildIdxs)

		for landIdx := range actualBuildIdxs {
			buildLand := state.BoardState.Lands[landIdx]
			// TODO build skips

			if buildLand.NumTowns > buildLand.NumCities {
				buildLand.NumCities++
			} else {
				buildLand.NumTowns++
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
		state.BoardState.Lands[landIdx].NumExplorers++
	}

	// Advance invader slots
	state.InvaderState = state.InvaderState.AdvanceInvaderCards()

	return state
}
