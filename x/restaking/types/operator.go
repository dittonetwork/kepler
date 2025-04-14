package types

import (
	"errors"
	"log"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	Keccak20AddressLength = 42
	CompressedPkLength    = 33
)

// IsBonded helper function to check if the operator is bonded.
func (v Operator) IsBonded() bool {
	return v.Status == Bonded
}

// IsUnbonding helper function to check if the operator is unbonding.
func (v Operator) IsUnbonding() bool {
	return v.Status == Unbonding
}

// IsOperatorAddress checks if the given public key matches the operator address (in EVM address format).
func IsOperatorAddress(pk cryptotypes.PubKey, addr string) (bool, error) {
	if pk == nil {
		return false, errors.New("public key is nil")
	}

	if pk.Type() != "secp256k1" {
		return false, errors.New("public key type is not secp256k1")
	}

	if len(addr) != Keccak20AddressLength {
		return false, errors.New("invalid operator address length")
	}

	if len(pk.Bytes()) != CompressedPkLength {
		return false, errors.New("invalid public key length")
	}

	// Convert to secp256k1 public key
	pubKey, err := btcec.ParsePubKey(pk.Bytes())
	if err != nil {
		return false, err
	}
	address := crypto.PubkeyToAddress(*pubKey.ToECDSA())

	return strings.EqualFold(addr, address.Hex()), nil
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
