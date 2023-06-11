package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readAndValidateSelection(numChoices, minNumChoices, maxNumChoices int) ([]int, error) {
	fmt.Print("Select: ")
	reader := bufio.NewReader(os.Stdin)
	// TODO handle isPrefix case?
	line, _, err := reader.ReadLine()
	if err != nil {
		panic("Got an error when reading user input: " + err.Error())
	}

	fields := strings.Fields(string(line))
	numFields := len(fields)
	if numFields < minNumChoices {
		return nil, fmt.Errorf("Require at least %d option(s) selected, but got %d", minNumChoices, numFields)
	}
	if numFields > maxNumChoices {
		return nil, fmt.Errorf("Require at max %d option(s) selected, but got %d", maxNumChoices, numFields)
	}

	allSelectionIdxs := make([]int, numFields)
	for i, field := range fields {
		selectionIdx, err := base26Decode(field)
		if err != nil {
			return nil, fmt.Errorf("Selection '%s' is invalid", field)
		}

		if selectionIdx < 0 || selectionIdx >= numChoices {
			return nil, fmt.Errorf(
				"Selection '%s' is not in range %s-%s",
				field,
				base26Encode(0),
				base26Encode(numChoices-1),
			)
		}
		allSelectionIdxs[i] = selectionIdx
	}

	return allSelectionIdxs, nil
}
