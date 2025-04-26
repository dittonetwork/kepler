package keeper

import (
	"context"
	"fmt"

	addresscodec "cosmossdk.io/core/address"
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
	CreateCommittee(ctx sdk.Context, epoch int64) (types.Committee, error)

	// GetAuthority returns the module's authority.
	GetAuthority() string

	// SetParams updates the committee module's parameters.
	SetParams(ctx context.Context, params types.Params) error

	// HandleReport handles a report message.
	HandleReport(ctx sdk.Context, msg *types.MsgSendReport) error
}

type Keeper struct {
	cdc             codec.Codec
	amino           *codec.LegacyAmino
	router          baseapp.MessageRouter
	valAddressCodec addresscodec.Codec
	logger          log.Logger

	repository types.Repository

	account   types.AccountKeeper
	bank      types.BankKeeper
	restaking types.RestakingKeeper

	epochMainID string

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

func NewKeeper(
	authority string,
	account types.AccountKeeper,
	bank types.BankKeeper,
	restaking types.RestakingKeeper,
	repo types.Repository,
	logger log.Logger,
	router baseapp.MessageRouter,
	amino *codec.LegacyAmino,
	cdc codec.Codec,
	epochMainID string,
	valAddressCodec addresscodec.Codec,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	k := Keeper{
		authority:       authority,
		account:         account,
		bank:            bank,
		restaking:       restaking,
		repository:      repo,
		router:          router,
		logger:          logger,
		amino:           amino,
		cdc:             cdc,
		epochMainID:     epochMainID,
		valAddressCodec: valAddressCodec,
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
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
