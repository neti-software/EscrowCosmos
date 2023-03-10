package keeper

import (
	"context"
	"strconv"

	"escrow/x/escrow/escrowCore"
	"escrow/x/escrow/escrowCore/strategy"
	"escrow/x/escrow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) NewEscrow(goCtx context.Context, msg *types.MsgNewEscrow) (*types.MsgNewEscrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}

	instigatorWager, instigatorWagerErr := strconv.ParseInt(msg.InstigatorWager, 10, 64)
	riderWager, riderWagerErr := strconv.ParseInt(msg.RiderWager, 10, 64)

	if instigatorWagerErr != nil {
		panic(instigatorWagerErr)
	}

	if riderWagerErr != nil {
		panic(riderWagerErr)
	}

	newEscrow, err := escrowCore.NewEscrow(
		strategy.Index(systemInfo.NextId),
		strategy.Strategies[msg.Strategy],
		strategy.Address(msg.Instigator),
		strategy.Coin(instigatorWager),
		strategy.Address(msg.Rider),
		strategy.Coin(riderWager),
	)

	if err != nil {
		panic(err)
	}

	nextId := strconv.FormatUint(systemInfo.NextId, 10)

	k.Keeper.SetStoredEscrow(ctx, types.StoredEscrow{
		Index:  nextId,
		Escrow: newEscrow,
	})

	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	_ = ctx

	return &types.MsgNewEscrowResponse{
		EscrowIndex: nextId,
	}, nil
}
