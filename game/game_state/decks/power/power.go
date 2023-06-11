package power

import "github.com/bobg/go-generics/v2/set"

type PowerCardID int
type PowerCard struct {
	Title string

	Elements set.Of[Element]
}
