package island

import (
	"fmt"
	"github.com/bobg/go-generics/v2/set"
	"github.com/imdario/mergo"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island"
	"github.com/mieubrisse/open-spirit-island/game/game_state/island/filter"
	"sort"
	"strconv"
	"strings"
)

// The base column config, on which the overrides will be applied
var baseColumnConfig = table.ColumnConfig{
	Align:        text.AlignCenter,
	VAlign:       text.VAlignMiddle,
	AlignHeader:  text.AlignCenter,
	VAlignHeader: text.VAlignMiddle,
}

// Overrides per column to apply on top of the base
var columnConfigOverrides = []table.ColumnConfig{
	{Name: "Land"},
	{Name: "Type"},
	{Name: "Adjacencies", Transformer: func(val interface{}) string {
		adjacentLandsSet := val.(set.Of[int])
		adjacentLandsIntList := adjacentLandsSet.Slice()
		sort.Ints(adjacentLandsIntList)
		adjacentLandsStrList := make([]string, len(adjacentLandsIntList))
		for j, adjacentIndx := range adjacentLandsIntList {
			adjacentLandsStrList[j] = strconv.Itoa(adjacentIndx)
		}
		return strings.Join(adjacentLandsStrList, ",")
	}},
	{Name: "🪔"},
	{Name: "Dahan"},
	{Name: "Invaders"},
	{Name: "Invader Damage", Transformer: renderIntegerSkip0},
	{Name: "Blight", Transformer: renderIntegerSkip0},
}

func RenderIslandState(state island.IslandBoardState) string {
	tableWriter := table.NewWriter()
	tableWriter.SetStyle(table.StyleLight)

	columnConfigs := make([]table.ColumnConfig, len(columnConfigOverrides))
	for idx, override := range columnConfigOverrides {
		overriden := baseColumnConfig
		mergo.Merge(&overriden, override)
		columnConfigs[idx] = overriden
	}

	tableWriter.SetColumnConfigs(columnConfigs)

	headerRow := make([]interface{}, len(columnConfigs))
	for idx, columnConfig := range columnConfigs {
		headerRow[idx] = columnConfig.Name
	}
	tableWriter.AppendHeader(headerRow)

	for i, land := range state.Lands {
		adjacentLandIdxs := state.FilterLands(filter.IslandFilter{
			SourceNumbers: set.New(i),
			MinRange:      1,
			MaxRange:      1,
		})

		dahanStrs := renderObjects("Dahan", land.DahanHealth, island.DahanBaseHealth)
		dahanCell := strings.Join(dahanStrs, "\n")

		invaderStrs := make([]string, 0, len(land.CityHealth)+len(land.TownHealth)+len(land.ExplorerHealth))
		invaderStrs = append(invaderStrs, renderObjects("City", land.CityHealth, island.CityBaseHealth)...)
		invaderStrs = append(invaderStrs, renderObjects("Town", land.TownHealth, island.TownBaseHealth)...)
		invaderStrs = append(invaderStrs, renderObjects("Explorer", land.ExplorerHealth, island.ExplorerBaseHealth)...)
		invaderCell := strings.Join(invaderStrs, "\n")

		invaderDamage := island.CityBaseDamage*len(land.CityHealth) + island.TownBaseDamage*len(land.TownHealth) + island.ExplorerBaseDamage*len(land.ExplorerHealth)

		row := []interface{}{
			i,
			land.LandType,
			adjacentLandIdxs,
			strings.Repeat("🪔", land.NumPresence),
			dahanCell,
			invaderCell,
			// TODO account for defense
			invaderDamage,
			land.NumBlight,
		}

		tableWriter.AppendRow(row)
		if i < len(state.Lands)-1 {
			tableWriter.AppendSeparator()
		}
	}

	return tableWriter.Render()
}

func renderIntegerSkip0(val interface{}) string {
	casted := val.(int)
	if casted == 0 {
		return ""
	}
	return strconv.Itoa(casted)
}

func renderObjects(objectTitle string, allObjectsHealth []int, baseObjectHealth int) []string {
	result := make([]string, len(allObjectsHealth))
	for i, health := range allObjectsHealth {
		result[i] = fmt.Sprintf("%s (%d/%d)", objectTitle, health, baseObjectHealth)
	}
	return result
}
