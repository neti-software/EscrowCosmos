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

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedEscrow, found := k.Keeper.GetStoredEscrow(ctx, msg.EscrowIndex)

	if !found {
		panic("Escrow not found")
	}

	escrow, err := escrowCore.InitEscrow(storedEscrow.Escrow)

	if err != nil {
		panic(fmt.Sprintf(err.Error()))
	}

	depositAmount, err := escrow.Deposit(strategy.Address(msg.Creator))

	if err != nil {
		panic(fmt.Sprintf(err.Error()))
	}

	senderAddress, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		panic(fmt.Sprintf(err.Error()))
	}

	err = k.bank.SendCoinsFromAccountToModule(
		ctx,
		senderAddress,
		types.ModuleName,
		sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(depositAmount)))),
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

	return &types.MsgDepositResponse{}, nil
}
