package keeper

import (
	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"kepler/x/horizon/types"
)

var StateSchema = &ormv1alpha1.ModuleSchemaDescriptor{
	SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{
			Id:            1,
			ProtoFileName: types.File_x_horizon_types_state_proto.Path(),
		},
	},
}
