package invader_board

import (
	"fmt"
	"github.com/imdario/mergo"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/blighted_island"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/fear"
	"github.com/mieubrisse/open-spirit-island/game/game_state/decks/invader_deck"
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

type MaybeInvaderCard struct {
	IsCardPresent bool
	MaybeCard     invader_deck.InvaderCard
}

type InvaderBoardState struct {
	UnearnedFear int
	EarnedFear   int

	// The number of fear cards needed to earn Terror Level 1/2/3/4
	TerrorLevelThresholds []int

	UnearnedFearCards []fear.FearCard // New cards are popped from the first element
	EarnedFearCards   []fear.FearCard // Processed from left to right

	// TODO blight card & pool!!

	BlightedIslandCard blighted_island.BlightedIslandCard
	IsBlightedIsland   bool
	BlightPool         int

	RemainingInvaderDeck []invader_deck.InvaderCard
	BuildSlot            MaybeInvaderCard
	RavageSlot           MaybeInvaderCard
	InvaderDeckDiscard   []invader_deck.InvaderCard
}

func (state InvaderBoardState) AddFear(fear int) InvaderBoardState {
	for i := 0; i < fear; i++ {
		state.UnearnedFear -= 1
		state.EarnedFear += 1

		// Earned a fear card
		if state.UnearnedFear == 0 {
			state.UnearnedFear = state.EarnedFear
			state.EarnedFear = 0

			if len(state.UnearnedFearCards) > 0 {
				state.EarnedFearCards = append(state.EarnedFearCards, state.UnearnedFearCards[0])
				state.UnearnedFearCards = state.UnearnedFearCards[1:]
			}
		}
	}
	return state
}

// Terror level
func (state InvaderBoardState) GetTerrorLevel() int {
	numEarnedFearCards := len(state.EarnedFearCards)
	highestIdxReached := 0
	for idx, terrorLevelThreshold := range state.TerrorLevelThresholds {
		if numEarnedFearCards >= terrorLevelThreshold {
			highestIdxReached = idx
		}
	}
	return highestIdxReached + 1 // Because terror level is 1-indexed
}

func (state InvaderBoardState) AdvanceInvaderCards() InvaderBoardState {
	if len(state.RemainingInvaderDeck) == 0 {
		return state
	}

	state.RavageSlot = state.BuildSlot
	state.BuildSlot = MaybeInvaderCard{
		IsCardPresent: true,
		MaybeCard:     state.RemainingInvaderDeck[0],
	}
	state.RemainingInvaderDeck = state.RemainingInvaderDeck[1:]
	return state
}

func (state InvaderBoardState) String() string {
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
			fmt.Sprintf("Cards Till Next: %d", state.countFearCardsTillNextTerrorLevel()),
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
		invasionLine,
		fearTable,
		blightedIslandLine,
	}

	return strings.Join(lines, "\n")
}

func (state InvaderBoardState) countFearCardsTillNextTerrorLevel() int {
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
