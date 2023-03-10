package keeper

import (
	"escrow/x/escrow/types"
)

var _ types.QueryServer = Keeper{}
