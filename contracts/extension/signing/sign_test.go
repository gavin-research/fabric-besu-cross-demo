package signing_test

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/datachainlab/fabric-besu-cross-demo/contracts/extension/signing"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	// Setup
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	cfg := authtx.NewTxConfig(cdc, []signingtypes.SignMode{signingtypes.SignMode_SIGN_MODE_DIRECT})
	txBuilder := cfg.NewTxBuilder()
	stdTx := txBuilder.GetTx()
	privKey, err := toPrivKey("713071a0b7101f177ae9c9ab0412eb7e43812bd289650f8db63f3055f2bcb029")
	assert.NoError(t, err)
	pubKey, err := crypto.DecompressPubkey(privKey.PubKey().Bytes())
	assert.NoError(t, err)
	address := crypto.PubkeyToAddress(*pubKey).Bytes()

	// Test
	sig, err := signing.Sign(stdTx, *privKey)
	assert.NoError(t, err)
	err = signing.VerifyEthSignature(stdTx, address, *sig)
	assert.NoError(t, err)
}

func toPrivKey(hexString string) (*secp256k1.PrivKey, error) {
	bz, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}
	return hd.Secp256k1.Generate()(bz).(*secp256k1.PrivKey), nil
}
