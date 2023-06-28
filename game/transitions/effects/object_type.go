package effects

// Useful for DRY gather/push mechanics
type ObjectType int

const (
	Dahan ObjectType = iota
	Explorer
	Town
	City
)
