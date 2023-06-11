package input

import (
	"fmt"
)

func GetSingleSelection(prompt string, options []string) int {
	return GetMultipleSelections(prompt, options, 1, 1)[0]
}

func GetMultipleSelections(prompt string, options []string, minSelectionsAllowed int, maxSelectionsAllowed int) []int {
	fmt.Println(prompt)
	for idx, option := range options {
		fmt.Println(fmt.Sprintf(" %s) %s", base26Encode(idx), option))
	}

	// TODO handle this error!

	for {
		selectionIdxs, err := readAndValidateSelection(len(options), minSelectionsAllowed, maxSelectionsAllowed)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		return selectionIdxs
	}
}
