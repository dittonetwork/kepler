package types_test

import (
	"testing"

	"kepler/x/alliance/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				SharedEntropy: &types.SharedEntropy{
					Entropy: 32,
				},
				QuorumParams: &types.QuorumParams{
					MaxParticipants:  52,
					ThresholdPercent: 37,
					LifetimeInBlocks: 94,
				},
				AlliancesTimelineList: []types.AlliancesTimeline{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				AlliancesTimelineCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated alliancesTimeline",
			genState: &types.GenesisState{
				AlliancesTimelineList: []types.AlliancesTimeline{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid alliancesTimeline count",
			genState: &types.GenesisState{
				AlliancesTimelineList: []types.AlliancesTimeline{
					{
						Id: 1,
					},
				},
				AlliancesTimelineCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
