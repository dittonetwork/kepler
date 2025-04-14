package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// GetValidator returns the validator by address.
func (k Keeper) GetValidator(ctx sdk.Context, addr sdk.ValAddress) (types.Validator, error) {
	return k.repository.GetValidator(ctx, addr)
}
