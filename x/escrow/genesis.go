package escrow

import (
	"escrow/x/escrow/keeper"
	"escrow/x/escrow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	k.SetSystemInfo(ctx, genState.SystemInfo)

	// Set all the storedEscrow
	for _, elem := range genState.StoredEscrowList {
		k.SetStoredEscrow(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all systemInfo
	systemInfo, found := k.GetSystemInfo(ctx)
	if found {
		genesis.SystemInfo = systemInfo
	}
	genesis.StoredEscrowList = k.GetAllStoredEscrow(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
