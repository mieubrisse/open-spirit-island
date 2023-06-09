package game_state

import (
	"github.com/mieubrisse/open-spirit-island/game_state/island"
	"github.com/mieubrisse/open-spirit-island/game_state/island/land_state"
	"github.com/mieubrisse/open-spirit-island/game_state/island/land_type"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGameState_RunInvaderPhase(t *testing.T) {
	game := NewTestGame()

	game = game.RunInvaderPhase()

	// At this point -- Ravage: none, Build: jungle
	require.False(t, game.InvaderState.RavageSlot.IsCardPresent)
	require.True(t, game.InvaderState.BuildSlot.IsCardPresent)
	require.Equal(
		t,
		[]land_state.LandState{
			// 0
			{
				LandType: land_type.Ocean,
			},
			// 1
			{
				LandType:    land_type.Mountain,
				NumPresence: 1,
			},
			// 2
			{
				LandType:    land_type.Wetlands,
				CityHealth:  []int{island.CityBaseHealth},
				DahanHealth: []int{island.DahanBaseHealth},
			},
			// 3
			{
				LandType:       land_type.Jungle,
				DahanHealth:    []int{island.DahanBaseHealth, island.DahanBaseHealth},
				ExplorerHealth: []int{island.ExplorerBaseHealth},
			},
			// 4
			{
				LandType:  land_type.Desert,
				NumBlight: 1,
			},
			// 5
			{
				LandType: land_type.Wetlands,
			},
			// 6
			{
				LandType:    land_type.Mountain,
				DahanHealth: []int{island.DahanBaseHealth},
			},
			// 7
			{
				LandType:    land_type.Desert,
				DahanHealth: []int{island.DahanBaseHealth, island.DahanBaseHealth},
			},
			// 8
			{
				LandType:       land_type.Jungle,
				TownHealth:     []int{island.TownBaseHealth},
				ExplorerHealth: []int{island.ExplorerBaseHealth},
			},
		},
		game.BoardState.Lands,
	)

	game = game.RunInvaderPhase()

	// At this point -- Ravage: jungle, build: desert
	require.True(t, game.InvaderState.RavageSlot.IsCardPresent)
	require.True(t, game.InvaderState.BuildSlot.IsCardPresent)
	require.Equal(
		t,
		[]land_state.LandState{
			// 0
			{
				LandType: land_type.Ocean,
			},
			// 1
			{
				LandType:    land_type.Mountain,
				NumPresence: 1,
			},
			// 2
			{
				LandType:    land_type.Wetlands,
				CityHealth:  []int{island.CityBaseHealth},
				DahanHealth: []int{island.DahanBaseHealth},
			},
			// 3
			{
				LandType:       land_type.Jungle,
				DahanHealth:    []int{island.DahanBaseHealth, island.DahanBaseHealth},
				ExplorerHealth: []int{island.ExplorerBaseHealth},
				TownHealth:     []int{island.TownBaseHealth},
			},
			// 4
			{
				LandType:       land_type.Desert,
				NumBlight:      1,
				ExplorerHealth: []int{island.ExplorerBaseHealth},
			},
			// 5
			{
				LandType: land_type.Wetlands,
			},
			// 6
			{
				LandType:    land_type.Mountain,
				DahanHealth: []int{island.DahanBaseHealth},
			},
			// 7
			{
				LandType:       land_type.Desert,
				DahanHealth:    []int{island.DahanBaseHealth, island.DahanBaseHealth},
				ExplorerHealth: []int{island.ExplorerBaseHealth},
			},
			// 8
			{
				LandType:       land_type.Jungle,
				TownHealth:     []int{island.TownBaseHealth},
				ExplorerHealth: []int{island.ExplorerBaseHealth},
				CityHealth:     []int{island.CityBaseHealth},
			},
		},
		game.BoardState.Lands,
	)

	game = game.RunInvaderPhase()

	// At this point -- ravage: desert, build: wetlands
	require.True(t, game.InvaderState.RavageSlot.IsCardPresent)
	require.True(t, game.InvaderState.BuildSlot.IsCardPresent)
	require.Equal(
		t,
		[]land_state.LandState{
			// 0
			{
				LandType: land_type.Ocean,
			},
			// 1
			{
				LandType:    land_type.Mountain,
				NumPresence: 1,
			},
			// 2
			{
				LandType:       land_type.Wetlands,
				CityHealth:     []int{island.CityBaseHealth},
				DahanHealth:    []int{island.DahanBaseHealth},
				ExplorerHealth: []int{island.ExplorerBaseHealth},
			},
			// 3
			{
				LandType:       land_type.Jungle,
				DahanHealth:    []int{1},
				ExplorerHealth: []int{island.ExplorerBaseHealth},
				TownHealth:     []int{island.TownBaseHealth},
				NumBlight:      1,
			},
			// 4
			{
				LandType:       land_type.Desert,
				NumBlight:      1,
				ExplorerHealth: []int{island.ExplorerBaseHealth},
				TownHealth:     []int{island.TownBaseHealth},
			},
			// 5
			{
				LandType:       land_type.Wetlands,
				ExplorerHealth: []int{island.ExplorerBaseHealth},
			},
			// 6
			{
				LandType:    land_type.Mountain,
				DahanHealth: []int{island.DahanBaseHealth},
			},
			// 7
			{
				LandType:       land_type.Desert,
				DahanHealth:    []int{island.DahanBaseHealth, island.DahanBaseHealth},
				ExplorerHealth: []int{island.ExplorerBaseHealth},
				TownHealth:     []int{island.TownBaseHealth},
			},
			// 8
			{
				LandType:       land_type.Jungle,
				TownHealth:     []int{island.TownBaseHealth},
				ExplorerHealth: []int{island.ExplorerBaseHealth},
				CityHealth:     []int{island.CityBaseHealth},
				NumBlight:      1,
			},
		},
		game.BoardState.Lands,
	)

	// Can't go further because we need user input for the Blight cascade that occurs on Desert 4
}

func TestGameState_efficientlyDamageDahan(t *testing.T) {
	// No Dahan
	require.Equal(t, []int(nil), efficientlyDamageDahan([]int{}, 10))

	// No damage
	require.Equal(t, []int{2, 1}, efficientlyDamageDahan([]int{2, 1}, 0))

	// Efficient culling
	require.Equal(t, []int{1, 2}, efficientlyDamageDahan([]int{2, 1, 2, 1, 1}, 4))

	// Strong Dahan
	require.Equal(t, []int{1, 4}, efficientlyDamageDahan([]int{4, 4}, 3))
}
