package power

//go:generate go run github.com/dmarkham/enumer -type=Element
type Element int

const (
	Sun Element = iota
	Moon
	Fire
	Air
	Water
	Earth
	Plant
	Animal
)

var ElementSymbols = map[Element]string{
	Sun:    "☀️ ",
	Moon:   "🌒 ",
	Fire:   "🔥",
	Air:    "🪶 ",
	Water:  "💧",
	Earth:  "⛰️ ",
	Plant:  "🌿",
	Animal: "🦞",
}
