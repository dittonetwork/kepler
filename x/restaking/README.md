# Restaking Module

The Restaking module is a Kepler module that enables validator bonding and management functionality within the Kepler blockchain. It integrates with major restaking protocols like Symbiotic and EigenLayer to allow their operators to become validators on the Kepler network.
This module requires bridging operators from L1 to its network.

## Overview

This module provides functionality for:
- Bonding validators to the network
- Managing validator descriptions and metadata
- Handling validator-related state transitions
- Integration with restaking protocols (Symbiotic, EigenLayer)
- Converting restaking protocol operators into Kepler validators

## Module Structure

```
x/restaking/
├── client/         # CLI and REST client implementations
├── keeper/         # Module state management
├── module/         # Module configuration and setup
├── repository/     # Data access layer
├── simulation/     # Simulation testing utilities
├── testutil/       # Testing utilities
└── types/          # Type definitions and message handlers
```

## Becoming Validator Process Flow

### Step 1: Node Setup
- Set up the Kepler node
- Start the node and ensure it's properly configured

### Step 2: L1 Bridging
- Bridging L1 operators to Kepler
- Wait for the operator to be bridged and appear as a pending validator on Kepler
You can check this by query `keplerd query restaking pending-operators`

### Step 3: Operator Bonding
- Submit `MsgBondValidator` transaction to complete the operator bonding process
- After successful bonding, the operator becomes an active validator on Kepler

## Key Components

### Messages

The module implements the following message types:
- `MsgBondValidator`: Used to bond a new validator

### State Management

The module maintains state for:
- Validator information
- Bonding status
- Validator descriptions
- Restaking protocol operator mappings

### L1 Operator Bridging

The Restaking module requires a bridge mechanism to connect L1 operators to the Kepler network. In the case of Ditto, this is achieved through:

**Tess Decentralized Execution Layer**
- Provides a secure and decentralized way to bridge L1 operators
- Enables cross-chain operator verification and validation
- Ensures trustless operation between L1 and Kepler network

## Running Node from Genesis

### Prerequisites
- Make sure you have the Kepler binary (`keplerd`) installed
- Ensure you have sufficient permissions to create and modify files in the working directory
- You need to have `x/genutil` module 

### Setup Steps

0. Create a key in your keyring:
```bash
keplerd keys add [key_name]
```
You can specify the keyring backend using the `--keyring-backend` flag.

1. Initialize the blockchain with your moniker:
```bash
keplerd init [moniker] --chain-id kepler
```

2. Add a genesis account with initial tokens:
```bash
keplerd genesis add-genesis-account [key_name] [amount]
```

3. Add genesis pending validators using operators.json:
```bash
keplerd genesis bulk-add-genesis-operators ./operators.json
```

Example `operators.json`:
```json
{
    "operators": [
        {
            "address": "0xOperator_Address_From_L1",
            "consensus_pubkey": {
                "@type": "/cosmos.crypto.ed25519.PubKey",
                "key": "base64 bytes" // from ~/.kepler/config/priv_validator_key.json
            },
            "is_emergency": false,
            "status": "BOND_STATUS_BONDED",
            "voting_power": "100",
            "protocol": "PROTOCOL_DITTO" // restaking protocol name or something else (depends on L1 contracts)
        }
    ]
}
```

4. Create a genesis transaction:
```bash
keplerd genesis gentx [address] [amount] --chain-id kepler
```

5. Collect genesis transactions:
```bash
keplerd genesis collect-gentxs
```

6. Start the blockchain:
```bash
keplerd start
```

### Important Notes
- The `operators.json` file must be properly formatted with valid operator addresses and consensus public keys
- Make sure to use the correct chain ID when creating genesis transactions
- The voting power and protocol fields in `operators.json` should match your L1 contract configuration
