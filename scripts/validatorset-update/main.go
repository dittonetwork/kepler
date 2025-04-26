package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/cometbft/cometbft/rpc/client/http"
	cryptomultisig "github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	restakingtypes "github.com/dittonetwork/kepler/x/restaking/types"

	sdkmath "cosmossdk.io/math"
	"github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/dittonetwork/kepler/api/kepler/restaking"
	committeetypes "github.com/dittonetwork/kepler/x/committee/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	epochNum   = 1
	defaultFee = 30
	gasLimit   = 100_000
)

var (
	// Setup interface registry and register necessary interfaces.
	interfaceRegistry = codectypes.NewInterfaceRegistry()
	amino             = codec.NewLegacyAmino()

	marshaler = codec.NewProtoCodec(interfaceRegistry)
	txConfig  = authtx.NewTxConfig(marshaler, authtx.DefaultSignModes)

	// Create a protocodec for encoding/decoding.
	protoCodec = codec.NewProtoCodec(interfaceRegistry)

	config = sdk.GetConfig()
	ctx    = context.Background()
)

type Proposal struct {
	sdk.Tx

	Sequence uint64
	AccNum   uint64 // can be calculated by each signer but provided for simplify example
}

func main() {
	cometRPC, err := client.NewClientFromNode("http://localhost:26657")
	if err != nil {
		log.Fatal("create client: ", err)
	}

	clientConn, err := grpc.NewClient(
		"localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Register interfaces
	banktypes.RegisterInterfaces(interfaceRegistry)
	authtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	committeetypes.RegisterInterfaces(interfaceRegistry)
	restakingtypes.RegisterInterfaces(interfaceRegistry)
	sdktypes.RegisterInterfaces(interfaceRegistry)

	// Configure the SDK
	config.SetBech32PrefixForAccount("ditto", "dittopub")

	homeDir := fmt.Sprintf("%s/.kepler", os.Getenv("HOME"))
	ckr, err := keyring.New("kepler", "test", homeDir, os.Stdin, protoCodec)
	if err != nil {
		log.Fatal(err)
	}

	participants := []string{"alice"}

	// Create a new report message
	proposal, clientCtx, acc, err := CreateProposal(clientConn, cometRPC, ckr, participants)
	if err != nil {
		log.Fatal(err)
	}

	signatures := make([]signingtypes.SignatureV2, 0, len(participants))
	for _, name := range participants {
		var sig *signingtypes.SignatureV2
		sig, err = SignByParticipant(ctx, ckr, proposal, name, clientConn)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(
			"Transaction signed successfully. ", "name: ", name, "sig data: ", sig.Data,
		)

		signatures = append(signatures, *sig)
	}

	signedTx, err := AggregateSignatures(clientCtx, signatures, proposal, acc)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Sign aggregated successfully. ", "txb: ", signedTx.GetFee().String())

	txJSON, err := clientCtx.TxConfig.TxJSONEncoder()(signedTx)
	if err != nil {
		log.Fatal("to json: ", err)
	}

	prettyPrintJSON(txJSON)

	txBytes, err := clientCtx.TxConfig.TxEncoder()(signedTx)
	if err != nil {
		log.Fatal("to bytes: ", err)
	}

	txRes, err := clientCtx.BroadcastTxSync(txBytes)
	if err != nil {
		log.Fatal("broadcast tx: ", err)
	}

	if txRes.Code != 0 {
		log.Fatalf("error log: %s", txRes.RawLog)
	}

	log.Println(
		"Transaction sent successfully. ",
		txRes.TxHash,
		txRes.Info,
		txRes.Code,
	)
}

func CreateProposal(
	clientConn *grpc.ClientConn, cometRPC *http.HTTP, ckr keyring.Keyring, participants []string,
) (*Proposal, *client.Context, *multisig.LegacyAminoPubKey, error) {
	participantsPks := make([]cryptotypes.PubKey, 0, len(participants))

	for _, name := range participants {
		record, err := ckr.Key(name)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to get key: %w", err)
		}

		pk, err := record.GetPubKey()
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to get key: %w", err)
		}

		participantsPks = append(participantsPks, pk)
	}

	multisigAcc := multisig.NewLegacyAminoPubKey(getThreshold(len(participants)), participantsPks)

	clientCtx := client.Context{}.
		WithCodec(protoCodec).
		WithInterfaceRegistry(interfaceRegistry).
		WithTxConfig(txConfig).
		WithLegacyAmino(amino).
		WithChainID("kepler").
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithKeyring(ckr).
		WithNodeURI("http://localhost:26657").
		WithGRPCClient(clientConn).
		WithClient(cometRPC).
		WithFromAddress(sdk.AccAddress(multisigAcc.Address()))

	accnum, seq, err := clientCtx.AccountRetriever.GetAccountNumberSequence(
		clientCtx, sdk.AccAddress(multisigAcc.Address()),
	)
	if err != nil {
		log.Fatal("get account: ", err)
	}

	txf := tx.Factory{}.
		WithChainID("kepler").
		WithTxConfig(txConfig).
		WithKeybase(clientCtx.Keyring).
		WithSignMode(signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON).
		WithSequence(seq).
		WithAccountNumber(accnum)

	// Create a new report message
	validatorSetUpdateMsg := BuildValidatorSetUpdateMsg(epochNum)
	reportMsg := BuildReportMsg(multisigAcc.Address(), epochNum, []sdk.Msg{validatorSetUpdateMsg})

	if m, ok := reportMsg.(sdk.HasValidateBasic); ok {
		if err = m.ValidateBasic(); err != nil {
			log.Fatal(err)
		}
	}

	txBuilder := clientCtx.TxConfig.NewTxBuilder()

	txBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewCoin("ditto", sdkmath.NewInt(defaultFee))))
	txBuilder.SetGasLimit(gasLimit)

	if err = txBuilder.SetMsgs(reportMsg); err != nil {
		return nil, nil, nil, fmt.Errorf("set messages: %w", err)
	}

	proposal := &Proposal{
		Tx:       txBuilder.GetTx(),
		Sequence: txf.Sequence(),
		AccNum:   txf.AccountNumber(),
	}

	log.Printf(
		"generated proposal with accnum: %d and sequence: %d", txf.AccountNumber(), txf.Sequence(),
	)

	return proposal, &clientCtx, multisigAcc, nil
}

func AggregateSignatures(
	clientCtx *client.Context, signatures []signingtypes.SignatureV2, proposal *Proposal,
	committeePk *multisig.LegacyAminoPubKey,
) (signing.Tx, error) {
	txf := tx.Factory{}.
		WithChainID("kepler").
		WithTxConfig(txConfig).
		WithKeybase(clientCtx.Keyring).
		WithSignMode(signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON).
		WithSequence(proposal.Sequence).
		WithAccountNumber(proposal.AccNum)

	// create signatures slice for add them to committee signature
	sigs := cryptomultisig.NewMultisig(len(signatures))

	// aggregate all signatures to one struct
	for _, sig := range signatures {
		if err := cryptomultisig.AddSignatureV2(sigs, sig, committeePk.GetPubKeys()); err != nil {
			return nil, fmt.Errorf("add signature: %w", err)
		}
	}

	txb, err := txConfig.WrapTxBuilder(proposal.Tx)
	if err != nil {
		return nil, fmt.Errorf("wrap tx: %w", err)
	}

	signature := signingtypes.SignatureV2{
		PubKey:   committeePk,
		Data:     sigs,
		Sequence: txf.Sequence(),
	}

	if err = txb.SetSignatures(signature); err != nil {
		return nil, fmt.Errorf("set signatures: %w", err)
	}

	return txb.GetTx(), nil
}

func SignByParticipant(
	ctx context.Context, ckr keyring.Keyring, proposal *Proposal, name string,
	clientConn *grpc.ClientConn,
) (*signingtypes.SignatureV2, error) {
	clientCtx := client.Context{}.
		WithCodec(protoCodec).
		WithInterfaceRegistry(interfaceRegistry).
		WithTxConfig(txConfig).
		WithLegacyAmino(amino).
		WithChainID("kepler").
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithKeyring(ckr).
		WithNodeURI("http://localhost:26657").
		WithGRPCClient(clientConn).
		WithFromName(name)

	txf := tx.Factory{}.
		WithChainID("kepler").
		WithTxConfig(txConfig).
		WithKeybase(clientCtx.Keyring).
		WithSignMode(signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON).
		WithSequence(proposal.Sequence).
		WithAccountNumber(proposal.AccNum)

	txb, err := txConfig.WrapTxBuilder(proposal.Tx)
	if err != nil {
		log.Fatal("wrap tx: ", err)
	}

	if err = tx.Sign(ctx, txf, name, txb, false); err != nil {
		return nil, fmt.Errorf("sign by %s: %w", name, err)
	}

	signatures, err := txb.GetTx().GetSignaturesV2()
	if err != nil {
		return nil, fmt.Errorf("get signatures: %w", err)
	}

	if len(signatures) != 1 {
		return nil, fmt.Errorf("expected one signature, got %d", len(signatures))
	}

	return &signatures[0], nil
}

func BuildReportMsg(addr types.Address, epochNum int64, messages []sdk.Msg) sdk.Msg {
	messagesAnyType := make([]*codectypes.Any, len(messages))
	for i, msg := range messages {
		msgAny, err := codectypes.NewAnyWithValue(msg)
		if err != nil {
			panic(err)
		}
		messagesAnyType[i] = msgAny
	}

	return &committeetypes.MsgSendReport{
		Creator: sdk.AccAddress(addr).String(),
		EpochId: epochNum,
		Report: &committeetypes.Report{
			CommitteeId:      strconv.Itoa(int(epochNum)),
			ExecutionContext: nil,
			Messages:         messagesAnyType,
		},
	}
}

// BuildValidatorSetUpdateMsg creates a MsgUpdateValidatorsSet message.
func BuildValidatorSetUpdateMsg(epochNum int64) sdk.Msg {
	aliceValPk := &ed25519.PubKey{
		Key: []byte("KJPZk9bTZGdgsJ0xDIpcg9RbIRHuPwjlh3oiZ2C+5Cs="),
	}

	aliceValPkAny, err := codectypes.NewAnyWithValue(aliceValPk)
	if err != nil {
		panic(err)
	}

	bobValPk := &ed25519.PubKey{
		Key: []byte("dcHikhOeHpXzLzX+IFlz6HuNs5ILr3v/OG5NfGQuvKE="),
	}

	bobValPkAny, err := codectypes.NewAnyWithValue(bobValPk)
	if err != nil {
		panic(err)
	}

	return &restaking.MsgUpdateValidatorsSet{
		Authority: authtypes.NewModuleAddress(committeetypes.ModuleName).String(),
		Operators: []*restaking.Operator{
			// Bob
			{
				Address: "0xf7e1dCd1d199f7C42BD12502b68f58301b5b4d98",
				ConsensusPubkey: &anypb.Any{
					TypeUrl: "/cosmos.crypto.ed25519.PubKey",
					Value:   aliceValPkAny.Value,
				},
				IsEmergency: true,
				Status:      restaking.BondStatus_BOND_STATUS_BONDED,
				VotingPower: 1000000, //nolint: mnd // no matter
				Protocol:    restaking.Protocol_PROTOCOL_DITTO,
			},
			// Alice
			{
				Address: "0x910cB6A0937ECeBA1EDF4F505F1b86D3234a4Fe9",
				ConsensusPubkey: &anypb.Any{
					TypeUrl: "/cosmos.crypto.ed25519.PubKey",
					Value:   bobValPkAny.Value,
				},
				IsEmergency: true,
				Status:      restaking.BondStatus_BOND_STATUS_BONDED,
				VotingPower: 1000000, //nolint: mnd // no matter
				Protocol:    restaking.Protocol_PROTOCOL_DITTO,
			},
		},
		Info: &restaking.UpdateInfo{
			EpochNum:    epochNum,
			Timestamp:   timestamppb.New(time.Now()),
			BlockHeight: 14188, //nolint: mnd // no matter
			BlockHash:   "0x1234567890abcdef",
		},
	}
}

func getThreshold(participantsCount int) int {
	if participantsCount == 1 {
		return 1
	}

	// Calculate the threshold for BFT consensus
	// https://pmg.csail.mit.edu/papers/osdi99.pdf
	return int(math.Floor(2*(float64(participantsCount)-1)/3) + 1) //nolint:mnd // no matter
}

func prettyPrintJSON(data []byte) {
	var pretty bytes.Buffer
	err := json.Indent(&pretty, data, "", "  ")
	if err != nil {
		log.Fatalf("failed to indent json: %v", err)
		return
	}
	log.Println(pretty.String())
}
