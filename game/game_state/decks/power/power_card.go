package power

import "github.com/bobg/go-generics/v2/set"

type PowerCard struct {
	Title string

	Cost int

	Speed PowerCardSpeed

	Elements set.Of[Element]

	TransitionsID PowerCardTransitionsID
}
