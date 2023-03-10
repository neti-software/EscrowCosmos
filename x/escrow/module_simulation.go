package escrow

import (
	"math/rand"

	"escrow/testutil/sample"
	escrowsimulation "escrow/x/escrow/simulation"
	"escrow/x/escrow/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = escrowsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgNewEscrow = "op_weight_msg_new_escrow"
	// TODO: Determine the simulation weight value
	defaultWeightMsgNewEscrow int = 100

	opWeightMsgApprove = "op_weight_msg_approve"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApprove int = 100

	opWeightMsgDeposit = "op_weight_msg_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeposit int = 100

	opWeightMsgConfirm = "op_weight_msg_confirm"
	// TODO: Determine the simulation weight value
	defaultWeightMsgConfirm int = 100

	opWeightMsgWithdraw = "op_weight_msg_withdraw"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdraw int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	escrowGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&escrowGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgNewEscrow int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgNewEscrow, &weightMsgNewEscrow, nil,
		func(_ *rand.Rand) {
			weightMsgNewEscrow = defaultWeightMsgNewEscrow
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgNewEscrow,
		escrowsimulation.SimulateMsgNewEscrow(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApprove int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApprove, &weightMsgApprove, nil,
		func(_ *rand.Rand) {
			weightMsgApprove = defaultWeightMsgApprove
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApprove,
		escrowsimulation.SimulateMsgApprove(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeposit, &weightMsgDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgDeposit = defaultWeightMsgDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeposit,
		escrowsimulation.SimulateMsgDeposit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgConfirm int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgConfirm, &weightMsgConfirm, nil,
		func(_ *rand.Rand) {
			weightMsgConfirm = defaultWeightMsgConfirm
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgConfirm,
		escrowsimulation.SimulateMsgConfirm(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdraw int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdraw, &weightMsgWithdraw, nil,
		func(_ *rand.Rand) {
			weightMsgWithdraw = defaultWeightMsgWithdraw
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdraw,
		escrowsimulation.SimulateMsgWithdraw(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
