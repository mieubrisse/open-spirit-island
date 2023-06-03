package island

import (
	"github.com/imdario/mergo"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
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
		adjacentLandIdxs := val.([]int)
		adjacentLandIdxStrs := make([]string, len(adjacentLandIdxs))
		for j, adjacentLandIdx := range adjacentLandIdxs {
			adjacentLandIdxStrs[j] = strconv.Itoa(adjacentLandIdx)
		}
		return strings.Join(adjacentLandIdxStrs, ",")
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
	// Same as the numbers on each board, with the ocean being 0
	Lands []LandState

	// Mapping of land_index -> bordering_land_indexes
	Adjacencies [][]int
}

// TODO nice error-handling
func (state IslandBoardState) AddPresence(landIdx int) IslandBoardState {
	state.Lands[landIdx].NumPresence++
	return state
}

// Gets the indexes of the adjacent lands, returned in sorted order
func (state IslandBoardState) GetAdjacentLands(landIdx int) []int {
	resultSet := map[int]bool{}
	for _, pair := range state.Adjacencies {
		if pair[0] == landIdx {
			resultSet[pair[1]] = true
		}
		if pair[1] == landIdx {
			resultSet[pair[0]] = true
		}
	}

	result := make([]int, 0, len(resultSet))
	for adjacentLandIdx := range resultSet {
		result = append(result, adjacentLandIdx)
	}

	sort.Ints(result)

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
		adjacentLandIdxs := state.GetAdjacentLands(i)

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
