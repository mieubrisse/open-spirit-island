package input

import "fmt"

const (
	alphabeticalEncodingChars string = "abcdefghijklmnopqrstuvwxyz"
)

var encodingCharOrdinals map[string]int

func init() {
	encodingCharOrdinals = make(map[string]int, len(alphabeticalEncodingChars))
	for i, char := range alphabeticalEncodingChars {
		encodingCharOrdinals[string(char)] = i
	}
}

// Converts a decimal string to a letter (base 26)
func base26Encode(number int) string {
	// Mayyyyybe we'll need to support > 26 choices, but I doubt it - can always support those in the future if we need
	return string(alphabeticalEncodingChars[number])
}

func base26Decode(str string) (int, error) {
	decimal, found := encodingCharOrdinals[str]
	if !found {
		return 0, fmt.Errorf("invalid string '%v'", str)
	}
	return decimal, nil
}
