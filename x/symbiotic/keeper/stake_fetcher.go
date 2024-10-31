package keeper

import (
	"context"
	"errors"
	contracts "kepler/x/symbiotic/abi"
	"kepler/x/symbiotic/types"
	"time"

	"kepler/pkg/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	SLEEP_ON_RETRY = 200
	RETRIES        = 5
)

func (k Keeper) UpdateStakedAmounts(ctx context.Context) error {
	if !k.beaconKeeper.SyncNeeded(ctx) {
		k.Logger().Debug("no need to update stakes yet")
		return nil
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.Logger().Debug("updating stakes amount", "height", sdkCtx.BlockHeight())

	finalizedBlock, exists := k.beaconKeeper.GetFinalizedBlockInfo(ctx)
	if !exists {
		return types.ErrNoFinalizedBlockInBeaconModule
	}

	amountsInfo, err := k.fetchStakedAmountsFromBlock(ctx, ethcommon.HexToHash(finalizedBlock.BlockHash))
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
		stakedAmountInfo.LastUpdateTs = uint64(sdk.UnwrapSDKContext(ctx).BlockTime().Unix())
		k.SetStakedAmountInfo(ctx, stakedAmountInfo)
	}

	k.Logger().Info("successfully updated staked amounts")
	return nil
}

func (k Keeper) fetchStakedAmountsFromBlock(ctx context.Context, blkHash ethcommon.Hash) ([]types.StakedAmountInfo, error) {
	contractAddr, exists := k.GetContractAddress(ctx)
	if !exists {
		return nil, types.ErrContractNotInStorage
	}

	var stakes []types.StakedAmountInfo

	if len(contractAddr.Address) == 0 {
		k.Logger().Warn("contract address is not set")
		return nil, nil
	}

	err := common.Retry(
		func() error {
			client, err := ethclient.Dial(k.rpcUrls.GetCurrentUrl())
			if err != nil {
				k.Logger().Error("rpc error: ethclient dial error", "url", k.rpcUrls.GetCurrentUrl(), "err", err)
				k.rpcUrls.RotateUrl()
				return err
			}
			defer client.Close()

			k.Logger().Debug("creating contracts caller", "address", ethcommon.HexToAddress(contractAddr.Address))
			contractCaller, err := contracts.NewContractsCaller(ethcommon.HexToAddress(contractAddr.Address), client)
			if err != nil {
				k.Logger().Error("contract caller creation error", "err", err)
				k.rpcUrls.RotateUrl()
				return err
			}

			k.Logger().Debug("getting validator set", "blkHash", blkHash.Hex(), "blkNum")
			validatorSet, err := contractCaller.GetValidatorSet(&bind.CallOpts{BlockHash: blkHash})
			if err != nil {
				k.Logger().Error("get validator set error", "err", err)
				k.rpcUrls.RotateUrl()
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
