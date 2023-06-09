package island

import (
	"fmt"
	"github.com/bobg/go-generics/v2/set"
	"github.com/imdario/mergo"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/mieubrisse/open-spirit-island/game_state/island/filter"
	"github.com/mieubrisse/open-spirit-island/game_state/island/land_state"
	"github.com/yourbasic/graph"
	"sort"
	"strconv"
	"strings"
)

// TODO one day make these customizable?
const (
	DahanBaseHealth = 2
	DahanBaseDamage = 2

	CityBaseHealth = 3
	CityBaseDamage = 3

	TownBaseHealth = 2
	TownBaseDamage = 2

	ExplorerBaseHealth = 1
	ExplorerBaseDamage = 1
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
	{Name: "Presence", Transformer: renderIntegerSkip0},
	{Name: "Dahan"},
	{Name: "Invaders"},
	{Name: "Invader Damage", Transformer: renderIntegerSkip0},
	{Name: "Blight", Transformer: renderIntegerSkip0},
}

// TODO separate boards
type IslandBoardState struct {
	Graph *graph.Immutable

	Lands []land_state.LandState

	/*
		// Mapping of land_index -> bordering_land_indexes
		Adjacencies [][]int

	*/
}

// TODO nice error-handling
func (state IslandBoardState) AddPresence(landIdx int) IslandBoardState {
	state.Lands[landIdx].NumPresence++
	return state
}

// Swiss army knife for selecting lands on the island
func (state IslandBoardState) FilterLands(filter filter.IslandFilter) set.Of[int] {
	sourcesIdx := set.New[int]()
	for idx, land := range state.Lands {
		if filter.SourceNumbers != nil {
			if !filter.SourceNumbers.Has(idx) {
				continue
			}
		}

		if filter.SourceFilter.Match(land) {
			sourcesIdx.Add(idx)
		}
	}

	result := set.New[int]()
	for sourceIdx := range sourcesIdx {
		_, distancesInt64 := graph.ShortestPaths(state.Graph, sourceIdx)
		distances := make([]int, len(distancesInt64))
		for i, value := range distancesInt64 {
			distances[i] = int(value)
		}

		for targetLandIdx, distance := range distances {
			targetLand := state.Lands[targetLandIdx]

			if distance < filter.MinRange {
				continue
			}

			if distance > filter.MaxRange {
				continue
			}

			if !filter.TargetFilter.Match(targetLand) {
				continue
			}

			result.Add(targetLandIdx)
		}
	}

	return result
}

func (state IslandBoardState) String() string {
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

		dahanStrs := renderObjects("Dahan", land.DahanHealth, DahanBaseHealth)
		dahanCell := strings.Join(dahanStrs, "\n")

		invaderStrs := make([]string, 0, len(land.CityHealth)+len(land.TownHealth)+len(land.ExplorerHealth))
		invaderStrs = append(invaderStrs, renderObjects("City", land.CityHealth, CityBaseHealth)...)
		invaderStrs = append(invaderStrs, renderObjects("Town", land.TownHealth, TownBaseHealth)...)
		invaderStrs = append(invaderStrs, renderObjects("Explorer", land.ExplorerHealth, ExplorerBaseHealth)...)
		invaderCell := strings.Join(invaderStrs, "\n")

		invaderDamage := CityBaseDamage*len(land.CityHealth) + TownBaseDamage*len(land.TownHealth) + ExplorerBaseDamage*len(land.ExplorerHealth)

		row := []interface{}{
			i,
			land.LandType,
			adjacentLandIdxs,
			land.NumPresence,
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
