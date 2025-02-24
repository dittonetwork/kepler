# kepler

**kepler** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

### Multi node local testing via docker-compose

1. Build container image via command `docker build docker build -f Dockerfile_multinode -t keplerd_i .`
2. Generate testnets via command `keplerd multi-node`
3. Replace `localhost` to `validatorN` in `persistent_peers` field of `config.toml` of each testnet member

   like `persistent_peers = "<address>:localhost:26656"` -> `persistent_peers = "<address>:validator0:26656"`
4. Run `docker-compose -p kepler up -d` to start the containers