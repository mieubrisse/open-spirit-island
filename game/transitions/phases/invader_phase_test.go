package phases

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEfficientlyDamageDahan_Normal(t *testing.T) {
	damageTaken := []int{0, 1, 1, 0}

	newDamageTaken := efficientlyDamageDahan(damageTaken, 2, 3)
	require.Equal(t, []int{1, 0}, newDamageTaken)
}

func TestEfficientlyDamageDahan_BuffedDahan(t *testing.T) {
	damageTaken := []int{0, 1, 2, 3, 4}

	newDamageTaken := efficientlyDamageDahan(damageTaken, 5, 5)
	require.Equal(t, []int{0, 1, 4}, newDamageTaken)
}
