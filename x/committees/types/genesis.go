package types

import

// DefaultGenesis returns the default genesis state
"fmt"

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(), CommitteesList: []Committees{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	committeesIdMap := make(map[uint64]bool)
	committeesCount := gs.GetCommitteesCount()
	for _, elem := range gs.CommitteesList {
		if _, ok := committeesIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for committees")
		}
		if elem.Id >= committeesCount {
			return fmt.Errorf("committees id should be lower or equal than the last id")
		}
		committeesIdMap[elem.Id] = true
	}

	return gs.Params.Validate()
}
