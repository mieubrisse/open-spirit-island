package fear

type FearCard struct {
	// Unique ID representing this particular fear card
	id int

	// Flavor text title
	title string

	/*
		level1Action action.Action

		level2Action action.Action

		level3Action action.Action

	*/
}

// TODO replace with real fear card
func NewDummyFearCard() FearCard {
	return FearCard{
		id:    0,
		title: "DOING NOTHING",
	}
}
