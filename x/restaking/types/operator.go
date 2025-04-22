package types

import (
	"errors"
	"log"

	"github.com/btcsuite/btcd/btcec/v2"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	CompressedPkLength = 33
)

// IsBonded helper function to check if the operator is bonded.
func (v Operator) IsBonded() bool {
	return v.Status == Bonded
}

// IsUnbonding helper function to check if the operator is unbonding.
func (v Operator) IsUnbonding() bool {
	return v.Status == Unbonding
}

func ToKeccakLast20(pk cryptotypes.PubKey) (common.Address, error) {
	if pk == nil {
		return common.Address{}, errors.New("public key is nil")
	}

	if pk.Type() != "secp256k1" {
		return common.Address{}, errors.New("public key type is not secp256k1")
	}

	if len(pk.Bytes()) != CompressedPkLength {
		return common.Address{}, errors.New("invalid public key length")
	}

	pubKey, err := btcec.ParsePubKey(pk.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	return crypto.PubkeyToAddress(*pubKey.ToECDSA()), nil
}
