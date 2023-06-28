package input

import (
	"fmt"
	"github.com/bobg/go-generics/v2/set"
)

func GetSingleSelection(prompt string, options []string) int {
	for selectionIdx := range GetMultipleSelections(prompt, options, 1, 1) {
		return selectionIdx
	}
	panic("Should never get here - we should always have exactly one selection")
}

func GetMultipleSelections(prompt string, options []string, minSelectionsAllowed int, maxSelectionsAllowed int) set.Of[int] {
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
