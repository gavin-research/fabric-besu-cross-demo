package keeper

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	authtypes "github.com/datachainlab/cross/x/core/auth/types"
	txtypes "github.com/datachainlab/cross/x/core/tx/types"
	"github.com/datachainlab/fabric-besu-cross-demo/contracts/erc20/erc20mgr/types"
	"github.com/pkg/errors"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	erc20keeper "github.com/datachainlab/cross-cdt/modules/erc20/keeper"
)

type Keeper struct {
	m           codec.Codec
	erc20Keeper erc20keeper.Keeper
	paramSpace  paramtypes.Subspace
}

// NewKeeper creates a new keeper instance
func NewKeeper(m codec.Codec, erc20Keeper erc20keeper.Keeper, paramSpace paramtypes.Subspace) Keeper {
	paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	return Keeper{m: m, erc20Keeper: erc20Keeper, paramSpace: paramSpace}
}

// HandleContractCall is called by ContractModule
func (k Keeper) HandleContractCall(goCtx context.Context, signers []authtypes.Account, callInfo txtypes.ContractCallInfo) (*txtypes.ContractCallResult, error) {
	var req types.ContractCallRequest
	if err := k.m.UnmarshalJSON(callInfo, &req); err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	if len(signers) == 0 {
		return nil, fmt.Errorf("can't get signer")
	}
	id := types.IDFromAccount(signers[0].Id)

	switch req.Method {
	case "mint":
		return k.HandleMint(ctx, id, req)
	case "burn":
		return k.HandleBurn(ctx, id, req)
	case "transfer":
		return k.HandleTransfer(ctx, id, req)
	case "approve":
		return k.HandleApprove(ctx, id, req)
	case "transferFrom":
		return k.HandleTransferFrom(ctx, id, req)
	default:
		panic(fmt.Sprintf("unknown method '%v'", req.Method))
	}
}

func (k Keeper) HandleMint(ctx sdk.Context, id types.ID, req types.ContractCallRequest) (*txtypes.ContractCallResult, error) {
	admin, err := k.getAdmin(ctx)
	if err != nil {
		return nil, err
	}
	if id != admin {
		return nil, errors.Errorf("failed to authorize. id: %s", id)
	}

	if len(req.Args) != 2 {
		return nil, errors.New("invalid argument")
	}

	account := types.ID(req.Args[0])
	amount, err := strconv.ParseInt(req.Args[1], 10, 64)
	if err != nil {
		return nil, err
	}
	err = k.erc20Keeper.Mint(ctx, string(account), amount)
	if err != nil {
		return nil, err
	}

	// TODO: set return values
	return &txtypes.ContractCallResult{Data: []byte("")}, nil
}

func (k Keeper) HandleBurn(ctx sdk.Context, id types.ID, req types.ContractCallRequest) (*txtypes.ContractCallResult, error) {
	admin, err := k.getAdmin(ctx)
	if err != nil {
		return nil, err
	}
	if id != admin {
		return nil, errors.Errorf("failed to authorize. id: %s", id)
	}

	if len(req.Args) != 2 {
		return nil, errors.New("invalid argument")
	}
	account := types.ID(req.Args[0])
	amount, err := strconv.ParseInt(req.Args[1], 10, 64)
	if err != nil {
		return nil, err
	}
	err = k.erc20Keeper.Burn(ctx, string(account), amount)
	if err != nil {
		return nil, err
	}

	// TODO: set return values
	return &txtypes.ContractCallResult{Data: []byte("")}, nil
}

func (k Keeper) HandleTransfer(ctx sdk.Context, sender types.ID, req types.ContractCallRequest) (*txtypes.ContractCallResult, error) {
	if len(req.Args) != 2 {
		return nil, errors.New("invalid argument")
	}
	recipient := types.ID(req.Args[0])
	amount, err := strconv.ParseInt(req.Args[1], 10, 64)
	if err != nil {
		return nil, err
	}
	err = k.erc20Keeper.Transfer(ctx, string(sender), string(recipient), amount)
	if err != nil {
		return nil, err
	}

	// TODO: set return values
	return &txtypes.ContractCallResult{Data: []byte("")}, nil
}

func (k Keeper) HandleApprove(ctx sdk.Context, owner types.ID, req types.ContractCallRequest) (*txtypes.ContractCallResult, error) {
	if len(req.Args) != 2 {
		return nil, errors.New("invalid argument")
	}
	spender := types.ID(req.Args[0])
	// TODO: Check amount <= balance
	amount, err := strconv.ParseInt(req.Args[1], 10, 64)
	if err != nil {
		return nil, err
	}
	err = k.erc20Keeper.Approve(ctx, string(owner), string(spender), amount)
	if err != nil {
		return nil, err
	}

	// TODO: set return values
	return &txtypes.ContractCallResult{Data: []byte("")}, nil
}

func (k Keeper) HandleTransferFrom(ctx sdk.Context, spender types.ID, req types.ContractCallRequest) (*txtypes.ContractCallResult, error) {
	if len(req.Args) != 3 {
		return nil, errors.New("invalid argument")
	}
	owner := types.ID(req.Args[0])
	recipient := types.ID(req.Args[1])
	amount, err := strconv.ParseInt(req.Args[2], 10, 64)
	if err != nil {
		return nil, err
	}
	err = k.erc20Keeper.TransferFrom(ctx, string(owner), string(spender), string(recipient), amount)
	if err != nil {
		return nil, err
	}

	// TODO: set return values
	return &txtypes.ContractCallResult{Data: []byte("")}, nil
}

func (k Keeper) Codec() codec.Codec {
	return k.m
}

func (k Keeper) SetAdmin(ctx sdk.Context, admin types.ID) error {
	var isSetAdmin bool
	k.paramSpace.GetIfExists(ctx, types.KeyIsSetAdmin, &isSetAdmin)
	if isSetAdmin {
		return errors.New("failed to set admin")
	}
	k.paramSpace.Set(ctx, types.KeyAdmin, string(admin))
	k.paramSpace.Set(ctx, types.KeyIsSetAdmin, true)

	return nil
}

func (k Keeper) getAdmin(ctx sdk.Context) (types.ID, error) {
	var admin string
	k.paramSpace.GetIfExists(ctx, types.KeyAdmin, &admin)
	if admin == "" {
		return "", errors.New("failed to get admin")
	}

	return types.ID(admin), nil
}
