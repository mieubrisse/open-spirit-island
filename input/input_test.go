package input

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBase26Encode(t *testing.T) {
	require.Equal(t, "d", base26Encode(3))
	require.Equal(t, "a", base26Encode(0))
	require.Equal(t, "z", base26Encode(25))
}

func TestBase26Decode(t *testing.T) {
	result, err := base26Decode("a")
	require.NoError(t, err)
	require.Equal(t, 0, result)

	result, err = base26Decode("z")
	require.NoError(t, err)
	require.Equal(t, 25, result)

	result, err = base26Decode("aa")
	require.Error(t, err)

	result, err = base26Decode("")
	require.Error(t, err)
}
