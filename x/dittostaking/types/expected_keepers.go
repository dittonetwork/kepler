package types

import "context"

type ConsensusKeeper interface {
	ValidatorPubKeyTypes(context.Context) ([]string, error)
}
