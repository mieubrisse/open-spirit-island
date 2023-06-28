package input

import (
	"bufio"
	"fmt"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/land_state"
	"os"
	"strconv"
	"strings"
)

// Allocates damage to the invaders
func DamageInvaders(
	actorDescription string,
	land land_state.LandState,
	damageToAllocate int,
) land_state.LandState {

	numCities := len(land.CityDamageTaken)
	numTowns := len(land.TownDamageTaken)
	numExplorers := len(land.ExplorerDamageTaken)
	if damageToAllocate == 0 ||
		(numCities == 0 && numTowns == 0 && numExplorers == 0) {

		return land
	}

	// Convenience so we don't require user input when there's enough damage to kill all the explorers
	fullClearDamageNeeded := 0
	for _, damageTaken := range land.CityDamageTaken {
		fullClearDamageNeeded += land.CityHealth - damageTaken
	}
	for _, damageTaken := range land.TownDamageTaken {
		fullClearDamageNeeded += land.TownHealth - damageTaken
	}
	for _, damageTaken := range land.ExplorerDamageTaken {
		fullClearDamageNeeded += land.ExplorerHealth - damageTaken
	}
	if damageToAllocate >= fullClearDamageNeeded {
		land.CityDamageTaken = []int{}
		land.TownDamageTaken = []int{}
		land.ExplorerDamageTaken = []int{}
	}

	// Convenience so we don't require user input when cleaning out lands with just 1HP explorers
	if land.ExplorerHealth == 1 && numExplorers > 0 && numTowns == 0 && numCities == 0 {
		explorersToRemain := numExplorers - damageToAllocate
		land.ExplorerDamageTaken = land.ExplorerDamageTaken[:explorersToRemain]
		return land
	}

	fmt.Println(fmt.Sprintf("%s will deal %d damage to the invaders:", actorDescription, damageToAllocate))

	optionIdx := 0
	for _, damageTaken := range land.CityDamageTaken {
		fmt.Println(fmt.Sprintf(
			" %s) City (%d/%d)",
			base26Encode(optionIdx),
			land.CityHealth-damageTaken,
			land.CityHealth,
		))
		optionIdx++
	}
	for _, damageTaken := range land.TownDamageTaken {
		fmt.Println(fmt.Sprintf(
			" %s) Town (%d/%d)",
			base26Encode(optionIdx),
			land.TownHealth-damageTaken,
			land.TownHealth,
		))
		optionIdx++
	}
	for _, damageTaken := range land.ExplorerDamageTaken {
		fmt.Println(fmt.Sprintf(
			" %s) Explorer (%d/%d)",
			base26Encode(optionIdx),
			land.ExplorerHealth-damageTaken,
			land.ExplorerHealth,
		))
		optionIdx++
	}
	numOptions := optionIdx

	fmt.Println("Allocate damage to the invaders in space-separated invader=damage format (e.g. 'a=4 b=2 c=1')")
	var allocation map[int]int
allocationLoop:
	for {
		fmt.Print("Allocate: ")
		reader := bufio.NewReader(os.Stdin)
		// TODO handle isPrefix case?
		line, _, err := reader.ReadLine()
		if err != nil {
			panic("Got an error when reading user input: " + err.Error())
		}

		fields := strings.Fields(string(line))
		candidateAllocation := make(map[int]int, len(fields)) // invader_idx -> damage
		sumAllocatedDamage := 0
		for _, field := range fields {
			fragments := strings.Split(field, "=")
			if len(fragments) != 2 {
				fmt.Println(fmt.Sprintf("Allocation '%s' isn't in invader=damage format", field))
				continue allocationLoop
			}
			optionStr := fragments[0]

			invaderIdx, err := base26Decode(optionStr)
			if err != nil {
				fmt.Println(fmt.Sprintf("Allocation '%s' doesn't specify the alphabetical choice of an invader to damage", field))
				continue allocationLoop
			}
			if invaderIdx < 0 || invaderIdx >= numOptions {
				fmt.Println(fmt.Sprintf(
					"Allocation '%s' specifies an out-of-range invader selection (must be %s-%s)",
					field,
					base26Encode(0),
					base26Encode(numOptions-1),
				))
				continue allocationLoop
			}

			damage, err := strconv.Atoi(fragments[1])
			if err != nil {
				fmt.Println(fmt.Sprintf("Allocation '%s' doesn't specify numerical damage", field))
				continue allocationLoop
			}

			if damage < 0 {
				fmt.Println(fmt.Sprintf("Allocation '%s' attempts to allocate negative damage", field))
				continue allocationLoop
			}

			if invaderIdx < numCities {
				// First C entries are cities

				cityHealthRemaining := land.CityHealth - land.CityDamageTaken[invaderIdx]
				if damage > cityHealthRemaining {
					fmt.Println(fmt.Sprintf(
						"Allocation '%s' attempts to damage City '%s' for %d, but it only has %d HP",
						field,
						optionStr,
						damage,
						cityHealthRemaining,
					))
					continue allocationLoop
				}
			} else if invaderIdx < numCities+numTowns {
				// Next T entries are towns

				townArrIdx := invaderIdx - numCities
				townHealthRemaining := land.TownHealth - land.TownDamageTaken[townArrIdx]
				if damage > townHealthRemaining {
					fmt.Println(fmt.Sprintf(
						"Allocation '%s' attempts to damage Town '%s' for %d, but it only has %d HP",
						field,
						optionStr,
						damage,
						townHealthRemaining,
					))
					continue allocationLoop
				}
			} else if invaderIdx < numCities+numTowns+numExplorers {
				// Last E entries are explorers

				explorersArrIdx := invaderIdx - numCities - numTowns
				explorerHealthRemaining := land.TownHealth - land.TownDamageTaken[explorersArrIdx]
				if damage > explorerHealthRemaining {
					fmt.Println(fmt.Sprintf(
						"Allocation '%s' attempts to damage Explorer '%s' for %d, but it only has %d HP",
						field,
						optionStr,
						damage,
						explorerHealthRemaining,
					))
					continue allocationLoop
				}
			} else {
				// Should never happen
				panic(fmt.Sprintf("Attempted to validate invalid invader idx %d; this should not have passed validation", invaderIdx))
			}

			candidateAllocation[invaderIdx] = damage
			sumAllocatedDamage += damage
		}

		if sumAllocatedDamage != damageToAllocate {
			fmt.Println(fmt.Sprintf("Need to allocate %d damage, but %d was allocated", damageToAllocate, sumAllocatedDamage))
			continue allocationLoop
		}

		allocation = candidateAllocation
		break
	}

	var cityDamageTakenCopy, townDamageTakenCopy, explorerDamageTakenCopy []int
	copy(cityDamageTakenCopy, land.CityDamageTaken)
	copy(townDamageTakenCopy, land.TownDamageTaken)
	copy(explorerDamageTakenCopy, land.ExplorerDamageTaken)

	for invaderIdx, damage := range allocation {
		// User selected a city
		if invaderIdx < numCities {
			cityDamageTakenCopy[invaderIdx] += damage
		} else if invaderIdx < numCities+numTowns {
			townDamageTakenCopy[invaderIdx-numCities] += damage
		} else if invaderIdx < numCities+numTowns+numExplorers {
			explorerDamageTakenCopy[invaderIdx-numCities-numTowns] += damage
		} else {
			// Should never happen
			panic(fmt.Sprintf("Attempted to apply damage to unknown invader idx %d; this should not have passed validation", invaderIdx))
		}
	}

	var resultCityDamageTaken, resultTownDamageTaken, resultExplorerDamageTaken []int
	for _, damageTaken := range cityDamageTakenCopy {
		if damageTaken < land.CityHealth {
			resultCityDamageTaken = append(resultCityDamageTaken, damageTaken)
		}
	}
	for _, damageTaken := range townDamageTakenCopy {
		if damageTaken < land.TownHealth {
			resultTownDamageTaken = append(resultTownDamageTaken, damageTaken)
		}
	}
	for _, damageTaken := range explorerDamageTakenCopy {
		if damageTaken < land.ExplorerHealth {
			resultExplorerDamageTaken = append(resultExplorerDamageTaken, damageTaken)
		}
	}

	land.CityDamageTaken = resultCityDamageTaken
	land.TownDamageTaken = resultTownDamageTaken
	land.ExplorerDamageTaken = resultExplorerDamageTaken

	return land
}
