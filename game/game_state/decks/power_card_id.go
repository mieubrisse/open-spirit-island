package decks

//go:generate go run github.com/dmarkham/enumer -type=PowerCardID
type PowerCardID int

const (
	// Vital Strength of Earth's Unique Powers
	AYearOfPerfectStillness PowerCardID = iota
	DrawOfTheFruitfulEarth
	GuardTheHealingLand
	RitualsOfDestruction
)
