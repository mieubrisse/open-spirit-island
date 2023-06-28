package static_assets

import "github.com/mieubrisse/open-spirit-island/game/game_state/player"

var ElementSymbols = map[player.Element]string{
	player.Sun:    "☀️ ",
	player.Moon:   "🌒 ",
	player.Fire:   "🔥",
	player.Air:    "🪶 ",
	player.Water:  "💧",
	player.Earth:  "⛰️ ",
	player.Plant:  "🌿",
	player.Animal: "🦞",
}
