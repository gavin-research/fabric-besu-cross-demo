package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/datachainlab/cross/x/core/auth/types"
	fabauthtypes "github.com/hyperledger-labs/yui-fabric-ibc/x/auth/types"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

var _ authtypes.AuthExtensionVerifier = (*FabricAuthExtension)(nil)

func (FabricAuthExtension) Verify(ctx sdk.Context, signer authtypes.Account, signature signing.SignatureV2, tx sdk.Tx) error {
	client, err := getCreatorClient(ctx)
	if err != nil {
		return err
	}

	mspid, err := client.GetMSPID()
	if err != nil {
		return err
	}
	clientid, err := client.GetID()
	if err != nil {
		return err
	}
	id, err := NewID(MSPID(mspid), ClientID(clientid))
	if err != nil {
		return err
	}

	signerID := IDFromAccount(signer.Id)
	if id != signerID {
		return fmt.Errorf("got unexpected signer: expected=%v actual=%v", id, signerID)
	}
	return nil
}

func getCreatorClient(ctx sdk.Context) (*cid.ClientID, error) {
	stub := fabauthtypes.StubFromContext(ctx)
	return cid.New(stub)
}
