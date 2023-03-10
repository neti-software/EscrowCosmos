package keeper_test

import (
	"testing"

	testkeeper "escrow/testutil/keeper"
	"escrow/x/escrow/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EscrowKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
