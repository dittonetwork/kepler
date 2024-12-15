package keeper

import (
	"kepler/x/horizon/types/state"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
)

var StateSchema = &ormv1alpha1.ModuleSchemaDescriptor{
	SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{
			Id:            1,
			ProtoFileName: state.File_x_horizon_types_state_state_proto.Path(),
		},
	},
}
