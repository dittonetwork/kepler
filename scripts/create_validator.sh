#!/bin/sh

NODE_HOME=$1
NODE_DIR=$2
MONIKER=$3
account=$4

# Create validator_info.json
cat > $NODE_DIR/validator_info.json << EOF
{
  "pubkey": $(keplerd comet show-validator --home $NODE_HOME),
  "amount": "10000000000stake",
  "moniker": "$MONIKER",
  "identity": "",
  "website": "",
  "security": "",
  "details": "",
  "commission-rate": "0.1",
  "commission-max-rate": "0.2",
  "commission-max-change-rate": "0.01",
  "min-self-delegation": "1"
}
EOF

# Create validator
keplerd tx staking create-validator $NODE_DIR/validator_info.json \
  --from $account \
  --home $NODE_HOME \
  --fees 100000stake \
  --chain-id kepler \
  --broadcast-mode sync \
  --keyring-backend=test \
  --yes