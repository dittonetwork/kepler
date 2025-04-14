package types

import (
	"encoding/json"
	"fmt"
)

func (s BondStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *BondStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	switch str {
	case "UnspecifiedStatus":
		*s = UnspecifiedStatus
	case "Unbonded":
		*s = Unbonded
	case "Unbonding":
		*s = Unbonding
	case "Bonded":
		*s = Bonded
	default:
		return fmt.Errorf("unknown BondStatus: %s", str)
	}

	return nil
}

func (p Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *Protocol) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	switch str {
	case "UnspecifiedProtocol":
		*p = UnspecifiedProtocol
	case "Ditto":
		*p = Ditto
	case "Symbiotic":
		*p = Symbiotic
	case "EigenLayer":
		*p = EigenLayer
	default:
		return fmt.Errorf("unknown Protocol: %s", str)
	}

	return nil
}
