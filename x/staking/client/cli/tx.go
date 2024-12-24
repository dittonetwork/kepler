package cli

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	flag "github.com/spf13/pflag"
	"kepler/x/staking/types"
	"os"
)

// default values
var (
	DefaultTokens                  = sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction)
	defaultAmount                  = DefaultTokens.String() + sdk.DefaultBondDenom
	defaultCommissionRate          = "0.1"
	defaultCommissionMaxRate       = "0.2"
	defaultCommissionMaxChangeRate = "0.01"
	defaultMinSelfDelegation       = "1"
)

// CreateValidatorMsgFlagSet Returns the flagset, particular flags, and a description of defaults
// this is anticipated to be used with the gen-tx
func CreateValidatorMsgFlagSet(ipDefault string) (fs *flag.FlagSet, defaultsDesc string) {
	fsCreateValidator := flag.NewFlagSet("", flag.ContinueOnError)
	fsCreateValidator.String(FlagIP, ipDefault, "The node's public P2P IP")
	fsCreateValidator.Uint(FlagP2PPort, 26656, "The node's public P2P port")
	fsCreateValidator.String(FlagNodeID, "", "The node's NodeID")
	fsCreateValidator.String(FlagMoniker, "", "The validator's (optional) moniker")
	fsCreateValidator.String(FlagWebsite, "", "The validator's (optional) website")
	fsCreateValidator.String(FlagSecurityContact, "", "The validator's (optional) security contact email")
	fsCreateValidator.String(FlagDetails, "", "The validator's (optional) details")
	fsCreateValidator.String(FlagIdentity, "", "The (optional) identity signature (ex. UPort or Keybase)")
	fsCreateValidator.AddFlagSet(FlagSetCommissionCreate())
	fsCreateValidator.AddFlagSet(FlagSetMinSelfDelegation())
	fsCreateValidator.AddFlagSet(FlagSetAmount())
	fsCreateValidator.AddFlagSet(FlagSetPublicKey())

	defaultsDesc = fmt.Sprintf(`
	delegation amount:           %s
	commission rate:             %s
	commission max rate:         %s
	commission max change rate:  %s
	minimum self delegation:     %s
`,
		defaultAmount,
		defaultCommissionRate,
		defaultCommissionMaxRate,
		defaultCommissionMaxChangeRate,
		defaultMinSelfDelegation,
	)

	return fsCreateValidator, defaultsDesc
}

type TxCreateValidatorConfig struct {
	ChainID string
	NodeID  string
	Moniker string

	Amount string

	CommissionRate          string
	CommissionMaxRate       string
	CommissionMaxChangeRate string
	MinSelfDelegation       string

	PubKey cryptotypes.PubKey

	IP              string
	P2PPort         uint
	Website         string
	SecurityContact string
	Details         string
	Identity        string
}

func PrepareConfigForTxCreateValidator(flagSet *flag.FlagSet, moniker, nodeID, chainID string, valPubKey cryptotypes.PubKey) (TxCreateValidatorConfig, error) {
	c := TxCreateValidatorConfig{}

	ip, err := flagSet.GetString(FlagIP)
	if err != nil {
		return c, err
	}

	if ip == "" {
		_, _ = fmt.Fprintf(os.Stderr, "failed to retrieve an external IP; the tx's memo field will be unset")
	}

	p2pPort, err := flagSet.GetUint(FlagP2PPort)
	if err != nil {
		return c, err
	}

	website, err := flagSet.GetString(FlagWebsite)
	if err != nil {
		return c, err
	}

	securityContact, err := flagSet.GetString(FlagSecurityContact)
	if err != nil {
		return c, err
	}

	details, err := flagSet.GetString(FlagDetails)
	if err != nil {
		return c, err
	}

	identity, err := flagSet.GetString(FlagIdentity)
	if err != nil {
		return c, err
	}

	c.Amount, err = flagSet.GetString(FlagAmount)
	if err != nil {
		return c, err
	}

	c.CommissionRate, err = flagSet.GetString(FlagCommissionRate)
	if err != nil {
		return c, err
	}

	c.CommissionMaxRate, err = flagSet.GetString(FlagCommissionMaxRate)
	if err != nil {
		return c, err
	}

	c.CommissionMaxChangeRate, err = flagSet.GetString(FlagCommissionMaxChangeRate)
	if err != nil {
		return c, err
	}

	c.MinSelfDelegation, err = flagSet.GetString(FlagMinSelfDelegation)
	if err != nil {
		return c, err
	}

	c.IP = ip
	c.P2PPort = p2pPort
	c.Website = website
	c.SecurityContact = securityContact
	c.Identity = identity
	c.NodeID = nodeID
	c.PubKey = valPubKey
	c.Website = website
	c.SecurityContact = securityContact
	c.Details = details
	c.Identity = identity
	c.ChainID = chainID
	c.Moniker = moniker

	if c.Amount == "" {
		c.Amount = defaultAmount
	}

	if c.CommissionRate == "" {
		c.CommissionRate = defaultCommissionRate
	}

	if c.CommissionMaxRate == "" {
		c.CommissionMaxRate = defaultCommissionMaxRate
	}

	if c.CommissionMaxChangeRate == "" {
		c.CommissionMaxChangeRate = defaultCommissionMaxChangeRate
	}

	if c.MinSelfDelegation == "" {
		c.MinSelfDelegation = defaultMinSelfDelegation
	}

	return c, nil
}

// BuildCreateValidatorMsg makes a new MsgCreateValidator.
func BuildCreateValidatorMsg(clientCtx client.Context, config TxCreateValidatorConfig, txBldr tx.Factory, generateOnly bool) (tx.Factory, sdk.Msg, error) {
	amounstStr := config.Amount
	amount, err := sdk.ParseCoinNormalized(amounstStr)
	if err != nil {
		return txBldr, nil, err
	}

	valAddr := clientCtx.GetFromAddress()
	description := types.NewDescription(
		config.Moniker,
		config.Identity,
		config.Website,
		config.SecurityContact,
		config.Details,
		nil,
	)

	// get the initial validator commission parameters
	rateStr := config.CommissionRate
	maxRateStr := config.CommissionMaxRate
	maxChangeRateStr := config.CommissionMaxChangeRate
	commissionRates, err := buildCommissionRates(rateStr, maxRateStr, maxChangeRateStr)
	if err != nil {
		return txBldr, nil, err
	}

	// get the initial validator min self delegation
	msbStr := config.MinSelfDelegation
	minSelfDelegation, ok := math.NewIntFromString(msbStr)

	if !ok {
		return txBldr, nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "minimum self delegation must be a positive integer")
	}

	valCodec := clientCtx.TxConfig.SigningContext().ValidatorAddressCodec()
	valStr, err := valCodec.BytesToString(valAddr)
	if err != nil {
		return txBldr, nil, err
	}

	msg, err := types.NewMsgCreateValidator(
		valStr,
		config.PubKey,
		amount,
		description,
		commissionRates,
		minSelfDelegation,
	)
	if err != nil {
		return txBldr, msg, err
	}

	if generateOnly {
		ip := config.IP
		p2pPort := config.P2PPort
		nodeID := config.NodeID

		if nodeID != "" && ip != "" && p2pPort > 0 {
			txBldr = txBldr.WithMemo(fmt.Sprintf("%s@%s:%d", nodeID, ip, p2pPort))
		}
	}

	return txBldr, msg, nil
}
