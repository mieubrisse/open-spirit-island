package main

import (
	"fmt"
	"github.com/mieubrisse/open-spirit-island/game_state"
	"github.com/mieubrisse/open-spirit-island/game_state/status"
)

func main() {
	gameState := game_state.NewTestGame()

	// TODO presence-placing
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
