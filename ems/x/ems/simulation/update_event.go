package simulation

import (
	"math/rand"

	"ems/x/ems/keeper"
	"ems/x/ems/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgUpdateEvent(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdateEvent{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UpdateEvent simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "UpdateEvent simulation not implemented"), nil, nil
	}
}
