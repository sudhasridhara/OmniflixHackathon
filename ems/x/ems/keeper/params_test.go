package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "ems/testutil/keeper"
	"ems/x/ems/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.EmsKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
