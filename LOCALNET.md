## 1. Prerequisites
``
Golang 1.23
``

``
Docker
``

``
Ignite CLI (curl https://get.ignite.com/cli! | bash)
``

``
/home/<your-username>/go/bin and /usr/local/go/bin added to PATH (export PATH=$PATH:$HOME/go/bin:/usr/local/go/bin)
``

## 2. Init localnet state

``
make init
``

It command cleans old state, create a new one, add 5 genesis accounts with following mnemonics:

``
again glare leg choice input april tone brush goose seek forum dinosaur link speed digital question ticket caught strike quote release crowd fork deposit
``

``
heavy notice slice boil dune monitor pet slight denial sea train notable section boring fancy evolve kangaroo lazy pride cluster gaze any hotel chimney
``

``
crime salute atom that obey gaze among arrow decide bicycle robust glow diesel possible hill system three ethics share violin recipe oven dice height
``

``
follow appear body worry number laundry crucial fiction swing sure image poem alter hole motor chef traffic enact near arrange imitate shoulder saddle hybrid
``

``
cancel welcome evolve effort luggage long foam teach cage betray library damp pink artist dilemma foster mean blast this moon charge lobster swarm ordinary
``

You can use it for any txs

## 3. Start one-validator-node localnet

``
make start-node1
``

## 3.1. Add another node to localnet
In separated terminal:

``
make start-node2
``

## 3.2. Turn the second node into validator-node
In separated terminal:

``
make create-validator
``

Now you could check list of validators, there are should be 2:

``
make check-validators
``

## 4. Localnet with 5 validators

``
make localnet
``

build takes time (~5 minutes), as a result localnet in docker will be started. Now there are 1 validator and 4 nodes:

| Node ID       | P2P Port | Tendermint RPC Port | gRPC   | REST   |
|---------------|----------|---------------------|--------|--------|
| `keplernode1` | `26656`  | `26657`             | `9090` | `1317` |
| `keplernode2` | `26666`  | `26667`             | `9091` | `1318` |
| `keplernode3` | `26676`  | `26677`             | `9092` | `1319` |
| `keplernode4` | `26686`  | `26687`             | `9093` | `1320` |
| `keplernode5` | `26696`  | `26697`             | `9094` | `1321` |

For validators' initialization:

``
make create-validators-localnet
``

wait 5-10 seconds for txs applying and check validators' list, there should be 5:

``
make check-validators
``

## 5. Separate postgres without docker nodes

`make setup-postgres-indexer`

`docker compose up postgres -d`

## 6. Localnet with postgres

`make localnet-postgres`
