package power

// Because the GameState contains Power cards, but Power cards need to produce transitions which operate on the GameState,
// we end up with a circular dependency. To break this, we have power cards indirectly reference their effects
type PowerCardTransitionsID int

const (
	// Vital Strength of Earth unique powers
	DrawOfTheFruitfulEarth PowerCardTransitionsID = iota
	// TODO rest of unique powers
)
