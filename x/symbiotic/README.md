# Ditto Symbiotic Module
The Symbiotic module retrieves staking information from a symbiotic middleware contract deployed on the Ethereum blockchain.

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
  lastUpdatedTs: "1730953644"
  stakedAmount: "4926617248309304400000"
```

## Configuration
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
- `ETH_RPC_URLS`: Comma-separated list of execution node URLs


### Developer notes
In case middleware contrac ABI is updated, run `make gen-contracts` to regenerate related go code.
