#!/bin/sh

NODE2_HOME=$1
NODE2_DIR=$2

# Create validator_info.json
cat > $NODE2_DIR/validator_info.json << EOF
{
  "pubkey": $(keplerd comet show-validator --home $NODE2_HOME),
  "amount": "10000000000stake",
  "moniker": "node2",
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
keplerd tx staking create-validator $NODE2_DIR/validator_info.json \
  --from node2-account \
  --home $NODE2_HOME \
  --fees 100000stake \
  --chain-id kepler \
  --broadcast-mode sync \
  --keyring-backend=test \
  --yes