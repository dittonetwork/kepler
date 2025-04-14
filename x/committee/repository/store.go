package repository

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/dittonetwork/kepler/x/committee/types"
)

type CommitteeRepository struct {
	Schema     collections.Schema
	Committees *collections.IndexedMap[uint32, types.Committee, Idx]
	LastEpoch  collections.Item[uint32]
}

var _ types.Repository = &CommitteeRepository{}

func NewCommitteeRepository(storeService store.KVStoreService, cdc codec.BinaryCodec) *CommitteeRepository {
	sb := collections.NewSchemaBuilder(storeService)

	cr := &CommitteeRepository{
		Committees: collections.NewIndexedMap(
			sb,
			types.CommitteesStoreKeyPrefix,
			"committees",
			collections.Uint32Key,
			codec.CollValue[types.Committee](cdc),
			NewIndexes(sb),
		),
		LastEpoch: collections.NewItem(
			sb,
			types.LatestEpochStorePrefix,
			"last_epoch",
			collections.Uint32Value,
		),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	cr.Schema = schema

	return cr
}

func (r *CommitteeRepository) GetLastEpoch(ctx sdk.Context) (uint32, error) {
	return r.LastEpoch.Get(ctx)
}

func (r *CommitteeRepository) SetLastEpoch(ctx sdk.Context, epoch uint32) error {
	return r.LastEpoch.Set(ctx, epoch)
}

func (r *CommitteeRepository) GetCommittee(ctx sdk.Context, epoch uint32) (types.Committee, error) {
	return r.Committees.Get(ctx, epoch)
}

func (r *CommitteeRepository) GetLastCommittee(ctx sdk.Context) (types.Committee, error) {
	epoch, err := r.GetLastEpoch(ctx)
	if err != nil {
		return types.Committee{}, err
	}
	return r.GetCommittee(ctx, epoch)
}

func (r *CommitteeRepository) SetCommittee(ctx sdk.Context, epoch uint32, committee types.Committee) error {
	return r.Committees.Set(ctx, epoch, committee)
}

func (r *CommitteeRepository) HasCommittee(ctx sdk.Context, epoch uint32) (bool, error) {
	return r.Committees.Has(ctx, epoch)
}

func (r *CommitteeRepository) IterateCommittees(ctx sdk.Context, fn func(committee types.Committee) error) error {
	i, err := r.Committees.Iterate(ctx, nil)
	if err != nil {
		return err
	}
	defer i.Close()

	for ; i.Valid(); i.Next() {
		committee, iterErr := i.Value()
		if iterErr != nil {
			return iterErr
		}
		if iterErr = fn(committee); iterErr != nil {
			return iterErr
		}
	}

	return nil
}
