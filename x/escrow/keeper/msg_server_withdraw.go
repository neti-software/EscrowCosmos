package keeper

import (
	"context"
	"fmt"
	"strconv"

	"escrow/x/escrow/escrowCore"
	"escrow/x/escrow/escrowCore/strategy"
	"escrow/x/escrow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
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

	if escrow.ReleasedByAgreementIndex != escrowCore.NullIndex {
		releaseAgreementIndex := strconv.FormatInt(int64(escrow.ReleasedByAgreementIndex), 10)
		storedReleaseEscrow, found := k.Keeper.GetStoredEscrow(ctx, releaseAgreementIndex)
		if !found {
			panic("Escrow not found")
		}
		releaseEscrow, err = escrowCore.InitEscrow(storedReleaseEscrow.Escrow)
		if err != nil {
			panic(err)
		}
	}

	withdrawAmount, err := escrow.Withdraw(strategy.Address(msg.Creator), &releaseEscrow)

	if err != nil {
		panic(err)
	}

	senderAddress, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		panic(fmt.Sprintf(err.Error()))
	}

	err = k.bank.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		senderAddress,
		sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(withdrawAmount)))),
	)
	if err != nil {
		panic(fmt.Sprintf(err.Error()))
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

	return &types.MsgWithdrawResponse{}, nil
}
