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
		fmt.Sprintf(" %d) %s", idx, option)
	}

	// TODO handle this error!

	for {
		fmt.Print("Select: ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			panic("Got an error when reading user input: " + err.Error())
		}

		selection, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Invalid selection")
			continue
		}

		if selection < 0 || selection >= len(options) {
			fmt.Println(fmt.Sprintf("Selection must be in range [0,%d]", len(options)-1))
			continue
		}

		return selection
	}
}
