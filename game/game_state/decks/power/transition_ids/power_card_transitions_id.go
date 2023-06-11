package transition_ids

// Because the GameState contains Power cards, but Power cards need to produce transitions which operate on the GameState,
// we end up with a circular dependency. To break this, we have power cards indirectly reference their effects
type PowerCardTransitionsID int

const (
	// Indicates no transitions
	None PowerCardTransitionsID = iota

	// Vital Strength of Earth unique powers
	DrawOfTheFruitfulEarth
	// TODO rest of unique powers
)
