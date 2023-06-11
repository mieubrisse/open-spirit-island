package input

import (
	"fmt"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/power"
	"strings"
)

func PlayCards(handCards []power.PowerCard, energyAvailable int, cardPlaysAvailable int) []int {
	fmt.Println("Choose cards to play:")
	for i, card := range handCards {
		firstLinePrefix := fmt.Sprintf(" %s) ", base26Encode(i))
		secondPlusLinePrefix := strings.Repeat(" ", len(firstLinePrefix))

		cardStr := card.String()
		optionLines := strings.Split(cardStr, "\n")
		for j, line := range optionLines {
			prefix := firstLinePrefix
			if j > 0 {
				prefix = secondPlusLinePrefix
			}
			optionLines[j] = prefix + line
		}
		optionStr := strings.Join(optionLines, "\n")
		fmt.Println(optionStr)
	}
	fmt.Println(fmt.Sprintf("You have %d⚡ and %d plays available", energyAvailable, cardPlaysAvailable))

	for {
		selectionIdxs, err := readAndValidateSelection(len(handCards), 0, cardPlaysAvailable)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		totalEnergyCost := 0
		for _, selectionIdx := range selectionIdxs {
			totalEnergyCost += handCards[selectionIdx].Cost
		}
		if totalEnergyCost > energyAvailable {
			fmt.Println(fmt.Sprintf("The selected cards cost %d⚡, but you only have %d⚡", totalEnergyCost, energyAvailable))
		}

		return selectionIdxs
	}
}
