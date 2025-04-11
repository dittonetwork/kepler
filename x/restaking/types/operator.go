package types

import (
	"errors"
	"log"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	EvmAddressLength   = 42
	CompressedPkLength = 33
)

// IsOperatorAddress checks if the given public key matches the operator address (in EVM address format).
func IsOperatorAddress(pk cryptotypes.PubKey, addr string) (bool, error) {
	if pk == nil {
		return false, errors.New("public key is nil")
	}

	if pk.Type() != "secp256k1" {
		return false, errors.New("public key type is not secp256k1")
	}

	if len(addr) != EvmAddressLength {
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
	log.Println(address.Hex(), addr)

	return strings.EqualFold(addr, address.Hex()), nil
}
