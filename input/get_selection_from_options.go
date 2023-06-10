package input

import (
	"bufio"
	"fmt"
	"os"
)

func GetSelectionFromOptions(prompt string, options []string) int {
	fmt.Println(prompt)
	for idx, option := range options {
		fmt.Println(fmt.Sprintf(" %s) %s", base26Encode(idx), option))
	}

	// TODO handle this error!

	for {
		fmt.Print("Select: ")
		reader := bufio.NewReader(os.Stdin)
		// TODO handle isPrefix case?
		line, _, err := reader.ReadLine()
		if err != nil {
			panic("Got an error when reading user input: " + err.Error())
		}

		selection, err := base26Decode(string(line))
		if err != nil {
			fmt.Println("Unrecognized selection string")
			continue
		}

		if selection < 0 || selection >= len(options) {
			fmt.Println(fmt.Sprintf(
				"Selection must be %s-%s",
				base26Encode(0),
				base26Encode(len(options)-1),
			))
			continue
		}

		return selection
	}
}
