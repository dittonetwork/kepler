package keeper_test

import (
	"encoding/hex"
	"fmt"

	"kepler/x/committee/keeper"
	"kepler/x/committee/types"
)

func (s *KeeperTestSuite) TestMsgServer_RandaoCommitReveal() {
	valAddr := "cosmos16wfryel63g7axeamw68630wglalcnk3l0zuadc"

	cases := map[string]struct {
		preRun    func() (*types.MsgCommitRandao, error)
		expErr    bool
		expErrMsg string
	}{
		"all ok": {
			preRun: func() (*types.MsgCommitRandao, error) {
				return &types.MsgCommitRandao{
					CommitmentHash:   []byte("commitmentHash"),
					Validator:        valAddr,
					ExecutionChainId: "0",
					EpochId:          1,
				}, nil
			},
		},
	}

	for name, tc := range cases {
		s.Run(name, func() {
			msg, err := tc.preRun()
			s.Require().NoError(err)
			res, err := s.msgSrvr.CommitRandao(s.ctx, msg)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(res)
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgServer_RevealRandao() {
	valAddr := "cosmos16wfryel63g7axeamw68630wglalcnk3l0zuadc"

	cases := map[string]struct {
		preRun    func() (*types.MsgRevealRandao, error)
		expErr    bool
		expErrMsg string
	}{
		"all ok": {
			preRun: func() (*types.MsgRevealRandao, error) {
				hashToCommit, err := hex.DecodeString(
					"cff3db0365d00ccd4662087c5158082fd71ec5fdd97bfaeb9e7634043521a0f2",
				)
				s.Require().NoError(err)

				_, err = s.msgSrvr.CommitRandao(s.ctx, &types.MsgCommitRandao{
					CommitmentHash:   hashToCommit,
					Validator:        valAddr,
					ExecutionChainId: "0",
					EpochId:          1,
				})
				s.Require().NoError(err)

				realSeed := "randomseed"
				seed, err := hex.DecodeString(fmt.Sprintf("%X", realSeed))
				s.Require().NoError(err)

				return &types.MsgRevealRandao{
					RandomSeed:       seed,
					Validator:        valAddr,
					ExecutionChainId: "0",
					EpochId:          2,
				}, nil
			},
		},
		"invalid commitment hash": {
			preRun: func() (*types.MsgRevealRandao, error) {
				hashToCommit, err := hex.DecodeString(
					"cff3db0365d00ccd4662087c5158082fd71ec5fdd97bfaeb9e7634043521a0f4",
				)
				s.Require().NoError(err)

				_, err = s.msgSrvr.CommitRandao(s.ctx, &types.MsgCommitRandao{
					CommitmentHash:   hashToCommit,
					Validator:        valAddr,
					ExecutionChainId: "0",
					EpochId:          2,
				})
				s.Require().NoError(err)

				realSeed := "not_real_seed"
				seed, err := hex.DecodeString(fmt.Sprintf("%X", realSeed))
				s.Require().NoError(err)

				return &types.MsgRevealRandao{
					RandomSeed:       seed,
					Validator:        valAddr,
					ExecutionChainId: "0",
					EpochId:          3,
				}, nil
			},
			expErr:    true,
			expErrMsg: keeper.ErrInvalidCommitmentHash.Error(),
		},
	}

	for name, tc := range cases {
		s.Run(name, func() {
			msg, err := tc.preRun()
			s.Require().NoError(err)
			res, err := s.msgSrvr.RevealRandao(s.ctx, msg)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(res)
			}
		})
	}
}
