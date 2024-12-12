package keeper

import (
	"context"

	"kepler/x/horizon/types"
)

func (k msgServer) ScheduleAutomation(ctx context.Context, req *types.MsgScheduleAutomationParams) (*types.MsgScheduleAutomationResponse, error) {
	// Validating params
	reqValidateError := req.Validate()

	if reqValidateError != nil {
		return nil, reqValidateError
	}

	panic("implement me")
}
