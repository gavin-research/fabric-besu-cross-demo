package cross

import (
	"encoding/hex"
	"log"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/datachainlab/fabric-besu-cross-demo/contracts/extension/signing"
)

func hexToSecp256k1PrivKey(hexString string) (*secp256k1.PrivKey, error) {
	bz, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}

	return hd.Secp256k1.Generate()(bz).(*secp256k1.PrivKey), nil
}

func hexToEthereumAddress(hexString string) ([]byte, error) {
	if privKey, err := crypto.HexToECDSA(hexString); err != nil {
		return nil, err
	} else {
		return crypto.PubkeyToAddress(privKey.PublicKey).Bytes(), nil
	}
}

func buildTx(registry codectypes.InterfaceRegistry, signKey *secp256k1.PrivKey, msgs ...sdk.Msg) ([]byte, error) {
	log.Printf("fabric.SendMsgs: %v", msgs)

	m := codec.NewProtoCodec(registry)
	cfg := authtx.NewTxConfig(m, []signingtypes.SignMode{signingtypes.SignMode_SIGN_MODE_DIRECT})

	txBuilder := cfg.NewTxBuilder()
	if err := txBuilder.SetMsgs(msgs...); err != nil {
		return nil, err
	}

	sigV2, err := signing.Sign(txBuilder.GetTx(), *signKey)
	if err != nil {
		return nil, err
	}

	if err := txBuilder.SetSignatures(*sigV2); err != nil {
		return nil, err
	}
	tx := txBuilder.GetTx()
	return cfg.TxJSONEncoder()(tx)
}
