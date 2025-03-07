package simulation

import (
	"math/rand"

	"github.com/dittonetwork/kepler/x/workflow/keeper"
	"github.com/dittonetwork/kepler/x/workflow/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAddAutomation(
	_ types.AccountKeeper,
	_ types.BankKeeper,
	_ keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, _ *baseapp.BaseApp, _ sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddAutomation{
			Creator: simAccount.Address.String(),
		}

		// TODO: Implement simulation later in #98.
		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "AddAutomation simulation not implemented"), nil, nil
	}
}
