package island

import (
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
	{Name: "Dahan", Transformer: renderIntegerSkip0},
	{Name: "Explorers", Transformer: renderIntegerSkip0},
	{Name: "Towns", Transformer: renderIntegerSkip0},
	{Name: "Cities", Transformer: renderIntegerSkip0},
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

	rows := make([]table.Row, len(state.Lands))
	for i, land := range state.Lands {
		adjacentLandIdxs := state.FilterLands(filter.IslandFilter{
			SourceNumbers: set.New(i),
			MinRange:      1,
			MaxRange:      1,
		})

		row := []interface{}{
			i,
			land.LandType,
			adjacentLandIdxs,
			land.NumPresence,
			land.NumDahan,
			land.NumExplorers,
			land.NumTowns,
			land.NumCities,
			land.NumBlight,
		}

		rows = append(rows, row)
	}
	tableWriter.AppendRows(rows)

	return tableWriter.Render()
}

func renderIntegerSkip0(val interface{}) string {
	casted := val.(int)
	if casted == 0 {
		return ""
	}
	return strconv.Itoa(casted)
}
