package types

func NewGenesisState(params Params, validators []Validator) *GenesisState {
	return &GenesisState{
		Params:     params,
		Validators: validators,
	}
}

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}
