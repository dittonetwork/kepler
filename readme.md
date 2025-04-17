# kepler

**kepler** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

1. **Build chain:**
```bash
$ ignite chain build
```

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

> [!IMPORTANT]
> Folder '<root>/.localnet/{moniker}_priv_validator_key.json' is created for more convenient local testing.
> So that a new consensus key is not generated each time you reset the state of keplerd,
> this one will be used by the validator.
> Under no circumstances should it be used anywhere else except for local development.

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project.


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

### Multi node local testing via docker-compose

1. Build container image via command `docker build -f Dockerfile_multinode -t keplerd_i .`
2. Generate testnets via command `keplerd multi-node`
3. Replace `localhost` to `validatorN` in `persistent_peers` field of `config.toml` of each testnet member

   like `persistent_peers = "<address>:localhost:26656"` -> `persistent_peers = "<address>:validator0:26656"`
4. Run `docker-compose -p kepler up -d` to start the containers


## Production setup

Initialize your blockchain with a moniker (name) of your choice.
This will create the necessary configuration files and directories at `~/.keplerd` by default.
```bash
$ keplerd init [moniker]
```

**Create a new account with initial tokens**
```bash
$ keplerd genesis add-genesis-account [address] [amount]
```

**Add genesis pending validators**
```bash
$ keplerd genesis bulk-add-genesis-operators ./operators.json
```

Example `operators.json` file:
```json
{
	"operators": [
		{
			"address": "0x910cB6A0937ECeBA1EDF4F505F1b86D3234a4Fe9",
			"consensus_pubkey": {
				"@type": "/cosmos.crypto.ed25519.PubKey",
				"key": "dcHikhOeHpXzLzX+IFlz6HuNs5ILr3v/OG5NfGQuvKE="
			},
			"is_emergency": true,
			"status": "BOND_STATUS_BONDED",
			"voting_power": "10000000",
			"protocol": "PROTOCOL_DITTO"
		}
	]
}
```


**Create genesis transactions:**
```bash
$ keplerd genesis gentx [address] [amount]
```

> [!IMPORTANT]
> The amount can be anything, as it does not play a role in the logic of forming the transaction.
> This interface is needed to maintain backward compatibility during local development.
> In the future, we might create a separate command for production runs without unnecessary parameters.

Example:
```bash
$ keplerd genesis gentx alice 0ditto
```

Collect the genesis transactions:
```bash
$ keplerd genesis collect-gentxs
```

Now you can start your blockchain with:
```bash
$ keplerd start
```
