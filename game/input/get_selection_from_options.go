package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

outerLoop:
	for {
		fmt.Print("Select: ")
		reader := bufio.NewReader(os.Stdin)
		// TODO handle isPrefix case?
		line, _, err := reader.ReadLine()
		if err != nil {
			panic("Got an error when reading user input: " + err.Error())
		}

		fields := strings.Fields(string(line))
		numFields := len(fields)
		if numFields < minSelectionsAllowed {
			fmt.Println(fmt.Sprintf("Require at least %d option(s) selected, but got %d", minSelectionsAllowed, numFields))
			continue
		}
		if numFields > maxSelectionsAllowed {
			fmt.Println(fmt.Sprintf("Require at max %d option(s) selected, but got %d", minSelectionsAllowed, numFields))
			continue
		}

		allSelectionIdxs := make([]int, numFields)
		for i, field := range fields {
			selectionIdx, err := base26Decode(field)
			if err != nil {
				fmt.Println("Selection '%s' is invalid", field)
				continue outerLoop
			}

			if selectionIdx < 0 || selectionIdx >= len(options) {
				fmt.Println(fmt.Sprintf(
					"Selection '%s' is not in range %s-%s",
					field,
					base26Encode(0),
					base26Encode(len(options)-1),
				))
				continue outerLoop
			}
			allSelectionIdxs[i] = selectionIdx
		}

		return allSelectionIdxs
	}
}
