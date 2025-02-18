package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) IsCommitteeExists(_ sdk.Context, _ string) (bool, error) {
	// TODO implement me
	panic("implement me")
}

// func (k Keeper) CanBeSigned(_ sdk.Context, _ string, _ string, _ [][]byte) (bool, error) {
//	// TODO implement me
//	panic("implement me")
//}
