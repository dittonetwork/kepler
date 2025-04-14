package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dittonetwork/kepler/x/committee/types"
)

type CommitteeKeeper interface {
	// IsCommitteeExists returns true if the committee exists
	// Deprecated: CanBeSigned returns true if the message can be signed.
	IsCommitteeExists(ctx sdk.Context, committeeID string) (bool, error)

	// CreateCommittee creates a new committee for the given epoch.
	CreateCommittee(ctx sdk.Context, epoch uint32) (types.Committee, error)

	// GetAuthority returns the module's authority.
	GetAuthority() string

	// SetParams updates the committee module's parameters.
	SetParams(ctx context.Context, params types.Params) error

	// HandleReport handles a report message.
	HandleReport(ctx sdk.Context, msg *types.MsgSendReport) error
}

type Keeper struct {
	cdc    codec.Codec
	amino  *codec.LegacyAmino
	router baseapp.MessageRouter

	repository types.Repository

	executors types.ExecutorsKeeper
	restaking types.RestakingKeeper

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

func NewKeeper(
	authority string,
	executors types.ExecutorsKeeper,
	restaking types.RestakingKeeper,
	repo types.Repository,
	router baseapp.MessageRouter,
	amino *codec.LegacyAmino,
	cdc codec.Codec,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	k := Keeper{
		authority:  authority,
		executors:  executors,
		restaking:  restaking,
		repository: repo,
		router:     router,
		amino:      amino,
		cdc:        cdc,
	}

	return k
}

// Deprecated: This method is deprecated and will be reworked.
// CanBeSigned returns true if the message can be signed.
func (k Keeper) CanBeSigned(
	_ sdk.Context,
	_ string,
	_ string,
	_ [][]byte,
	_ []byte,
) (bool, error) {
	return true, nil
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx context.Context) log.Logger {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	return sdkCtx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
