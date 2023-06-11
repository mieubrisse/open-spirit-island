package game

import (
	"fmt"
	"github.com/bobg/go-generics/v2/set"
	"github.com/mieubrisse/open-spirit-island/game/game_state"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/filter"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/land_type"
	"github.com/mieubrisse/open-spirit-island/game/game_state/status"
	"github.com/mieubrisse/open-spirit-island/game/utils"
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


	// TODO first explore

	fmt.Println(gameState.String())

	for gameState.GetStatus() == status.Undecided {
		fmt.Scanln()
		fmt.Println("\n\n\n")
		gameState = gameState.RunGrowthPhase()
		gameState = gameState.RunInvaderPhase()
		fmt.Println(gameState.String())
	}
}

func RunGrowthPhase() GameState {
	// Can't change the game state beyond victory or defeat
	if state.GetStatus() != status.Undecided {
		return state
	}

	growthChoices :=

	// TODO growth choice

	// TODO elemental income

	// TODO play & pay power cards
}
