package keeper

import (
	"context"
	"errors"
	contracts "kepler/x/symbiotic/abi"
	"kepler/x/symbiotic/types"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (k Keeper) UpdateStakedAmounts(ctx context.Context) error {
	if !k.SyncNeeded(ctx) {
		k.Logger().Debug("no need to update stakes yet")
		return nil
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.Logger().Debug("updating stakes amount", "height", sdkCtx.BlockHeight())

	blkHash, err := k.GetFinalizedBlockHash(ctx)
	if err != nil {
		return err
	}

	amountsInfo, err := k.fetchStakedAmountsFromBlock(ctx, blkHash)
	if errors.Is(err, types.ErrContractNotInStorage) {
		k.Logger().Warn("contract address is not set, stakes won't be updated")
		return nil
	}
	if err != nil {
		return err
	}

	for _, stakedAmountInfo := range amountsInfo {
		prevInfo, found := k.GetStakedAmountInfo(ctx, stakedAmountInfo.EthereumAddress)
		if found && prevInfo.StakedAmount == stakedAmountInfo.StakedAmount {
			continue
		}
		stakedAmountInfo.LastUpdated = uint64(sdk.UnwrapSDKContext(ctx).BlockTime().Unix())
		k.SetStakedAmountInfo(ctx, stakedAmountInfo)
	}

	k.Logger().Info("successfully updated staked amounts")
	return nil
}

func (k Keeper) fetchStakedAmountsFromBlock(ctx context.Context, blkHash common.Hash) ([]types.StakedAmountInfo, error) {
	contractAddr, exists := k.GetContractAddress(ctx)
	if !exists {
		return nil, types.ErrContractNotInStorage
	}

	var stakes []types.StakedAmountInfo

	if len(contractAddr.Address) == 0 {
		k.Logger().Warn("contract address is not set")
		return nil, nil
	}

	err := Retry(
		func() error {
			client, err := ethclient.Dial(k.apiUrls.GetEthApiUrl())
			if err != nil {
				k.Logger().Error("rpc error: ethclient dial error", "url", k.apiUrls.GetEthApiUrl(), "err", err)
				k.apiUrls.RotateEthUrl()
				return err
			}
			defer client.Close()

			k.Logger().Debug("creating contracts caller", "address", common.HexToAddress(contractAddr.Address))
			contractCaller, err := contracts.NewContractsCaller(common.HexToAddress(contractAddr.Address), client)
			if err != nil {
				k.Logger().Error("contract caller creation error", "err", err)
				k.apiUrls.RotateEthUrl()
				return err
			}

			k.Logger().Debug("getting validator set", "blkHash", blkHash.Hex(), "blkNum")
			validatorSet, err := contractCaller.GetValidatorSet(&bind.CallOpts{BlockHash: blkHash})
			if err != nil {
				k.Logger().Error("get validator set error", "err", err)
				k.apiUrls.RotateEthUrl()
				return err
			}

			// for some reason symbiotic mixes operator and validator here (seems the same)
			for _, operator := range validatorSet {
				stakes = append(stakes, types.StakedAmountInfo{
					EthereumAddress: operator.Operator.Hex(),
					StakedAmount:    operator.Stake.String(),
				})
			}
			return nil
		},
		RETRIES,
		SLEEP_ON_RETRY*time.Millisecond,
	)
	if err != nil {
		return nil, err
	}

	return stakes, nil
}

func (k *Keeper) SyncNeeded(ctx context.Context) bool {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return (sdkCtx.BlockHeader().Height % SYNC_PERIOD) == 0
}
