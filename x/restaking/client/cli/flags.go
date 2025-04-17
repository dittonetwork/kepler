package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagPubKey = "pubkey"
	FlagAmount = "amount"

	FlagMoniker         = "moniker"
	FlagIdentity        = "identity"
	FlagWebsite         = "website"
	FlagSecurityContact = "security-contact"
	FlagDetails         = "details"

	FlagCommissionRate          = "commission-rate"
	FlagCommissionMaxRate       = "commission-max-rate"
	FlagCommissionMaxChangeRate = "commission-max-change-rate"

	FlagMinSelfDelegation = "min-self-delegation"

	FlagNodeID  = "node-id"
	FlagIP      = "ip"
	FlagP2PPort = "p2p-port"
)

//nolint:gochecknoinits // no matter
func init() {

}

// FlagSetCommissionCreate Returns the FlagSet used for commission create.
// !! DO NOT USE THIS FLAG !!
func FlagSetCommissionCreate() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagCommissionRate, "", "DO NOT USE THIS FLAG. The commission rate percentage")
	fs.String(FlagCommissionMaxRate, "", "DO NOT USE THIS FLAG. The maximum commission rate percentage")
	fs.String(
		FlagCommissionMaxChangeRate, "",
		"DO NOT USE THIS FLAG. The maximum commission change rate percentage (per day)",
	)

	return fs
}

// FlagSetMinSelfDelegation Returns the FlagSet used for minimum set delegation.
// !! DO NOT USE THIS FLAG !!
func FlagSetMinSelfDelegation() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagMinSelfDelegation, "",
		"DO NOT USE THIS FLAG. The minimum self delegation required on the validator")
	return fs
}

// FlagSetAmount Returns the FlagSet for amount related operations.
// !! DO NOT USE THIS FLAG !!
func FlagSetAmount() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagAmount, "", "DO NOT USE THIS FLAG. The amount to bond")
	return fs
}

// FlagSetPublicKey Returns the flagset for Public Key related operations.
// !! DO NOT USE THIS FLAG !!
func FlagSetPublicKey() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagPubKey, "", "DO NOT USE THIS FLAG. The validator's Protobuf JSON encoded public key")
	return fs
}
