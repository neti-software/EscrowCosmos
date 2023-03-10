package keeper_test

import (
	"strconv"
	"testing"

	keepertest "escrow/testutil/keeper"
	"escrow/testutil/nullify"
	"escrow/x/escrow/keeper"
	"escrow/x/escrow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStoredEscrow(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredEscrow {
	items := make([]types.StoredEscrow, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredEscrow(ctx, items[i])
	}
	return items
}

func TestStoredEscrowGet(t *testing.T) {
	keeper, ctx := keepertest.EscrowKeeper(t)
	items := createNStoredEscrow(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredEscrow(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredEscrowRemove(t *testing.T) {
	keeper, ctx := keepertest.EscrowKeeper(t)
	items := createNStoredEscrow(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredEscrow(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredEscrow(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredEscrowGetAll(t *testing.T) {
	keeper, ctx := keepertest.EscrowKeeper(t)
	items := createNStoredEscrow(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredEscrow(ctx)),
	)
}
