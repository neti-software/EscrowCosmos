package keeper

import (
	"context"
	"strconv"

	"escrow/x/escrow/escrowCore"
	"escrow/x/escrow/escrowCore/strategy"
	"escrow/x/escrow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Confirm(goCtx context.Context, msg *types.MsgConfirm) (*types.MsgConfirmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedEscrow, found := k.Keeper.GetStoredEscrow(ctx, msg.EscrowIndex)

	if !found {
		panic("Escrow not found")
	}

	escrow, err := escrowCore.InitEscrow(storedEscrow.Escrow)

	if err != nil {
		panic(err)
	}

	err = escrow.Confirm(strategy.Address(msg.Creator), strategy.Address(msg.Winner))

	if err != nil {
		panic(err)
	}

	serializedEscrow, err := escrow.Serialize()

	if err != nil {
		panic(err)
	}

	stringIndex := strconv.FormatUint(uint64(escrow.Index), 10)

	k.Keeper.SetStoredEscrow(ctx, types.StoredEscrow{
		Index:  stringIndex,
		Escrow: serializedEscrow,
	})
	_ = ctx

	return &types.MsgConfirmResponse{}, nil
}
