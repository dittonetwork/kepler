package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SharedEntropy:         nil,
		QuorumParams:          nil,
		AlliancesTimelineList: []AlliancesTimeline{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in alliancesTimeline
	alliancesTimelineIdMap := make(map[uint64]bool)
	alliancesTimelineCount := gs.GetAlliancesTimelineCount()
	for _, elem := range gs.AlliancesTimelineList {
		if _, ok := alliancesTimelineIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for alliancesTimeline")
		}
		if elem.Id >= alliancesTimelineCount {
			return fmt.Errorf("alliancesTimeline id should be lower or equal than the last id")
		}
		alliancesTimelineIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
