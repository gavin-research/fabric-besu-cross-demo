package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/datachainlab/cross/x/core/auth/types"
	extsigning "github.com/datachainlab/fabric-besu-cross-demo/contracts/extension/signing"
)

var _ authtypes.AuthExtensionVerifier = (*BesuAuthExtension)(nil)

func (BesuAuthExtension) Verify(ctx sdk.Context, signer authtypes.Account, signature signing.SignatureV2, tx sdk.Tx) error {
	return extsigning.VerifyEthSignature(tx, signer.Id, signature)
}
