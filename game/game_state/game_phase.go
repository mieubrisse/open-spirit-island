package game_state

type GamePhase int

const (
	SpiritGrow GamePhase = iota
	SpiritGainTrackBenefits
	SpiritPowerPlays
	FastPower
	Invader
	/*
		InvaderBlightedIsland
		InvaderFear
		InvaderRavage
		InvaderBuild
		InvaderExplore
		InvaderAdvance
	*/
	SlowPower
	TimePasses
)
