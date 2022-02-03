package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/datachainlab/fabric-besu-cross-demo/contracts/erc20/erc20mgr/types"
)

// InitGenesis initializes the bridge module state and binds to PortID.
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	err := k.SetAdmin(ctx, types.ID(state.Params.Admin))
	if err != nil {
		panic(err)
	}
}

// ExportGenesis exports bridge module's portID and denom trace info into its genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{}
}
