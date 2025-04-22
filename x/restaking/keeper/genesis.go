package keeper

import (
	"errors"

	"cosmossdk.io/collections"
	sdkerrors "cosmossdk.io/errors"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/crypto"
	cmttypes "github.com/cometbft/cometbft/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// WriteValidators returns a slice of bonded genesis validators.
func (k Keeper) WriteValidators(ctx sdk.Context) ([]cmttypes.GenesisValidator, error) {
	operators, err := k.repository.GetBondedValidators(ctx)
	if err != nil {
		return nil, err
	}

	validators := make([]cmttypes.GenesisValidator, 0, len(operators))

	for _, operator := range operators {
		if err = operator.UnpackInterfaces(k.cdc); err != nil {
			return nil, err
		}

		var pk cryptotypes.PubKey
		pk, err = operator.ConsPubKey()
		if err != nil {
			return nil, err
		}

		var cmtPk crypto.PubKey
		cmtPk, err = cryptocodec.ToCmtPubKeyInterface(pk)
		if err != nil {
			return nil, err
		}

		validators = append(validators, cmttypes.GenesisValidator{
			Address: sdk.ConsAddress(cmtPk.Address()).Bytes(),
			PubKey:  cmtPk,
			Power:   operator.VotingPower,
		})
	}

	return validators, nil
}

// InitGenesis initializes the genesis state of the restaking module.
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) ([]abci.ValidatorUpdate, error) {
	if err := k.repository.SetLastUpdate(ctx, state.LastUpdate); err != nil {
		return nil, sdkerrors.Wrap(err, "last update")
	}

	for _, operator := range state.PendingValidators {
		if err := k.repository.SetPendingOperator(ctx, operator.Address, operator); err != nil {
			return nil, sdkerrors.Wrap(err, "set operator")
		}
	}

	for _, validator := range state.Validators {
		if err := k.repository.AddValidatorsChange(ctx, validator, types.ValidatorChangeTypeCreate); err != nil {
			return nil, sdkerrors.Wrap(err, "validators")
		}
	}

	return k.ApplyAndReturnValidatorSetUpdates(ctx)
}

func (k Keeper) ExportGenesis(ctx sdk.Context) (*types.GenesisState, error) {
	genesis := types.DefaultGenesis()

	var err error
	genesis.LastUpdate, err = k.repository.GetLastUpdate(ctx)
	if err != nil && !errors.Is(err, collections.ErrNotFound) {
		return genesis, sdkerrors.Wrap(err, "last update")
	}

	genesis.Validators, err = k.repository.GetAllValidators(ctx)
	if err != nil {
		return genesis, err
	}

	genesis.PendingValidators, err = k.repository.GetPendingOperators(ctx)

	return genesis, err
}
