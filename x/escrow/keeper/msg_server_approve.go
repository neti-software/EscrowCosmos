package keeper

import (
	"context"
	"escrow/x/escrow/escrowCore"
	"escrow/x/escrow/escrowCore/strategy"
	"escrow/x/escrow/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Approve(goCtx context.Context, msg *types.MsgApprove) (*types.MsgApproveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedEscrow, found := k.Keeper.GetStoredEscrow(ctx, msg.EscrowIndex)

	if !found {
		panic("Escrow not found")
	}

	escrow, err := escrowCore.InitEscrow(storedEscrow.Escrow)

	if err != nil {
		panic(err)
	}

	var releaseEscrow escrowCore.Escrow

	if escrow.ReleaseAgreementIndex != escrowCore.NullIndex {
		releaseAgreementIndex := strconv.FormatInt(int64(escrow.ReleaseAgreementIndex), 10)
		storedReleaseEscrow, found := k.Keeper.GetStoredEscrow(ctx, releaseAgreementIndex)
		if !found {
			panic("Escrow not found")
		}
		releaseEscrow, err = escrowCore.InitEscrow(storedReleaseEscrow.Escrow)
		if err != nil {
			panic(err)
		}
	}

	err = escrow.Approve(strategy.Address(msg.Creator), &releaseEscrow)

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

	return &types.MsgApproveResponse{}, nil
}
