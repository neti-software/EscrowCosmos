package keeper

import (
	"context"

	"escrow/x/escrow/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StoredEscrowAll(c context.Context, req *types.QueryAllStoredEscrowRequest) (*types.QueryAllStoredEscrowResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedEscrows []types.StoredEscrow
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	storedEscrowStore := prefix.NewStore(store, types.KeyPrefix(types.StoredEscrowKeyPrefix))

	pageRes, err := query.Paginate(storedEscrowStore, req.Pagination, func(key []byte, value []byte) error {
		var storedEscrow types.StoredEscrow
		if err := k.cdc.Unmarshal(value, &storedEscrow); err != nil {
			return err
		}

		storedEscrows = append(storedEscrows, storedEscrow)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredEscrowResponse{StoredEscrow: storedEscrows, Pagination: pageRes}, nil
}

func (k Keeper) StoredEscrow(c context.Context, req *types.QueryGetStoredEscrowRequest) (*types.QueryGetStoredEscrowResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStoredEscrow(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStoredEscrowResponse{StoredEscrow: val}, nil
}
