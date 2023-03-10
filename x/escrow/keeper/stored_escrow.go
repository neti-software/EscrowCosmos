package keeper

import (
	"escrow/x/escrow/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStoredEscrow set a specific storedEscrow in the store from its index
func (k Keeper) SetStoredEscrow(ctx sdk.Context, storedEscrow types.StoredEscrow) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredEscrowKeyPrefix))
	b := k.cdc.MustMarshal(&storedEscrow)
	store.Set(types.StoredEscrowKey(
		storedEscrow.Index,
	), b)
}

// GetStoredEscrow returns a storedEscrow from its index
func (k Keeper) GetStoredEscrow(
	ctx sdk.Context,
	index string,

) (val types.StoredEscrow, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredEscrowKeyPrefix))

	b := store.Get(types.StoredEscrowKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredEscrow removes a storedEscrow from the store
func (k Keeper) RemoveStoredEscrow(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredEscrowKeyPrefix))
	store.Delete(types.StoredEscrowKey(
		index,
	))
}

// GetAllStoredEscrow returns all storedEscrow
func (k Keeper) GetAllStoredEscrow(ctx sdk.Context) (list []types.StoredEscrow) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredEscrowKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredEscrow
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
