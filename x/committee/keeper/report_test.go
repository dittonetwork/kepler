package keeper_test

import (
	"testing"

	"cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/dittonetwork/kepler/x/committee/keeper"
	committeetestutil "github.com/dittonetwork/kepler/x/committee/testutil"
	"github.com/dittonetwork/kepler/x/committee/types"
)

type ReportTestSuite struct {
	suite.Suite
	ctx         sdk.Context
	keeper      keeper.Keeper
	committeeID string
	committee   types.Committee
	repository  *committeetestutil.MockRepository
}

func TestReportTestSuite(t *testing.T) {
	suite.Run(t, new(ReportTestSuite))
}

func (suite *ReportTestSuite) SetupTest() {
	// Create a new context
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(suite.T(), storeKey, storetypes.NewTransientStoreKey("transient_test"))
	suite.ctx = testCtx.Ctx

	// Create a new controller for mocks
	ctrl := gomock.NewController(suite.T())
	suite.repository = committeetestutil.NewMockRepository(ctrl)

	// Create a test committee
	suite.committeeID = "test-committee"
	suite.committee = types.Committee{
		Epoch:   10,
		Address: "cosmos1kkyr80lkuku58h7e2v84egemscmem304mdra4f",
	}

	// Create the keeper with mocks
	suite.keeper = keeper.NewKeeper(
		"cosmos1kkyr80lkuku58h7e2v84egemscmem304mdra4f",
		nil, // executor keeper
		nil, // bank keeper,
		nil, // restaking keeper
		suite.repository,
		suite.ctx.Logger(),
		nil,                    // router
		codec.NewLegacyAmino(), // amino codec
		codec.NewProtoCodec(cdctypes.NewInterfaceRegistry()), // proto codec
		"hour",
		address.NewBech32Codec("dittovaloper"),
	)
}

func (suite *ReportTestSuite) TestHandleReport_InvalidCommitteeAddress() {
	// Setup mock expectations
	suite.repository.EXPECT().
		GetLastCommittee(suite.ctx).
		Return(suite.committee, nil)

	suite.repository.EXPECT().
		GetCommittee(suite.ctx, gomock.Any()).
		Return(types.Committee{
			Epoch:   10,
			Address: "cosmos1testaddressererer",
		}, nil)

	// Create a report with invalid committee address
	msg := &types.MsgSendReport{
		Creator: "cosmos1testaddress",
		Report: &types.Report{
			CommitteeId: suite.committeeID,
		},
	}

	err := suite.keeper.HandleReport(suite.ctx, msg)
	suite.Require().Error(err)
	suite.Require().Contains(err.Error(), "invalid committee")
}

func (suite *ReportTestSuite) TestHandleReport_InvalidEpoch() {
	// Setup mock expectations
	suite.repository.EXPECT().
		GetLastCommittee(suite.ctx).
		Return(suite.committee, nil)

	suite.repository.EXPECT().
		GetCommittee(suite.ctx, gomock.Any()).
		Return(types.Committee{
			Epoch:   8,
			Address: "cosmos1testaddress",
		}, nil)

	// Create a report with invalid committee address
	msg := &types.MsgSendReport{
		Creator: "cosmos1testaddress",
		Report: &types.Report{
			CommitteeId: suite.committeeID,
		},
	}

	err := suite.keeper.HandleReport(suite.ctx, msg)
	suite.Require().Error(err)
	suite.Require().Contains(err.Error(), "invalid committee")
}

func (suite *ReportTestSuite) TestHandleReport_ValidReport() {
	// Setup mock expectations
	suite.repository.EXPECT().
		GetLastCommittee(suite.ctx).
		Return(suite.committee, nil)

	// Create a valid report with a test message
	msg := &types.MsgSendReport{
		Creator: suite.committee.Address,
		Report: &types.Report{
			CommitteeId: suite.committeeID,
			Messages:    []*cdctypes.Any{},
		},
	}

	err := suite.keeper.HandleReport(suite.ctx, msg)
	suite.Require().NoError(err)

	// Verify events were emitted
	events := suite.ctx.EventManager().Events()
	suite.Require().Len(events, 1)
	suite.Require().Equal("kepler.committee.EventReportReceived", events[0].Type)
	suite.Require().Equal("\""+suite.committee.Address+"\"", events[0].Attributes[0].Value)
	suite.Require().Equal("\"0\"", events[0].Attributes[2].Value) // report_count
}

func (suite *ReportTestSuite) TestHandleReport_InvalidMessage() {
	// Setup mock expectations
	suite.repository.EXPECT().
		GetLastCommittee(suite.ctx).
		Return(suite.committee, nil)

	// Create a report with an invalid message
	msg := &types.MsgSendReport{
		Creator: suite.committee.Address,
		Report: &types.Report{
			CommitteeId: suite.committeeID,
			Messages: []*cdctypes.Any{
				{
					TypeUrl: "invalid-type-url",
					Value:   []byte("invalid-message"),
				},
			},
		},
	}

	err := suite.keeper.HandleReport(suite.ctx, msg)
	suite.Require().Error(err)
	suite.Require().Contains(err.Error(), "cannot unpack message")
}

func (suite *ReportTestSuite) TestHandleReport_GetLastCommitteeError() {
	// Setup mock expectations to return an error
	suite.repository.EXPECT().
		GetLastCommittee(suite.ctx).
		Return(types.Committee{}, errors.New("9", 9, "repository error"))

	// Create a report
	msg := &types.MsgSendReport{
		Creator: suite.committee.Address,
		Report: &types.Report{
			CommitteeId: suite.committeeID,
		},
	}

	err := suite.keeper.HandleReport(suite.ctx, msg)
	suite.Require().Error(err)
	suite.Require().Contains(err.Error(), "repository error")
}
