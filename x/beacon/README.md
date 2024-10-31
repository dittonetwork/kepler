# Ditto Beacon Module
The Beacon module retrieves finalized block info from Etherium Beacon chain. To maintain data consistency and prevent reorganization issues we fetch only finalized block at the expense of experiencing significant update delays. This ensures that the fetched data remains stable for a given block height.

## Source Block Selection
The process of finding the last finalized block follows these steps:
1. Calculate the first slot of the current epoch
1. Move back 3 epochs to ensure we're working with finalized data
1. Iterate through all slots within the identified epoch until a non-empty slot is found
1. If all slots in the epoch are empty, raise an error
```go
slot := (ts.Unix() - BEACON_GENESIS_TIMESTAMP) / SLOT_DURATION // get beacon slot
slot = slot / SLOTS_IN_EPOCH * SLOTS_IN_EPOCH                  // first slot of epoch
slot -= 3 * SLOTS_IN_EPOCH
...
for j := 0; j < SLOTS_IN_EPOCH; j++ {
    block, err = k.getBlockForSlot(slot + j)
    if err != nil {
        // retry logic
    }
}
if block == nil {
    // raise err
}
```
Refer to [api.go](keeper/api.go) for constants' values and more implementation details.

## Data Retrieval
The middleware contract stores stake values in USD equivalent. The module fetches this data for all operators and stores it in the keeper's key-value storage.
To query the stored stake information, use:
```sh
keplerd query beacon show-finalized-block-info
```
Example output:
```sh
FinalizedBlockInfo:
  blockHash: 0x67aafbd324259b053e27f4af3b1336795e4f3e913fadd874f5bb8ccfc89958b5
  blockNum: "7148679"
  blockTimestamp: "1732525152"
  slotNum: "6399296"
```

## Configuration
### Environment Variables
The module requires the following environment variables for Ethereum node connections:
- `BEACON_API_URLS`: Comma-separated list of beacon node URLs
