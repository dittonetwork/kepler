package types

import (
	"cosmossdk.io/errors"
	cmtprotocrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ConsPubKey returns the validator's PubKey as cryptotypes.PubKey.
func (v Validator) ConsPubKey() (cryptotypes.PubKey, error) {
	pk, ok := v.ConsensusPubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, errors.Wrapf(
			sdkerrors.ErrInvalidType,
			"expecting cryptotypes.PubKey, got %T", v.ConsensusPubkey,
		)
	}

	return pk, nil
}

// CmtConsPublicKey casts Validator.ConsPublicKey to cmtprotocrypto.PublicKey.
func (v Validator) CmtConsPublicKey() (cmtprotocrypto.PublicKey, error) {
	pk, err := v.ConsPubKey()
	if err != nil {
		return cmtprotocrypto.PublicKey{}, err
	}

	tmPk, err := cryptocodec.ToCmtProtoPublicKey(pk)
	if err != nil {
		return cmtprotocrypto.PublicKey{}, err
	}

	return tmPk, nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces.
func (v Validator) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var pk cryptotypes.PubKey
	return unpacker.UnpackAny(v.ConsensusPubkey, &pk)
}

// ConvertToOperator converts a Validator to an Operator.
func (v Validator) ConvertToOperator() *Operator {
	return &Operator{
		Address:         v.OperatorAddress,
		ConsensusPubkey: v.ConsensusPubkey,
		IsEmergency:     v.IsEmergency,
		Status:          v.Status,
		VotingPower:     v.VotingPower,
		Protocol:        v.Protocol,
	}
}

func (v *Validator) UpdateOperatorInfo(operator Operator) {
	v.ConsensusPubkey = operator.ConsensusPubkey
	v.Status = operator.Status
	v.Protocol = operator.Protocol
	v.VotingPower = operator.VotingPower
	v.IsEmergency = operator.IsEmergency
}
