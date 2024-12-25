package cli

import (
	flag "github.com/spf13/pflag"

	"kepler/x/staking/types"
)

const (
	FlagAddressValidator    = "validator"
	FlagAddressValidatorSrc = "addr-validator-source"
	FlagAddressValidatorDst = "addr-validator-dest"
	FlagPubKey              = "pubkey"
	FlagAmount              = "amount"

	FlagMoniker         = "moniker"
	FlagEditMoniker     = "new-moniker"
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

// FlagSetCommissionCreate Returns the FlagSet used for commission create.
func FlagSetCommissionCreate() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagCommissionRate, "", "The initial commission rate percentage")
	fs.String(FlagCommissionMaxRate, "", "The maximum commission rate percentage")
	fs.String(FlagCommissionMaxChangeRate, "", "The maximum commission change rate percentage (per day)")

	return fs
}

// FlagSetMinSelfDelegation Returns the FlagSet used for minimum set delegation.
func FlagSetMinSelfDelegation() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagMinSelfDelegation, "", "The minimum self delegation required on the validator")
	return fs
}

// FlagSetAmount Returns the FlagSet for amount related operations.
func FlagSetAmount() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagAmount, "", "Amount of coins to bond")
	return fs
}

// FlagSetPublicKey Returns the flagset for Public Key related operations.
func FlagSetPublicKey() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagPubKey, "", "The validator's Protobuf JSON encoded public key")
	return fs
}

func flagSetDescriptionEdit() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagEditMoniker, types.DoNotModifyDesc, "The validator's name")
	fs.String(FlagIdentity, types.DoNotModifyDesc, "The (optional) identity signature (ex. UPort or Keybase)")
	fs.String(FlagWebsite, types.DoNotModifyDesc, "The validator's (optional) website")
	fs.String(FlagSecurityContact, types.DoNotModifyDesc, "The validator's (optional) security contact email")
	fs.String(FlagDetails, types.DoNotModifyDesc, "The validator's (optional) details")

	return fs
}

func flagSetCommissionUpdate() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagCommissionRate, "", "The new commission rate percentage")

	return fs
}
