package keeper

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtprotocrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

// ApplyAndReturnValidatorSetUpdates applies and return accumulated updates to the bonded validator set.
// Notes: we always return the full set of validators, even if they are not updated.
func (k Keeper) ApplyAndReturnValidatorSetUpdates(ctx context.Context) ([]abci.ValidatorUpdate, error) {
	validators, err := k.repository.GetBondedValidators(sdk.UnwrapSDKContext(ctx))

	if err != nil {
		k.Logger().With("error", err).Error("failed to collect bonded validators")
		return nil, err
	}

	updates := make([]abci.ValidatorUpdate, 0, len(validators))

	for _, validator := range validators {
		if validator.Status != types.Bonded {
			k.Logger().
				With("validator", validator.OperatorAddress).
				Warn("validator is not bonded, skipping update")
			continue
		}

		var tmProtoPk cmtprotocrypto.PublicKey
		tmProtoPk, err = validator.CmtConsPublicKey()
		if err != nil {
			k.Logger().With("error", err, "validator", validator.OperatorAddress).
				Error("failed to get consensus public key")
			continue
		}

		update := abci.ValidatorUpdate{
			PubKey: tmProtoPk,
			Power:  validator.VotingPower,
		}

		updates = append(updates, update)
	}

	return updates, nil
}
