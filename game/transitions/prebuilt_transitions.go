package transitions

/*
var ReclaimAllCardsTransition = GameStateTransition{
	ReadableStr: "All✋",
	TransitionFunction: func(state game_state.GameState) game_state.GameState {
		state.PlayerState.Hand = append(state.PlayerState.Hand, state.PlayerState.Discard...)
		state.PlayerState.Discard = []power.PowerCard{}

		return state
	},
}

func NewNormalAddPresenceTransition(addRange int) GameStateTransition {
	transitionFunc := func(state game_state.GameState) game_state.GameState {
		landIdxOptionsSet := state.BoardState.FilterLands(filter.IslandFilter{
			SourceFilter: filter.LandFilter{
				PresenceMin: 1,
			},
			MinRange: 0,
			MaxRange: addRange,
			TargetFilter: filter.LandFilter{
				LandTypes: land_type.NonOceanLandTypes,
			},
		})
		landIdxOptions := landIdxOptionsSet.Slice()
		sort.Ints(landIdxOptions)

		options := make([]string, len(landIdxOptions))
		for idx, landIdx := range landIdxOptions {
			land := state.BoardState.Lands[landIdx]
			suffixStr := strings.Repeat("🪔", land.NumPresence)
			options[idx] = fmt.Sprintf("%s #%d\t%s", land.LandType.String(), landIdx, suffixStr)
		}

		selection := input.GetSingleSelection("Choose a land for +🪔:", options)
		selectedLandIdx := landIdxOptions[selection]

		state.BoardState.Lands[selectedLandIdx].NumPresence++

		return state
	}
	return GameStateTransition{
		ReadableStr:        fmt.Sprintf("+🪔 —%d→", addRange),
		TransitionFunction: transitionFunc,
	}
}

func NewGainEnergyTransition(energy int) GameStateTransition {
	transitionFunc := func(state game_state.GameState) game_state.GameState {
		state.PlayerState.Energy += energy
		return state
	}
	return GameStateTransition{
		ReadableStr:        fmt.Sprintf("+%d⚡", energy),
		TransitionFunction: transitionFunc,
	}
}

*/
