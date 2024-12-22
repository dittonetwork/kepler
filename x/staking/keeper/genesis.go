package keeper

import (
	"context"
	"cosmossdk.io/core/appmodule"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"kepler/x/staking/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) ([]appmodule.ValidatorUpdate, error) {
	if err := k.Params.Set(ctx, genState.Params); err != nil {
		return nil, err
	}

	if err := k.LastTotalPower.Set(ctx, genState.LastTotalPower); err != nil {
		return nil, err
	}

	for _, validator := range genState.Validators {
		if err := k.SetValidator(ctx, validator); err != nil {
			return nil, err
		}

		// Manually set indices for the first time
		if err := k.SetValidatorByConsAddr(ctx, validator); err != nil {
			return nil, err
		}

		if err := k.SetValidatorByPowerIndex(ctx, validator); err != nil {
			return nil, err
		}

		// Call the creation hook if not exported
		if !genState.Exported {
			valbz, err := k.ValidatorAddressCodec().StringToBytes(validator.GetOperator())
			if err != nil {
				return nil, err
			}
			if err := k.Hooks().AfterValidatorCreated(ctx, valbz); err != nil {
				return nil, err
			}
		}

		// update timeslice if necessary
		if validator.IsUnbonding() {
			if err := k.InsertUnbondingValidatorQueue(ctx, validator); err != nil {
				return nil, err
			}
		}
	}

	// don't need to run CometBFT updates if we exported
	var moduleValidatorUpdates []appmodule.ValidatorUpdate
	if genState.Exported {
		for _, lv := range genState.LastValidatorPowers {
			valAddr, err := k.validatorAddressCodec.StringToBytes(lv.Address)
			if err != nil {
				return nil, err
			}

			err = k.SetLastValidatorPower(ctx, valAddr, lv.Power)
			if err != nil {
				return nil, err
			}

			validator, err := k.GetValidator(ctx, valAddr)
			if err != nil {
				return nil, fmt.Errorf("validator %s not found", lv.Address)
			}

			powerReduction, err := k.PowerReduction(ctx)
			if err != nil {
				return nil, err
			}

			update := validator.ModuleValidatorUpdate(powerReduction)
			update.Power = lv.Power // keep the next-val-set offset, use the last powerReduction for the first block
			moduleValidatorUpdates = append(moduleValidatorUpdates, update)
		}
	} else {
		var err error

		moduleValidatorUpdates, err = k.ApplyAndReturnValidatorSetUpdates(ctx)
		if err != nil {
			return nil, err
		}
	}

	return moduleValidatorUpdates, nil
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var err error
	var fnErr error

	var lastValidatorPowers []types.LastValidatorPower

	err = k.IterateLastValidatorPowers(ctx, func(addr sdk.ValAddress, power int64) (stop bool) {
		addrStr, err := k.validatorAddressCodec.BytesToString(addr)
		if err != nil {
			fnErr = err
			return true
		}
		lastValidatorPowers = append(lastValidatorPowers, types.LastValidatorPower{Address: addrStr, Power: power})
		return false
	})
	if err != nil {
		return nil, err
	}
	if fnErr != nil {
		return nil, fnErr
	}

	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	totalPower, err := k.LastTotalPower.Get(ctx)
	if err != nil {
		return nil, err
	}

	allValidators, err := k.GetAllValidators(ctx)
	if err != nil {
		return nil, err
	}

	return &types.GenesisState{
		Params:              params,
		LastTotalPower:      totalPower,
		LastValidatorPowers: lastValidatorPowers,
		Exported:            true,
		Validators:          allValidators,
	}, nil
}
