package power

type PowerCardSpeed int

const (
	Fast PowerCardSpeed = iota
	Slow
)

var PowerCardSpeedSymbols = map[PowerCardSpeed]string{
	Fast: "🦅",
	Slow: "🐢",
}
