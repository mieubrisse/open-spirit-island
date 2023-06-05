package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetUserSelection(prompt string, options []string) int {
	fmt.Println(prompt)
	for idx, option := range options {
		fmt.Println(fmt.Sprintf(" %d) %s", idx, option))
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

		selection, err := strconv.Atoi(string(line))
		if err != nil {
			fmt.Println("Invalid selection: " + err.Error())
			continue
		}

		if selection < 0 || selection >= len(options) {
			fmt.Println(fmt.Sprintf("Selection must be in range [0,%d]", len(options)-1))
			continue
		}

		return selection
	}
}
