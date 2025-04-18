package sample

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	secp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Account struct {
	Name       string
	Address    sdk.AccAddress
	ValAddress sdk.ValAddress
	PubKey     cryptotypes.PubKey
}

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

func GetAccount(name string) Account {
	pvKey := secp.GenPrivKey()
	pubKey := pvKey.PubKey()
	pkAddr := pubKey.Address()

	return Account{
		Address:    sdk.AccAddress(pkAddr),
		ValAddress: sdk.ValAddress(pkAddr),
		PubKey:     pubKey,
		Name:       name,
	}
}
