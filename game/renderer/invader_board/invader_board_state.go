package invader_board

import (
	"fmt"
	"github.com/imdario/mergo"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/mieubrisse/open-spirit-island/game/game_state/invader_board"
	"math"
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
	{Name: ""},
	{Name: "Fear Cards"},
	{Name: "☠️"},
	{Name: "Terror Level"},
}

func RenderInvaderBoardState(state invader_board.InvaderBoardState) string {
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

	tableWriter.AppendRows([]table.Row{
		{
			"UNEARNED",
			len(state.UnearnedFearCards),
			state.UnearnedFear,
			fmt.Sprintf("Cards Till Next: %d", countFearCardsTillNextTerrorLevel(state)),
		},
		{
			"EARNED",
			len(state.EarnedFearCards),
			state.EarnedFear,
			fmt.Sprintf("Current: %d", state.GetTerrorLevel()),
		},
	})

	/*
		lines := []string{
			fmt.Sprintf("Unearned Fear Cards: %d             Unearned Fear: %d       Terror Level: %d", len(state.UnearnedFearCards), state.GetTerrorLevel()),
			fmt.Sprintf("Fear Till Next Card: %d", state.UnearnedFear),
			fmt.Sprintf("Fear Cards Till Next Terror Level: %d", state.countFearCardsTillNextTerrorLevel()),
			fmt.Sprintf("Earned Fear Cards: %d", len(state.EarnedFearCards)),
		}
	*/

	fearTable := tableWriter.Render()

	ravageLineContent := "<none>"
	if state.RavageSlot.IsCardPresent {
		ravageLineContent = state.RavageSlot.MaybeCard.String()
	}

	islandStatus := "Healthy"
	if state.IsBlightedIsland {
		islandStatus = "BLIGHTED"
	}
	blightedIslandLine := fmt.Sprintf(
		"                           Blight Pool: %d     Island: %v",
		state.BlightPool,
		islandStatus,
	)

	buildLineContent := "<none>"
	if state.BuildSlot.IsCardPresent {
		buildLineContent = state.BuildSlot.MaybeCard.String()
	}
	invasionLine := fmt.Sprintf(
		"     Ravage(%s) <- Build(%s) <- Deck(%d)",
		ravageLineContent,
		buildLineContent,
		len(state.RemainingInvaderDeck),
	)

	lines := []string{
		blightedIslandLine,
		fearTable,
		invasionLine,
	}

	return strings.Join(lines, "\n")
}

func countFearCardsTillNextTerrorLevel(state invader_board.InvaderBoardState) int {
	numEarnedFearCards := len(state.EarnedFearCards)
	numFearCardsTillNextTerrorLevel := math.MaxInt
	for _, terrorLevelThreshold := range state.TerrorLevelThresholds {
		numCardsTillTerrorLevel := terrorLevelThreshold - numEarnedFearCards
		if numCardsTillTerrorLevel > 0 && numCardsTillTerrorLevel < numFearCardsTillNextTerrorLevel {
			numFearCardsTillNextTerrorLevel = numCardsTillTerrorLevel
		}
	}
	return numFearCardsTillNextTerrorLevel
}
