package keeper_test

import (
	"crypto/ecdsa"
	"fmt"
	"testing"

	"github.com/dittonetwork/kepler/testutil/keeper"

	"github.com/dittonetwork/kepler/x/committee/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestKeeper_CanBeSigned_Success(t *testing.T) {
	k, ctx := keeper.CommitteeKeeper(t)

	chainID := "chainID"
	committeeID := "committeeID"

	privKeys := make([]*ecdsa.PrivateKey, 0, 5)
	addresses := make([]string, 0, 5)
	for i := 0; i < 5; i++ {
		key, err := crypto.GenerateKey()
		require.NoError(t, err)
		privKeys = append(privKeys, key)
		// Get the hex representation of the public address.
		addr := crypto.PubkeyToAddress(key.PublicKey).Hex()
		addresses = append(addresses, addr)
	}

	committee := types.Committee{
		Id:           committeeID,
		ChainId:      chainID,
		EpochCounter: 0,
		Active:       true,
		Seed:         []byte("seed"),
		Members: []*types.Member{
			{Address: addresses[0], Power: 1},
			{Address: addresses[1], Power: 1},
			{Address: addresses[2], Power: 1},
			{Address: addresses[3], Power: 1},
			{Address: addresses[4], Power: 1},
		},
	}

	err := k.Committees.Set(ctx, committeeID, committee)
	require.NoError(t, err)

	prefixedMsg := getPrefixedMessage()
	msgHash := crypto.Keccak256(prefixedMsg)

	signatures := make([][]byte, 0, 5)
	for i := 0; i < 5; i++ {
		sig, err := crypto.Sign(msgHash, privKeys[i])
		require.NoError(t, err)
		signatures = append(signatures, sig)
	}

	res, err := k.CanBeSigned(ctx, committeeID, chainID, signatures, prefixedMsg)
	require.NoError(t, err)
	require.True(t, res, "expected CanBeSigned to return true with valid signatures")
}

func TestKeeper_CanBeSigned_NoCommitteeForChain(t *testing.T) {
	k, ctx := keeper.CommitteeKeeper(t)

	commChainID := "commChainID"
	reqChainID := "ChainID"
	committeeID := "committeeID"

	privKeys := make([]*ecdsa.PrivateKey, 0, 5)
	addresses := make([]string, 0, 5)
	for i := 0; i < 5; i++ {
		key, err := crypto.GenerateKey()
		require.NoError(t, err)
		privKeys = append(privKeys, key)
		// Get the hex representation of the public address.
		addr := crypto.PubkeyToAddress(key.PublicKey).Hex()
		addresses = append(addresses, addr)
	}

	committee := types.Committee{
		Id:           committeeID,
		ChainId:      commChainID,
		EpochCounter: 0,
		Active:       true,
		Seed:         []byte("seed"),
		Members: []*types.Member{
			{Address: addresses[0], Power: 1},
			{Address: addresses[1], Power: 1},
			{Address: addresses[2], Power: 1},
			{Address: addresses[3], Power: 1},
			{Address: addresses[4], Power: 1},
		},
	}

	err := k.Committees.Set(ctx, committeeID, committee)
	require.NoError(t, err)

	prefixedMsg := getPrefixedMessage()
	msgHash := crypto.Keccak256(prefixedMsg)

	signatures := make([][]byte, 0, 5)
	for i := 0; i < 5; i++ {
		sig, err := crypto.Sign(msgHash, privKeys[i])
		require.NoError(t, err)
		signatures = append(signatures, sig)
	}

	_, err = k.CanBeSigned(ctx, committeeID, reqChainID, signatures, prefixedMsg)
	require.Error(t, err)
}

func TestKeeper_CanBeSigned_CommitteeIsNoLongerActive(t *testing.T) {
	k, ctx := keeper.CommitteeKeeper(t)
	chainID := "ChainID"
	committeeID := "committeeID"

	privKeys := make([]*ecdsa.PrivateKey, 0, 5)
	addresses := make([]string, 0, 5)
	for i := 0; i < 5; i++ {
		key, err := crypto.GenerateKey()
		require.NoError(t, err)
		privKeys = append(privKeys, key)
		// Get the hex representation of the public address.
		addr := crypto.PubkeyToAddress(key.PublicKey).Hex()
		addresses = append(addresses, addr)
	}

	committee := types.Committee{
		Id:           committeeID,
		ChainId:      chainID,
		EpochCounter: 0,
		Active:       true,
		Seed:         []byte("seed"),
		Members: []*types.Member{
			{Address: addresses[0], Power: 1},
			{Address: addresses[1], Power: 1},
			{Address: addresses[2], Power: 1},
			{Address: addresses[3], Power: 1},
			{Address: addresses[4], Power: 1},
		},
	}

	err := k.Committees.Set(ctx, committeeID, committee)
	require.NoError(t, err)

	prefixedMsg := getPrefixedMessage()
	msgHash := crypto.Keccak256(prefixedMsg)

	signatures := make([][]byte, 0, 5)
	for i := 0; i < 5; i++ {
		sig, err := crypto.Sign(msgHash, privKeys[i])
		require.NoError(t, err)
		signatures = append(signatures, sig)
	}

	wrongCommitteeID := "wrongCommitteeID"
	_, err = k.CanBeSigned(ctx, wrongCommitteeID, chainID, signatures, prefixedMsg)
	require.Error(t, err)
}

func TestKeeper_CanBeSigned_SignerNotInCommittee(t *testing.T) {
	k, ctx := keeper.CommitteeKeeper(t)

	chainID := "chainID"
	committeeID := "committeeID"

	privKeys := make([]*ecdsa.PrivateKey, 0, 5)
	addresses := make([]string, 0, 5)
	for i := 0; i < 5; i++ {
		key, err := crypto.GenerateKey()
		require.NoError(t, err)
		privKeys = append(privKeys, key)
		// Get the hex representation of the public address.
		addr := crypto.PubkeyToAddress(key.PublicKey).Hex()
		addresses = append(addresses, addr)
	}

	committee := types.Committee{
		Id:           committeeID,
		ChainId:      chainID,
		EpochCounter: 0,
		Active:       true,
		Seed:         []byte("seed"),
		Members: []*types.Member{
			{Address: addresses[0], Power: 1},
			{Address: addresses[1], Power: 1},
			{Address: addresses[2], Power: 1},
			{Address: addresses[3], Power: 1},
			{Address: addresses[4], Power: 1},
		},
	}

	err := k.Committees.Set(ctx, committeeID, committee)
	require.NoError(t, err)

	prefixedMsg := getPrefixedMessage()
	msgHash := crypto.Keccak256(prefixedMsg)

	signatures := make([][]byte, 0, 5)
	for i := 0; i < 3; i++ {
		sig, err := crypto.Sign(msgHash, privKeys[i])
		require.NoError(t, err)
		signatures = append(signatures, sig)
	}

	// Add a signature from a non-member
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	sig, err := crypto.Sign(msgHash, key)
	require.NoError(t, err)
	signatures = append(signatures, sig)

	_, err = k.CanBeSigned(ctx, committeeID, chainID, signatures, prefixedMsg)
	require.Error(t, err)
}

func TestKeeper_CanBeSigned_NoSuperMajority(t *testing.T) {
	k, ctx := keeper.CommitteeKeeper(t)

	chainID := "chainID"
	committeeID := "committeeID"

	privKeys := make([]*ecdsa.PrivateKey, 0, 5)
	addresses := make([]string, 0, 5)
	for i := 0; i < 5; i++ {
		key, err := crypto.GenerateKey()
		require.NoError(t, err)
		privKeys = append(privKeys, key)
		// Get the hex representation of the public address.
		addr := crypto.PubkeyToAddress(key.PublicKey).Hex()
		addresses = append(addresses, addr)
	}

	committee := types.Committee{
		Id:           committeeID,
		ChainId:      chainID,
		EpochCounter: 0,
		Active:       true,
		Seed:         []byte("seed"),
		Members: []*types.Member{
			{Address: addresses[0], Power: 1},
			{Address: addresses[1], Power: 1},
			{Address: addresses[2], Power: 1},
			{Address: addresses[3], Power: 1},
			{Address: addresses[4], Power: 1},
		},
	}

	err := k.Committees.Set(ctx, committeeID, committee)
	require.NoError(t, err)

	prefixedMsg := getPrefixedMessage()
	msgHash := crypto.Keccak256(prefixedMsg)

	signatures := make([][]byte, 0, 5)
	for i := 0; i < 2; i++ {
		sig, err := crypto.Sign(msgHash, privKeys[i])
		require.NoError(t, err)
		signatures = append(signatures, sig)
	}

	_, err = k.CanBeSigned(ctx, committeeID, chainID, signatures, prefixedMsg)
	require.Error(t, err)
}

func TestKeeper_CanBeSigned_InvalidAddressFromSig(t *testing.T) {
	k, ctx := keeper.CommitteeKeeper(t)

	chainID := "chainID"
	committeeID := "committeeID"

	privKeys := make([]*ecdsa.PrivateKey, 0, 5)
	addresses := make([]string, 0, 5)
	for i := 0; i < 5; i++ {
		key, err := crypto.GenerateKey()
		require.NoError(t, err)
		privKeys = append(privKeys, key)
		// Get the hex representation of the public address.
		addr := crypto.PubkeyToAddress(key.PublicKey).Hex()
		addresses = append(addresses, addr)
	}

	committee := types.Committee{
		Id:           committeeID,
		ChainId:      chainID,
		EpochCounter: 0,
		Active:       true,
		Seed:         []byte("seed"),
		Members: []*types.Member{
			{Address: addresses[0], Power: 1},
			{Address: addresses[1], Power: 1},
			{Address: addresses[2], Power: 1},
			{Address: addresses[3], Power: 1},
			{Address: addresses[4], Power: 1},
		},
	}

	err := k.Committees.Set(ctx, committeeID, committee)
	require.NoError(t, err)

	prefixedMsg := getPrefixedMessage()
	msgHash := crypto.Keccak256(prefixedMsg)

	signatures := make([][]byte, 0, 5)
	for i := 0; i < 5; i++ {
		sig, err := crypto.Sign(msgHash, privKeys[i])
		require.NoError(t, err)
		if i == 2 {
			// Modify the signature to make it invalid
			sig[0] = sig[0] + 1
		}
		signatures = append(signatures, sig)
	}

	_, err = k.CanBeSigned(ctx, committeeID, chainID, signatures, prefixedMsg)
	require.Error(t, err)
}

func getPrefixedMessage() []byte {
	jobPayload := []byte("sample job payload")
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(jobPayload))
	prefixedMsg := []byte(prefix)
	prefixedMsg = append(prefixedMsg, jobPayload...)

	return prefixedMsg
}
