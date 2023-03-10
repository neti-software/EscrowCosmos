package keeper_test

import (
	"context"
	"testing"

	keepertest "escrow/testutil/keeper"
	"escrow/x/escrow/keeper"
	"escrow/x/escrow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EscrowKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
