package status

//go:generate go run github.com/dmarkham/enumer -type=GameStatus
type GameStatus int

const (
	Undecided GameStatus = iota
	Victory
	Defeat
)
