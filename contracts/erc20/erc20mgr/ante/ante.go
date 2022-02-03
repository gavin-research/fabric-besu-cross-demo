package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/datachainlab/cross/x/core/auth/types"
	"github.com/datachainlab/fabric-besu-cross-demo/contracts/erc20/erc20mgr/types"
	"github.com/pkg/errors"
)

type FabricVerificationDecorator struct{}

func NewFabricVerificationDecorator() FabricVerificationDecorator {
	return FabricVerificationDecorator{}
}

func (dec FabricVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	// no need to verify signatures on recheck tx
	if ctx.IsReCheckTx() {
		return next(ctx, tx, simulate)
	}

	for _, msg := range tx.GetMsgs() {
		cctMsg, ok := msg.(*types.MsgContractCallTx)
		if !ok {
			continue
		}
		if len(cctMsg.Signers) != 1 {
			return ctx, errors.Errorf("the number of signers must be 1, msg = %v", msg.String())
		}

		account := authtypes.Account{Id: cctMsg.Signers[0]}
		authExt := types.FabricAuthExtension{}
		err = authExt.Verify(ctx, account, signing.SignatureV2{}, tx)
		if err != nil {
			return ctx, errors.Wrap(err, "verification error")
		}
	}

	return next(ctx, tx, simulate)
}
