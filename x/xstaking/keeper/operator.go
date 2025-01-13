package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	errorsmod "cosmossdk.io/errors"
	stakingtypes "cosmossdk.io/x/staking/types"

	"kepler/x/xstaking/types"
)

func (k Keeper) DecreaseStake(ctx context.Context, props types.DecreaseStakeProps) error {
	opAddr, err := k.validatorAddressCodec.StringToBytes(props.OperatorAddress)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid operator address: %s", err)
	}

	if !props.Amount.IsValid() || !props.Amount.Amount.IsPositive() {
		return errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"invalid shares amount",
		)
	}

	bondDenom, err := k.stakingKeeper.BondDenom(ctx)
	if err != nil {
		return err
	}

	if props.Amount.Denom != bondDenom {
		return errorsmod.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"invalid coin denomination: got %s, expected %s", props.Amount.Denom, bondDenom,
		)
	}

	shares, err := k.stakingKeeper.ValidateUnbondAmount(ctx, opAddr, opAddr, props.Amount.Amount)
	if err != nil {
		return err
	}

	_, err = k.stakingKeeper.GetValidator(ctx, opAddr)
	if err != nil {
		return err
	}

	amount, err := k.stakingKeeper.Unbond(ctx, opAddr, opAddr, shares)
	if err != nil {
		return err
	}

	err = k.bankKeeper.BurnCoins(ctx, opAddr, sdk.NewCoins(sdk.NewCoin(bondDenom, amount)))
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) IncreaseStake(ctx context.Context, props types.IncreaseStakeProps) error {
	opAddr, err := k.validatorAddressCodec.StringToBytes(props.OperatorAddress)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid operator address: %s", err)
	}

	validator, err := k.stakingKeeper.GetValidator(ctx, opAddr)
	if err != nil {
		return err
	}

	bondDenom, err := k.stakingKeeper.BondDenom(ctx)
	if err != nil {
		return err
	}

	if props.Amount.Denom != bondDenom {
		return errorsmod.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"invalid coin denomination: got %s, expected %s", props.Amount.Denom, bondDenom,
		)
	}

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(props.Amount))
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, opAddr, sdk.NewCoins(props.Amount))
	if err != nil {
		return err
	}

	_, err = k.stakingKeeper.Delegate(ctx, opAddr, props.Amount.Amount, stakingtypes.Unbonded, validator, true)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) RegisterOperator(ctx context.Context, props types.RegisterOperatorProps) error {
	opAddr, err := k.validatorAddressCodec.StringToBytes(props.OperatorAddress)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid operator address: %s", err)
	}

	minCommRate, err := k.minCommissionRate(ctx)
	if err != nil {
		return err
	}

	if props.Commission.Rate.LT(minCommRate) {
		return errorsmod.Wrapf(
			stakingtypes.ErrCommissionLTMinRate,
			"cannot set validator commission to less than minimum rate of %s",
			minCommRate,
		)
	}

	_, err = k.stakingKeeper.GetValidator(ctx, opAddr)
	if err == nil {
		return errorsmod.Wrapf(
			stakingtypes.ErrValidatorOwnerExists,
			"cosmos staking module error: %s", err,
		)
	}

	pubKeyTypes, err := k.consensusKeeper.ValidatorPubKeyTypes(ctx)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "failed to query consensus params: %s", err)
	}

	if err = validatePubKey(props.PubKey, pubKeyTypes); err != nil {
		return err
	}

	bondDenom, err := k.stakingKeeper.BondDenom(ctx)
	if err != nil {
		return err
	}

	if props.StakeValue.Denom != bondDenom {
		return errorsmod.Wrapf(
			sdkerrors.ErrInvalidRequest, "invalid coin denomination: got %s, expected %s", props.StakeValue.Denom, bondDenom,
		)
	}

	validator, err := stakingtypes.NewValidator(props.OperatorAddress, props.PubKey, props.Description)
	if err != nil {
		return err
	}

	operatorAddr, err := k.validatorAddressCodec.StringToBytes(validator.GetOperator())
	if err != nil {
		return err
	}

	err = k.stakingKeeper.Validators.Set(ctx, operatorAddr, validator)
	if err != nil {
		return err
	}

	err = k.stakingKeeper.SetValidatorByConsAddr(ctx, validator)
	if err != nil {
		return err
	}

	err = k.stakingKeeper.SetNewValidatorByPowerIndex(ctx, validator)
	if err != nil {
		return err
	}

	err = k.stakingKeeper.Hooks().AfterValidatorCreated(ctx, operatorAddr)
	if err != nil {
		return err
	}

	_, err = k.stakingKeeper.Delegate(ctx, operatorAddr, props.StakeValue.Amount, stakingtypes.Unbonded, validator, true)
	if err != nil {
		return err
	}

	return nil
}
