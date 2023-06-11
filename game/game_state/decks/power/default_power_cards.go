package power

import "github.com/bobg/go-generics/v2/set"

var DefaultPowerCards = map[PowerCardID]PowerCard{
	0: {
		Title:    "The Land Thrashes In Furious Pain",
		Elements: set.New(Moon, Fire, Earth),
	},
}
