package utils

func GetMintInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetMaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
