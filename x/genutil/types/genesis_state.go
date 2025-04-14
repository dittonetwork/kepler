package types

import (
	"encoding/json"
	"fmt"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesisState returns the genutil module's default genesis state.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		GenTxs: []json.RawMessage{},
	}
}

// GenesisStateFromAppGenesis creates the core parameters for genesis initialization
// for the application.
//
// NOTE: The pubkey input is this machines pubkey.
func GenesisStateFromAppGenesis(genesis *AppGenesis) (map[string]json.RawMessage, error) {
	var genesisState map[string]json.RawMessage

	if err := json.Unmarshal(genesis.AppState, &genesisState); err != nil {
		return genesisState, err
	}
	return genesisState, nil
}

// GenesisStateFromGenFile creates the core parameters for genesis initialization
// for the application.
//
// NOTE: The pubkey input is this machines pubkey.
func GenesisStateFromGenFile(genFile string) (map[string]json.RawMessage, *AppGenesis, error) {
	if _, err := os.Stat(genFile); os.IsNotExist(err) {
		return nil, nil, fmt.Errorf("%s does not exist, run `init` first", genFile)
	}

	genesis, err := AppGenesisFromFile(genFile)
	if err != nil {
		return nil, nil, err
	}

	genesisState, err := GenesisStateFromAppGenesis(genesis)
	return genesisState, genesis, err
}

// ValidateGenesis validates GenTx transactions.
func ValidateGenesis(genesisState *GenesisState, txJSONDecoder sdk.TxDecoder, validator MessageValidator) error {
	for _, genTx := range genesisState.GenTxs {
		_, err := ValidateAndGetGenTx(genTx, txJSONDecoder, validator)
		if err != nil {
			return err
		}
	}
	return nil
}

type MessageValidator func([]sdk.Msg) error

func DefaultMessageValidator(msgs []sdk.Msg) error {
	if len(msgs) != 1 {
		return fmt.Errorf("unexpected number of GenTx messages; got: %d, expected: 1", len(msgs))
	}

	if m, ok := msgs[0].(sdk.HasValidateBasic); ok {
		if err := m.ValidateBasic(); err != nil {
			return fmt.Errorf("invalid GenTx '%s': %w", msgs[0], err)
		}
	}

	return nil
}

// ValidateAndGetGenTx validates the genesis transaction and returns GenTx if valid
// it cannot verify the signature as it is stateless validation.
func ValidateAndGetGenTx(
	genTx json.RawMessage,
	txJSONDecoder sdk.TxDecoder,
	validator MessageValidator,
) (sdk.Tx, error) {
	tx, err := txJSONDecoder(genTx)
	if err != nil {
		return tx, fmt.Errorf("failed to decode gentx: %s, error: %w", genTx, err)
	}

	return tx, validator(tx.GetMsgs())
}
