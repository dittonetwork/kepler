package keeper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"kepler/x/beacon/types"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
)

// Struct to unmarshal the response from the Beacon Chain API
type Block struct {
	Finalized bool `json:"finalized"`
	Data      struct {
		Message struct {
			Slot string `json:"slot"`
			Body struct {
				ExecutionPayload struct {
					BlockHash   string `json:"block_hash"`
					BlockNumber string `json:"block_number"`
					Timestamp   string `json:"timestamp"`
				} `json:"execution_payload"`
			} `json:"body"`
		} `json:"message"`
	} `json:"data"`
}

type RPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type RPCResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
	Error   *RPCError       `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Validator struct {
	Stake    *big.Int
	ConsAddr [32]byte
}

const (
	SYNC_PERIOD                      = 10
	SLEEP_ON_RETRY                   = 200
	RETRIES                          = 5
	MAINNET_BEACON_GENESIS_TIMESTAMP = 1695902400
	SEPOLIA_BEACON_GENESIS_TIMESTAMP = 1655733600
	SLOTS_IN_EPOCH                   = 32
	SLOT_DURATION                    = 12
	BLOCK_PATH                       = "/eth/v2/beacon/blocks/"
)

func (k Keeper) getFinalizedEpochFirstSlot(ts time.Time) int {
	isMainnet, _ := strconv.ParseBool(os.Getenv("IS_MAINNET"))

	beaconGenesisTimestamp := int64(SEPOLIA_BEACON_GENESIS_TIMESTAMP)
	if isMainnet {
		beaconGenesisTimestamp = MAINNET_BEACON_GENESIS_TIMESTAMP
	}

	slot := (ts.Unix() - beaconGenesisTimestamp) / SLOT_DURATION // get beacon slot
	slot = slot / SLOTS_IN_EPOCH * SLOTS_IN_EPOCH                // first slot of epoch
	// unreliable way to get finalized slot: current - 3
	slot -= 3 * SLOTS_IN_EPOCH
	return int(slot)
}

func (k Keeper) getBlockForSlot(slot int) (Block, error) {
	url := k.beaconApiUrls.GetCurrentUrl() + BLOCK_PATH + strconv.Itoa(slot)

	var block Block
	resp, err := http.Get(url)
	if err != nil {
		k.Logger().Error("rpc error: beacon rpc call error", "url", url, "err", err)
		return block, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		k.Logger().Error(
			"rpc error: beacon rpc call error",
			"url", k.beaconApiUrls.GetCurrentUrl(),
			"err", "no err",
			"status", resp.StatusCode,
		)
	}

	if resp.StatusCode == http.StatusNotFound {
		return block, types.ErrBeaconNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return block, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return block, fmt.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &block)
	if err != nil {
		return block, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return block, nil
}

func (k *Keeper) getFinalizedBlock(ctx context.Context) (types.FinalizedBlockInfo, error) {
	sdkCtx := sdktypes.UnwrapSDKContext(ctx)
	slot := k.getFinalizedEpochFirstSlot(sdkCtx.BlockHeader().Time)

	var block Block
	err := Retry(
		func() error {
			var err error
			for j := 0; j < SLOTS_IN_EPOCH; j++ {
				block, err = k.getBlockForSlot(slot + j)

				// Since some slots could be empty, try the next one within epoch
				if err == nil && !errors.Is(err, types.ErrBeaconNotFound) {
					k.beaconApiUrls.RotateUrl()
					return err
				}
			}
			// almost impossible, no filled slot found within epoch
			return nil
		},
		RETRIES,
		time.Millisecond*SLEEP_ON_RETRY,
	)
	if err != nil {
		return types.FinalizedBlockInfo{}, fmt.Errorf("finding finalized block: %w", err)
	}

	if !block.Finalized {
		// if this fails, need to make sth more smart at `getFinalizedEpochFirstSlot` func
		return types.FinalizedBlockInfo{}, types.ErrNoFinalizedBlock
	}

	slotNum, err := strconv.ParseUint(block.Data.Message.Slot, 10, 64)
	if err != nil {
		return types.FinalizedBlockInfo{}, err
	}
	blkTs, err := strconv.ParseUint(block.Data.Message.Body.ExecutionPayload.Timestamp, 10, 64)
	if err != nil {
		return types.FinalizedBlockInfo{}, err
	}
	blkNum, err := strconv.ParseUint(block.Data.Message.Body.ExecutionPayload.BlockNumber, 10, 64)
	if err != nil {
		return types.FinalizedBlockInfo{}, err
	}

	return types.FinalizedBlockInfo{
		SlotNum:        slotNum,
		BlockTimestamp: blkTs,
		BlockNum:       blkNum,
		BlockHash:      block.Data.Message.Body.ExecutionPayload.BlockHash,
	}, nil
}
