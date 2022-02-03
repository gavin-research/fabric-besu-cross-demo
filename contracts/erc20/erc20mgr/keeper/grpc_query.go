package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/datachainlab/fabric-besu-cross-demo/contracts/erc20/erc20mgr/types"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ types.QueryServer = (*Keeper)(nil)

func (k Keeper) BalanceOf(ctx context.Context, req *types.QueryBalanceOfRequest) (*types.QueryBalanceOfResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	v, err := k.erc20Keeper.BalanceOf(sdkCtx, string(req.Id))
	if err != nil {
		return nil, err
	}
	return &types.QueryBalanceOfResponse{
		Balance: &types.Balance{
			Id:     req.Id,
			Amount: v,
		},
	}, nil
}

func (k Keeper) TotalSupply(ctx context.Context, req *emptypb.Empty) (*types.QueryTotalSupplyResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	v, err := k.erc20Keeper.TotalSupply(sdkCtx)
	if err != nil {
		return nil, err
	}
	return &types.QueryTotalSupplyResponse{TotalSupply: v}, nil
}

func (k Keeper) Allowance(ctx context.Context, req *types.QueryAllowanceRequest) (*types.QueryAllowanceResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	v, err := k.erc20Keeper.Allowance(sdkCtx, string(req.Owner), string(req.Spender))
	if err != nil {
		return nil, err
	}
	return &types.QueryAllowanceResponse{Amount: v}, nil
}
