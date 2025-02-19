package keeper

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/crypto"

	"cosmossdk.io/collections"
	"cosmossdk.io/collections/indexes"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"kepler/x/committee/types"
)

type CommitteeKeeper interface {
	// IsCommitteeExists returns true if the committee exists
	IsCommitteeExists(ctx sdk.Context, committeeID string) (bool, error)

	// GetAuthority returns the module's authority.
	GetAuthority() string

	// SetParams updates the committee module's parameters.
	SetParams(ctx context.Context, params types.Params) error
}

type committeeIndexes struct {
	// Key: ChainID | Value: CommitteeID | Type: MultiKey
	ChainID *indexes.Multi[string, string, types.Committee]

	// Key: ChainID | Value: CommitteeID | Type: Unique
	Active *indexes.Unique[string, string, types.Committee]
}

func (i committeeIndexes) IndexesList() []collections.Index[string, types.Committee] {
	return []collections.Index[string, types.Committee]{
		i.ChainID,
		i.Active,
	}
}

func newCommitteeIndexes(sb *collections.SchemaBuilder) committeeIndexes {
	return committeeIndexes{
		ChainID: indexes.NewMulti(
			sb,
			types.ChainIDStoreKeyPrefix,
			"committee_by_chain_id",
			collections.StringKey,
			collections.StringKey,
			func(_ string, value types.Committee) (string, error) {
				return value.ChainId, nil
			},
		),
		Active: indexes.NewUnique(
			sb,
			types.ActiveCommitteeStoreKeyPrefix,
			"active_committee",
			collections.StringKey,
			collections.StringKey,
			func(_ string, value types.Committee) (string, error) {
				return value.ChainId, nil
			},
		),
	}
}

type Keeper struct {
	cdc    codec.BinaryCodec
	logger log.Logger

	Schema     collections.Schema
	Committees *collections.IndexedMap[string, types.Committee, committeeIndexes]

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:       cdc,
		authority: authority,
		logger:    logger,

		Committees: collections.NewIndexedMap(
			sb,
			types.CommitteeStoreKeyPrefix,
			"committees",
			collections.StringKey,
			codec.CollValue[types.Committee](cdc),
			newCommitteeIndexes(sb),
		),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

func (k Keeper) CanBeSigned(
	goCtx context.Context,
	committeeID string,
	chainID string,
	signatures [][]byte,
	payload []byte,
) (bool, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	addresses := make([]string, 0, len(signatures))
	for _, signature := range signatures {
		address, err := getAddressFromSignature(payload, signature)
		if err != nil {
			return false, err
		}

		addresses = append(addresses, address)
	}

	commID, err := k.Committees.Indexes.Active.MatchExact(ctx, chainID)
	if err != nil {
		return false, err
	}

	activeCommittee, err := k.Committees.Get(ctx, commID)
	if err != nil {
		return false, err
	}

	if activeCommittee.Id != committeeID {
		return false, errors.New("committee ID does not match the active committee")
	}

	membersAddresses := make(map[string]struct{})
	for _, member := range activeCommittee.Members {
		membersAddresses[member.Address] = struct{}{}
	}

	for _, address := range addresses {
		if _, ok := membersAddresses[address]; !ok {
			return false, errors.New("address is not a member of the committee")
		}
	}

	if !isSuperMajority(len(activeCommittee.Members), len(addresses)) {
		return false, errors.New("not enough votes")
	}

	return true, nil
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func getAddressFromSignature(jobMsg []byte, signature []byte) (string, error) {
	hsh := crypto.Keccak256(jobMsg)
	pubKey, err := crypto.SigToPub(hsh, signature)
	if err != nil {
		return "", fmt.Errorf("failed to recover public key from signature: %w", err)
	}

	return crypto.PubkeyToAddress(*pubKey).Hex(), nil
}

func isSuperMajority(totalVotes, votesFor int) bool {
	if votesFor > totalVotes || totalVotes == 0 {
		return false
	}

	//nolint:mnd // just a formula
	requiredVotes := math.Ceil(float64(totalVotes) * 2 / 3)
	return votesFor >= int(requiredVotes)
}
