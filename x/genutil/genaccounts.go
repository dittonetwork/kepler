package genutil

import (
	"encoding/json"
	"fmt"

	"cosmossdk.io/core/address"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutiltypes "github.com/dittonetwork/kepler/x/genutil/types"
)

// AddGenesisAccount adds a genesis account to the genesis state.
// Where `cdc` is client codec, `genesisFileUrl` is the path/url of current genesis file,
// `accAddr` is the address to be added to the genesis state, `amountStr` is the list of initial coins
// to be added for the account, `appendAcct` updates the account if already exists.
// `moduleName` is the module name for which the account is being created
// and coins to be appended to the account already in the genesis.json file.
func AddGenesisAccount(
	cdc codec.Codec,
	accAddr sdk.AccAddress,
	appendAcct bool,
	genesisFileURL, amountStr,
	moduleName string,
) error {
	coins, err := sdk.ParseCoinsNormalized(amountStr)
	if err != nil {
		return fmt.Errorf("failed to parse coins: %w", err)
	}

	// create concrete account type based on input parameters
	var genAccount authtypes.GenesisAccount

	balances := banktypes.Balance{Address: accAddr.String(), Coins: coins.Sort()}
	baseAccount := authtypes.NewBaseAccount(accAddr, nil, 0, 0)

	if moduleName != "" {
		genAccount = authtypes.NewEmptyModuleAccount(moduleName, authtypes.Burner, authtypes.Minter)
	} else {
		genAccount = baseAccount
	}

	if err = genAccount.Validate(); err != nil {
		return fmt.Errorf("failed to validate new genesis account: %w", err)
	}

	appState, appGenesis, err := genutiltypes.GenesisStateFromGenFile(genesisFileURL)
	if err != nil {
		return fmt.Errorf("failed to unmarshal genesis state: %w", err)
	}

	authGenState := authtypes.GetGenesisStateFromAppState(cdc, appState)

	accs, err := authtypes.UnpackAccounts(authGenState.Accounts)
	if err != nil {
		return fmt.Errorf("failed to get accounts from any: %w", err)
	}

	bankGenState := banktypes.GetGenesisStateFromAppState(cdc, appState)
	//nolint: nestif // no matter
	if accs.Contains(accAddr) {
		if !appendAcct {
			return fmt.Errorf(" Account %s already exists\nUse `append` flag to append account at existing address", accAddr)
		}

		genesisB := banktypes.GetGenesisStateFromAppState(cdc, appState)
		for idx, acc := range genesisB.Balances {
			if acc.Address != accAddr.String() {
				continue
			}

			updatedCoins := acc.Coins.Add(coins...)
			bankGenState.Balances[idx] = banktypes.Balance{Address: accAddr.String(), Coins: updatedCoins.Sort()}
			break
		}
	} else {
		// Add the new account to the set of genesis accounts and sanitize the accounts afterwards.
		accs = append(accs, genAccount)
		accs = authtypes.SanitizeGenesisAccounts(accs)

		var genAccs []*types.Any
		genAccs, err = authtypes.PackAccounts(accs)
		if err != nil {
			return fmt.Errorf("failed to convert accounts into any's: %w", err)
		}
		authGenState.Accounts = genAccs

		var authGenStateBz []byte
		authGenStateBz, err = cdc.MarshalJSON(&authGenState)
		if err != nil {
			return fmt.Errorf("failed to marshal auth genesis state: %w", err)
		}
		appState[authtypes.ModuleName] = authGenStateBz

		bankGenState.Balances = append(bankGenState.Balances, balances)
	}

	bankGenState.Balances = banktypes.SanitizeGenesisBalances(bankGenState.Balances)

	bankGenState.Supply = bankGenState.Supply.Add(balances.Coins...)

	bankGenStateBz, err := cdc.MarshalJSON(bankGenState)
	if err != nil {
		return fmt.Errorf("failed to marshal bank genesis state: %w", err)
	}
	appState[banktypes.ModuleName] = bankGenStateBz

	appStateJSON, err := json.Marshal(appState)
	if err != nil {
		return fmt.Errorf("failed to marshal application genesis state: %w", err)
	}

	appGenesis.AppState = appStateJSON
	return ExportGenesisFile(appGenesis, genesisFileURL)
}

type GenesisAccount struct {
	// Base
	Address string    `json:"address"`
	Coins   sdk.Coins `json:"coins"`

	// Module
	ModuleName string `json:"module_name,omitempty"`
}

// AddGenesisAccounts adds genesis accounts to the genesis state.
// Where `cdc` is the client codec, `accounts` are the genesis accounts to add,
// `appendAcct` updates the account if already exists, and `genesisFileURL` is the path/url of the current genesis file.
//
//nolint:gocognit,funlen // no matter
func AddGenesisAccounts(
	cdc codec.Codec,
	ac address.Codec,
	accounts []GenesisAccount,
	appendAcct bool,
	genesisFileURL string,
) error {
	appState, appGenesis, err := genutiltypes.GenesisStateFromGenFile(genesisFileURL)
	if err != nil {
		return fmt.Errorf("failed to unmarshal genesis state: %w", err)
	}

	authGenState := authtypes.GetGenesisStateFromAppState(cdc, appState)
	bankGenState := banktypes.GetGenesisStateFromAppState(cdc, appState)

	accs, err := authtypes.UnpackAccounts(authGenState.Accounts)
	if err != nil {
		return fmt.Errorf("failed to get accounts from any: %w", err)
	}

	newSupplyCoinsCache := sdk.NewCoins()
	balanceCache := make(map[string]banktypes.Balance)
	for _, acc := range accs {
		for _, balance := range bankGenState.GetBalances() {
			if balance.Address == acc.GetAddress().String() {
				balanceCache[acc.GetAddress().String()] = balance
			}
		}
	}

	for _, acc := range accounts {
		addr := acc.Address
		coins := acc.Coins

		var accAddr []byte
		accAddr, err = ac.StringToBytes(addr)
		if err != nil {
			return fmt.Errorf("failed to parse account address %s: %w", addr, err)
		}

		// create concrete account type based on input parameters
		var genAccount authtypes.GenesisAccount

		balances := banktypes.Balance{Address: addr, Coins: coins.Sort()}
		baseAccount := authtypes.NewBaseAccount(accAddr, nil, 0, 0)

		if acc.ModuleName != "" {
			genAccount = authtypes.NewEmptyModuleAccount(acc.ModuleName, authtypes.Burner, authtypes.Minter)
		} else {
			genAccount = baseAccount
		}

		if err = genAccount.Validate(); err != nil {
			return fmt.Errorf("failed to validate new genesis account: %w", err)
		}

		if _, ok := balanceCache[addr]; ok {
			if !appendAcct {
				return fmt.Errorf(" Account %s already exists\nUse `append` flag to append account at existing address", accAddr)
			}

			for idx, account := range bankGenState.Balances {
				if account.Address != addr {
					continue
				}

				updatedCoins := account.Coins.Add(coins...)
				bankGenState.Balances[idx] = banktypes.Balance{Address: addr, Coins: updatedCoins.Sort()}
				break
			}
		} else {
			accs = append(accs, genAccount)
			bankGenState.Balances = append(bankGenState.Balances, balances)
		}

		newSupplyCoinsCache = newSupplyCoinsCache.Add(coins...)
	}

	accs = authtypes.SanitizeGenesisAccounts(accs)

	authGenState.Accounts, err = authtypes.PackAccounts(accs)
	if err != nil {
		return fmt.Errorf("failed to convert accounts into any's: %w", err)
	}

	appState[authtypes.ModuleName], err = cdc.MarshalJSON(&authGenState)
	if err != nil {
		return fmt.Errorf("failed to marshal auth genesis state: %w", err)
	}

	bankGenState.Balances = banktypes.SanitizeGenesisBalances(bankGenState.Balances)
	bankGenState.Supply = bankGenState.Supply.Add(newSupplyCoinsCache...)

	appState[banktypes.ModuleName], err = cdc.MarshalJSON(bankGenState)
	if err != nil {
		return fmt.Errorf("failed to marshal bank genesis state: %w", err)
	}

	appStateJSON, err := json.Marshal(appState)
	if err != nil {
		return fmt.Errorf("failed to marshal application genesis state: %w", err)
	}

	appGenesis.AppState = appStateJSON
	return ExportGenesisFile(appGenesis, genesisFileURL)
}
