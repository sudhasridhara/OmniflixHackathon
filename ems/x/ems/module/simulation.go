package ems

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"ems/testutil/sample"
	emssimulation "ems/x/ems/simulation"
	"ems/x/ems/types"
)

// avoid unused import issue
var (
	_ = emssimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateEvent = "op_weight_msg_create_event"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateEvent int = 100

	opWeightMsgUpdateEvent = "op_weight_msg_update_event"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateEvent int = 100

	opWeightMsgDeleteEvent = "op_weight_msg_delete_event"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteEvent int = 100

	opWeightMsgPurchaseTicket = "op_weight_msg_purchase_ticket"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPurchaseTicket int = 100

	opWeightMsgResaleTicket = "op_weight_msg_resale_ticket"
	// TODO: Determine the simulation weight value
	defaultWeightMsgResaleTicket int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	emsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&emsGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateEvent int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateEvent, &weightMsgCreateEvent, nil,
		func(_ *rand.Rand) {
			weightMsgCreateEvent = defaultWeightMsgCreateEvent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateEvent,
		emssimulation.SimulateMsgCreateEvent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateEvent int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateEvent, &weightMsgUpdateEvent, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateEvent = defaultWeightMsgUpdateEvent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateEvent,
		emssimulation.SimulateMsgUpdateEvent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteEvent int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteEvent, &weightMsgDeleteEvent, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteEvent = defaultWeightMsgDeleteEvent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteEvent,
		emssimulation.SimulateMsgDeleteEvent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPurchaseTicket int
	simState.AppParams.GetOrGenerate(opWeightMsgPurchaseTicket, &weightMsgPurchaseTicket, nil,
		func(_ *rand.Rand) {
			weightMsgPurchaseTicket = defaultWeightMsgPurchaseTicket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPurchaseTicket,
		emssimulation.SimulateMsgPurchaseTicket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgResaleTicket int
	simState.AppParams.GetOrGenerate(opWeightMsgResaleTicket, &weightMsgResaleTicket, nil,
		func(_ *rand.Rand) {
			weightMsgResaleTicket = defaultWeightMsgResaleTicket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgResaleTicket,
		emssimulation.SimulateMsgResaleTicket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateEvent,
			defaultWeightMsgCreateEvent,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				emssimulation.SimulateMsgCreateEvent(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateEvent,
			defaultWeightMsgUpdateEvent,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				emssimulation.SimulateMsgUpdateEvent(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteEvent,
			defaultWeightMsgDeleteEvent,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				emssimulation.SimulateMsgDeleteEvent(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgPurchaseTicket,
			defaultWeightMsgPurchaseTicket,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				emssimulation.SimulateMsgPurchaseTicket(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgResaleTicket,
			defaultWeightMsgResaleTicket,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				emssimulation.SimulateMsgResaleTicket(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
