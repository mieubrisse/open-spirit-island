package island

// This is a selector, for use with the board to grab lands based on certain conditions
type LandSelector struct {
	// Only lands matching these types will be considered as sources
	SourceLandTypes map[LandType]bool

	// Only cities that match all these minimums will be considered as sources
	SourceExplorersMin int
	SourceTownsMin     int
	SourceCitiesMin    int
	// TODO Dahan as source??

	SourcePresenceMin int

	MinRange int
	MaxRange int

	TargetLandTypes map[LandType]bool

	TargetExplorersMin int
	TargetTownsMin     int
	TargetCitiesMin    int

	TargetDahanMin int
}
