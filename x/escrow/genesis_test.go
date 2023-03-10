package escrow_test

import (
	"testing"

	keepertest "escrow/testutil/keeper"
	"escrow/testutil/nullify"
	"escrow/x/escrow"
	"escrow/x/escrow/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: types.SystemInfo{
			NextId: 80,
		},
		StoredEscrowList: []types.StoredEscrow{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EscrowKeeper(t)
	escrow.InitGenesis(ctx, *k, genesisState)
	got := escrow.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	require.ElementsMatch(t, genesisState.StoredEscrowList, got.StoredEscrowList)
	// this line is used by starport scaffolding # genesis/test/assert
}
