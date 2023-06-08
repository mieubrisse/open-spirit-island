package blighted_island

type BlightedIslandCard struct {
	Title string

	BlightPerPlayer int
}

func (b BlightedIslandCard) String() string {
	return b.Title
}
