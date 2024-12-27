package types

func NewMsgUpdateValidatorsStakes(creator string, aaddresses []string, stakes []uint64) *MsgUpdateValidatorsStakes {
	return &MsgUpdateValidatorsStakes{
		Creator:    creator,
		Aaddresses: aaddresses,
		Stakes:     stakes,
	}
}
