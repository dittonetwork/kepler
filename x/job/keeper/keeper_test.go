package keeper_test

import (
	"kepler/x/job/keeper"
	"kepler/x/job/types"
	"kepler/x/job/types/mock"
	"testing"

	"go.uber.org/mock/gomock"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectestutils "github.com/cosmos/cosmos-sdk/codec/testutil"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx       sdk.Context
	keeper    keeper.Keeper
	committee *mock.MockCommitteeKeeper
}

func (s *KeeperTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.committee = mock.NewMockCommitteeKeeper(ctrl)

	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	s.ctx = testutil.DefaultContext(storeKey, storetypes.NewTransientStoreKey("transient_"+types.StoreKey))

	interfaceRegistry := codectestutils.CodecOptions{}.NewInterfaceRegistry()
	c := codec.NewProtoCodec(interfaceRegistry)

	s.keeper = keeper.NewKeeper(c, runtime.NewKVStoreService(storeKey), log.NewNopLogger(), s.committee)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) Test_CreateJob() {
	var address = "cosmos1ghekyjucln7y67ntx7cf27m9dpuxxemn4c8g4r"

	cases := map[string]struct {
		preRun    func() (types.Job, error)
		expErr    bool
		expErrMsg string
	}{
		"ok": {
			preRun: func() (types.Job, error) {
				s.committee.EXPECT().IsCommitteeExists(s.ctx, "1").Return(true, nil)
				s.committee.EXPECT().CanBeSigned(s.ctx, "1", "1", [][]byte{
					[]byte("sign1"),
					[]byte("sign2"),
					[]byte("sign3"),
				}, gomock.Any()).Return(true, nil)
				return types.Job{
					Id:              123,
					Status:          types.Job_STATUS_EXECUTED,
					ChainId:         "1",
					AutomationId:    1,
					TxHash:          "0x",
					ExecutorAddress: address,
					CommitteeId:     "1",
					Signs: [][]byte{
						[]byte("sign1"),
						[]byte("sign2"),
						[]byte("sign3"),
					},
				}, nil
			},
		},
		"already exists": {
			preRun: func() (types.Job, error) {
				s.committee.EXPECT().IsCommitteeExists(s.ctx, "1").Return(true, nil)
				s.committee.EXPECT().CanBeSigned(s.ctx, "1", "1", [][]byte{
					[]byte("sign1"),
					[]byte("sign2"),
					[]byte("sign3"),
				}, gomock.Any()).Return(true, nil)
				err := s.keeper.CreateJob(s.ctx, types.Job{Id: 111, CommitteeId: "1", ChainId: "1", Signs: [][]byte{
					[]byte("sign1"),
					[]byte("sign2"),
					[]byte("sign3"),
				}})
				if err != nil {
					return types.Job{}, err
				}
				return types.Job{
					Id:              111,
					Status:          types.Job_STATUS_EXECUTED,
					ChainId:         "1",
					AutomationId:    1,
					TxHash:          "0x",
					ExecutorAddress: address,
					CommitteeId:     "1",
					Signs: [][]byte{
						[]byte("sign1"),
						[]byte("sign2"),
						[]byte("sign3"),
					},
				}, nil
			},
			expErr:    true,
			expErrMsg: "job with id 111 already exists: job already exists",
		},
		"committee doesn't exists": {
			preRun: func() (types.Job, error) {
				s.committee.EXPECT().IsCommitteeExists(s.ctx, "1").Return(false, nil)
				return types.Job{Id: 124,
					Status:          types.Job_STATUS_EXECUTED,
					ChainId:         "1",
					AutomationId:    1,
					TxHash:          "0x",
					ExecutorAddress: address,
					CommitteeId:     "1",
					Signs: [][]byte{
						[]byte("sign1"),
					},
				}, nil
			},
			expErr:    true,
			expErrMsg: "committee does not exist",
		},
		"signs invalid": {
			preRun: func() (types.Job, error) {
				s.committee.EXPECT().IsCommitteeExists(s.ctx, "1").Return(true, nil)
				s.committee.EXPECT().CanBeSigned(s.ctx, "1", "1", [][]byte{
					[]byte("sign1"),
				}, gomock.Any()).Return(false, nil)
				return types.Job{
					Id:              125,
					Status:          types.Job_STATUS_EXECUTED,
					ChainId:         "1",
					AutomationId:    1,
					TxHash:          "0x",
					ExecutorAddress: address,
					CommitteeId:     "1",
					Signs: [][]byte{
						[]byte("sign1"),
					},
				}, nil
			},
		},
		"job nil": {
			preRun: func() (types.Job, error) {
				return types.Job{}, nil
			},
			expErr:    true,
			expErrMsg: "job signs is nil",
		},
	}

	for testName, tc := range cases {
		s.Run(testName, func() {
			msg, err := tc.preRun()
			s.Require().NoError(err)
			err = s.keeper.CreateJob(s.ctx, msg)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				s.Require().NoError(err)
			}
		})
	}
}

func (s *KeeperTestSuite) Test_GetJob() {
	cases := map[string]struct {
		preRun func() (uint64, error)
		found  bool
	}{
		"found": {
			preRun: func() (uint64, error) {
				s.committee.EXPECT().IsCommitteeExists(s.ctx, "1").Return(true, nil)
				s.committee.EXPECT().CanBeSigned(s.ctx, "1", "1", [][]byte{
					[]byte("sign1"),
					[]byte("sign2"),
					[]byte("sign3"),
				}, gomock.Any()).Return(true, nil)
				err := s.keeper.CreateJob(s.ctx, types.Job{Id: 111, CommitteeId: "1", ChainId: "1", Signs: [][]byte{
					[]byte("sign1"),
					[]byte("sign2"),
					[]byte("sign3"),
				}})
				return uint64(111), err
			},
			found: true,
		},
		"not found": {
			preRun: func() (uint64, error) {
				return uint64(0), nil
			},
			found: false,
		},
	}
	for testName, tc := range cases {
		s.Run(testName, func() {
			id, err := tc.preRun()
			s.Require().NoError(err)
			_, found, err := s.keeper.GetJobByID(s.ctx, id)
			s.Require().NoError(err)
			s.Require().Equal(tc.found, found)
		})
	}
}
