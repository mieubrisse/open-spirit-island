package game_state

import (
	"github.com/mieubrisse/open-spirit-island/game_state/invader_board"
	"github.com/mieubrisse/open-spirit-island/game_state/island"
	"github.com/mieubrisse/open-spirit-island/game_state/player"
	"github.com/mieubrisse/open-spirit-island/game_state/status"
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

func (state GameState) Advance() GameState {
	// Can't advance beyond victory or defeat
	if state.GetStatus() != status.Undecided {
		return state
	}

	// Ravage
	ravageSlot := state.InvaderState.RavageSlot
	if ravageSlot.IsCardPresent {
		targetedLandIdxs := ravageSlot.MaybeCard.TargetedLandSelector(state.BoardState)

		for _, targetedIdx := range targetedLandIdxs {
			targetedLand := state.BoardState.Lands[targetedIdx]

			// TODO ravage skips

			// TODO battle mechanics

			damageDealt := targetedLand.NumCities*3 + targetedLand.NumTowns*2 + targetedLand.NumExplorers
			if damageDealt >= 2 {
				// TODO defend

				targetedLand.NumBlight++
				state.BoardState.Lands[targetedIdx] = targetedLand

				// TODO blight cascade
			}

			// TODO presence token destroy mechanics
		}
	}

	// Build
	buildSlot := state.InvaderState.BuildSlot
	if buildSlot.IsCardPresent {
		targetedLandIdxs := buildSlot.MaybeCard.TargetedLandSelector(state.BoardState)

		for _, targetedIdx := range targetedLandIdxs {
			targetedLand := state.BoardState.Lands[targetedIdx]
			// TODO build skips

			areInvadersPresent := targetedLand.NumCities > 0 || targetedLand.NumTowns > 0 || targetedLand.NumExplorers > 0
			if !areInvadersPresent {
				continue
			}

			if targetedLand.NumTowns > targetedLand.NumCities {
				targetedLand.NumCities++
			} else {
				targetedLand.NumTowns++
			}
			state.BoardState.Lands[targetedIdx] = targetedLand
		}
	}

	// Explore
	explorerCard := state.InvaderState.RemainingInvaderDeck[0]
	maybeLandsToExplore := explorerCard.TargetedLandSelector(state.BoardState)
	for _, maybeLandToExploreIdx := range maybeLandsToExplore {
		// TODO explore skip

		maybeLandToExplore := state.BoardState.Lands[maybeLandToExploreIdx]
		shouldExplorePredicates := []bool{
			maybeLandToExplore.NumTowns > 0,
			maybeLandToExplore.NumCities > 0,
		}

		// Drop lands not adjacent to the coast or town/city
		adjacentLandIdxs := state.BoardState.GetAdjacentLands(maybeLandToExploreIdx)
		for _, adjacentIdx := range adjacentLandIdxs {
			adjacentLand := state.BoardState.Lands[adjacentIdx]

			shouldExplorePredicates = append(
				shouldExplorePredicates,
				adjacentLand.LandType == island.Ocean,
				adjacentLand.NumCities > 0,
				adjacentLand.NumTowns > 0,
			)
		}

		for _, predicate := range shouldExplorePredicates {
			if predicate {
				state.BoardState.Lands[maybeLandToExploreIdx].NumExplorers++
				break
			}
		}
	}

	// Advance invader slots
	state.InvaderState = state.InvaderState.AdvanceInvaderCards()

	return state
}

func (state GameState) String() string {
	lines := []string{
		"========================================== INVADERS ==========================================",
		state.InvaderState.String(),
		"=========================================== BOARD ============================================",
		state.BoardState.String(),
		// TODO multiple players
		"=========================================== PLAYER ===========================================",
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
