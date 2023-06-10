package input

import (
	"bufio"
	"fmt"
	"github.com/mieubrisse/open-spirit-island/game_state/island"
	"os"
	"strconv"
	"strings"
)

// Allocates damage to the invaders
func DamageInvaders(
	actorDescription string,
	currentCityHp []int,
	currentTownHp []int,
	currentExplorerHp []int,
	damageToAllocate int,
) ([]int, []int, []int) {
	numCities := len(currentCityHp)
	numTowns := len(currentTownHp)
	numExplorers := len(currentExplorerHp)
	if damageToAllocate == 0 || (numCities == 0 && numTowns == 0 && numExplorers == 0) {
		return currentCityHp, currentTownHp, currentExplorerHp
	}

	totalInvaderHealth := island.CityBaseHealth*numCities + island.TownBaseHealth*numTowns + island.ExplorerBaseHealth*numExplorers
	if damageToAllocate > totalInvaderHealth {
		return []int(nil), []int(nil), []int(nil)
	}

	fmt.Println(fmt.Sprintf("%s will deal %d damage to the invaders:", actorDescription, damageToAllocate))

	optionIdx := 0
	for _, cityHp := range currentCityHp {
		fmt.Println(fmt.Sprintf(
			" %s) City (%d/%d)",
			base26Encode(optionIdx),
			cityHp,
			island.CityBaseHealth,
		))
		optionIdx++
	}
	for _, townHp := range currentTownHp {
		fmt.Println(fmt.Sprintf(
			" %s) Town (%d/%d)",
			base26Encode(optionIdx),
			townHp,
			island.TownBaseHealth,
		))
		optionIdx++
	}
	for _, explorerHp := range currentExplorerHp {
		fmt.Println(fmt.Sprintf(
			" %s) Explorer (%d/%d)",
			base26Encode(optionIdx),
			explorerHp,
			island.ExplorerBaseHealth,
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
				if damage > currentCityHp[invaderIdx] {
					fmt.Println(fmt.Sprintf(
						"Allocation '%s' attempts to damage City '%s' for %d, but it only has %d HP",
						field,
						optionStr,
						damage,
						currentCityHp[invaderIdx],
					))
					continue allocationLoop
				}
			} else if invaderIdx < numCities+numTowns {
				townArrIdx := invaderIdx - numCities
				if damage > currentTownHp[townArrIdx] {
					fmt.Println(fmt.Sprintf(
						"Allocation '%s' attempts to damage Town '%s' for %d, but it only has %d HP",
						field,
						optionStr,
						damage,
						currentTownHp[townArrIdx],
					))
					continue allocationLoop
				}
			} else if invaderIdx < numCities+numTowns+numExplorers {
				explorersArrIdx := invaderIdx - numCities - numTowns
				if damage > currentExplorerHp[explorersArrIdx] {
					fmt.Println(fmt.Sprintf(
						"Allocation '%s' attempts to damage Explorer '%s' for %d, but it only has %d HP",
						field,
						optionStr,
						damage,
						currentTownHp[explorersArrIdx],
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

	var workingCopyCityHp, workingCopyTownHp, workingCopyExplorerHp []int
	copy(workingCopyCityHp, currentCityHp)
	copy(workingCopyTownHp, currentTownHp)
	copy(workingCopyExplorerHp, currentExplorerHp)

	for invaderIdx, damage := range allocation {
		// User selected a city
		if invaderIdx < numCities {
			currentCityHp[invaderIdx] -= damage
		} else if invaderIdx < numCities+numTowns {
			currentTownHp[invaderIdx-numCities] -= damage
		} else if invaderIdx < numCities+numTowns+numExplorers {
			currentExplorerHp[invaderIdx-numCities-numTowns] -= damage
		} else {
			// Should never happen
			panic(fmt.Sprintf("Attempted to apply damage to unknown invader idx %d; this should not have passed validation", invaderIdx))
		}
	}

	var resultCityHp, resultTownHp, resultExplorerHp []int
	for _, cityHp := range currentCityHp {
		if cityHp > 0 {
			resultCityHp = append(resultCityHp, cityHp)
		}
	}
	for _, townHp := range currentTownHp {
		if townHp > 0 {
			resultTownHp = append(resultTownHp, townHp)
		}
	}
	for _, explorerHp := range currentExplorerHp {
		if explorerHp > 0 {
			resultExplorerHp = append(resultExplorerHp, explorerHp)
		}
	}

	return resultCityHp, resultTownHp, resultExplorerHp
}
