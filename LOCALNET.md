## 1. Prerequisites
``
Golang 1.23
``

``
Ignite CLI (curl https://get.ignite.com/cli! | bash)
``

``
/home/<your-username>/go/bin and /usr/local/go/bin added to PATH (export PATH=$PATH:$HOME/go/bin:/usr/local/go/bin)
``

## 2. Init local testnet state

``
make init
``

It command cleans old state, create a new one, add 2 genesis accounts with following mnemonics:

``
again glare leg choice input april tone brush goose seek forum dinosaur link speed digital question ticket caught strike quote release crowd fork deposit
``

``
heavy notice slice boil dune monitor pet slight denial sea train notable section boring fancy evolve kangaroo lazy pride cluster gaze any hotel chimney
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