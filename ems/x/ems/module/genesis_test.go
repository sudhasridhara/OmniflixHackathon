package ems_test

import (
	"testing"

	keepertest "ems/testutil/keeper"
	"ems/testutil/nullify"
	ems "ems/x/ems/module"
	"ems/x/ems/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EmsKeeper(t)
	ems.InitGenesis(ctx, k, genesisState)
	got := ems.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
