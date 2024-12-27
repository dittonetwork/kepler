package simulation

import (
	"context"

	"github.com/cosmos/cosmos-sdk/simsx"

	"kepler/x/committees/keeper"
	"kepler/x/committees/types"
)

func MsgUpdateValidatorsStakesFactory(k keeper.Keeper) simsx.SimMsgFactoryFn[*types.MsgUpdateValidatorsStakes] {
	return func(ctx context.Context, testData *simsx.ChainDataSource, reporter simsx.SimulationReporter) ([]simsx.SimAccount, *types.MsgUpdateValidatorsStakes) {
		from := testData.AnyAccount(reporter)

		msg := &types.MsgUpdateValidatorsStakes{
			Creator: from.AddressBech32,
		}

		// TODO: Handle the UpdateValidatorsStakes simulation

		return []simsx.SimAccount{from}, msg
	}
}
