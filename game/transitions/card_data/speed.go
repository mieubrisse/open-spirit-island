package card_data

type PowerCardSpeed int

const (
	Fast PowerCardSpeed = iota
	Slow
)

var PowerCardSpeedSymbols = map[PowerCardSpeed]string{
	Fast: "🦅",
	Slow: "🐢",
}
