# Ditto Symbiotic Module
The Symbiotic module retrieves staking information from a symbiotic middleware contract deployed on the Ethereum blockchain.

## Source Block Selection
To maintain data consistency and prevent reorganization issues, the module first identifies a finalized block. This ensures that the fetched data remains stable for a given block height.
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
Refer to `kepler/x/symbiotic/keeper/beacon.go` for constants values and more implementation details.

## Data Retrieval
The middleware contract stores stake values in USD equivalent. The module fetches this data for all operators and stores it in the keeper's key-value storage.
To query the stored stake information, use:
```sh
keplerd query symbiotic list-staked-amount-info
```
Example output:
```sh
pagination:
  total: "1"
stakedAmountInfo:
- ethereumAddress: 0x254696302Db703E30D6EbCe2D495527A0DD9F6f9
  lastUpdated: "1730953644"
  stakedAmount: "4926617248309304400000"
```

## Configuration
In addition to pre-defined values inside mentioned earlier `kepler/x/symbiotic/keeper/beacon.go` module has few options.

### Contract Address
The module stores the middleware contract address in its key-value storage. To set the address, use:
```sh
keplerd tx symbiotic create-contract-address '0xf7E667dc37bAAF2BFe582CA568d0a8D7E25Aa058' --from alice --fees 20000stake --gas auto
```
This address is stored on-chain and rarely requires updates.
To query the stored contract address, use:
```sh
keplerd query symbiotic show-contract-address
```

### Environment Variables
The module requires the following environment variables for Ethereum node connections:
- `BEACON_API_URLS`: Comma-separated list of beacon node URLs
- `ETH_API_URLS`: Comma-separated list of execution node URLs


### Developer notes
In case middleware contrac ABI is updated, run `make gen-contracts` to regenerate related go code.
