package types

import "cosmossdk.io/math"

// NewCommissionRates returns an initialized validator commission rates.
func NewCommissionRates(rate, maxRate, maxChangeRate math.LegacyDec) CommissionRates {
	return CommissionRates{
		Rate:          rate,
		MaxRate:       maxRate,
		MaxChangeRate: maxChangeRate,
	}
}
