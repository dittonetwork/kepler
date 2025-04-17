package cli

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	flag "github.com/spf13/pflag"

	"cosmossdk.io/core/address"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/dittonetwork/kepler/x/restaking/types"
)

const (
	DefaultP2PPort = 26656
)

// CreateValidatorMsgFlagSet creates a flag set for the create validator message.
// @WARN this flag set backwards compatibility with staking flagset for ignite tool.
func CreateValidatorMsgFlagSet(ipDefault string) *flag.FlagSet {
	fsCreateValidator := flag.NewFlagSet("", flag.ContinueOnError)
	fsCreateValidator.String(FlagIP, ipDefault, "The node's public P2P IP")
	fsCreateValidator.Uint(FlagP2PPort, DefaultP2PPort, "The node's public P2P port")
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

	return fsCreateValidator
}

type TxBondValidatorConfig struct {
	ChainID string
	NodeID  string
	Moniker string

	IP              string
	P2PPort         uint
	Website         string
	SecurityContact string
	Details         string
	Identity        string
}

// PrepareConfigForTxBondValidator prepares the configuration for the tx bond validator command.
func PrepareConfigForTxBondValidator(
	flagSet *flag.FlagSet,
	moniker, nodeID, chainID string,
) (TxBondValidatorConfig, error) {
	c := TxBondValidatorConfig{
		ChainID: chainID,
		NodeID:  nodeID,
		Moniker: moniker,
	}

	ip, err := flagSet.GetString(FlagIP)
	if err != nil {
		return c, err
	}

	if ip == "" {
		_, _ = fmt.Fprintf(os.Stderr, "failed to retrieve an external IP; the tx's memo field will be unset")
	}

	c.P2PPort, err = flagSet.GetUint(FlagP2PPort)
	if err != nil {
		return c, err
	}

	c.Website, err = flagSet.GetString(FlagWebsite)
	if err != nil {
		return c, err
	}

	c.SecurityContact, err = flagSet.GetString(FlagSecurityContact)
	if err != nil {
		return c, err
	}

	c.Details, err = flagSet.GetString(FlagDetails)
	if err != nil {
		return c, err
	}

	c.Identity, err = flagSet.GetString(FlagIdentity)
	if err != nil {
		return c, err
	}

	return c, nil
}

// BuildBondValidatorMsg builds the bond validator message.
func BuildBondValidatorMsg(
	clientCtx client.Context,
	config TxBondValidatorConfig,
	txBuilder tx.Factory,
	generateOnly bool,
	valCodec address.Codec,
) (tx.Factory, sdk.Msg, error) {
	valAddr := clientCtx.GetFromAddress()
	description := types.NewDescription(
		config.Moniker,
		config.Identity,
		config.Website,
		config.SecurityContact,
		config.Details,
	)
	valStr, err := valCodec.BytesToString(valAddr)
	if err != nil {
		return tx.Factory{}, nil, err
	}

	msg := types.NewMsgBondValidator(valStr, description)

	if generateOnly {
		if config.NodeID != "" && config.IP != "" && config.P2PPort > 0 {
			txBuilder = txBuilder.WithMemo(fmt.Sprintf("%s@%s:%d", config.NodeID, config.IP, config.P2PPort))
		}
	}

	return txBuilder, msg, nil
}
