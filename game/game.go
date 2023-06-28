package game

import (
	"fmt"
	"github.com/mieubrisse/open-spirit-island/game/game_state"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/power"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/power/default_power_cards"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/land_type"
	"github.com/mieubrisse/open-spirit-island/game/game_state/status"
	"github.com/mieubrisse/open-spirit-island/game/phases"
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
	gameState.PlayerState.Hand = []power.PowerCard{default_power_cards.DrawOfTheFruitfulEarth}

	// First explore
	printSection("Initial State")
	gameState = phases.RunInvaderPhase(gameState)
	fmt.Println(gameState.String())

	// This is the inner simulation loop; nearly every player action will result in coming back here
	// This is also what will allow us to do detailed undo/redo
	for gameState.GetStatus() == status.Undecided {
		// TODO checkpoint the game here

		switch gameState.Phase {
		case game_state.SpiritGrow:
			gameState = phases.RunSpiritGrowPhase(gameState)
		case game_state.SpiritGainTrackBenefits:
			gameState = phases.RunSpiritGainTrackBenefitsPhase(gameState)
		case game_state.SpiritPowerPlays:
			gameState = phases.RunSpiritPlayPowerPhase(gameState)
		case game_state.FastPower:
			// TODO implement
			gameState.Phase++
		case game_state.Invader:
			gameState = phases.RunInvaderPhase(gameState)
		case game_state.SlowPower:
			// TODO implement
			gameState.Phase++
		case game_state.TimePasses:

		default:
			panic(fmt.Errorf("Unrecognized game phase: %d", gameState.Phase))
		}
	}
}

func printSection(title string) {
	fmt.Println("\n\n")
	fmt.Println(fmt.Sprintf("===================================== %s ===============================================", title))
}

func printPhaseSection(phase string, turnNumber int) {
	title := fmt.Sprintf("%s Phase %d", phase, turnNumber)
	printSection(title)
}
