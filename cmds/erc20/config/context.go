package config

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/hyperledger-labs/yui-relayer/chains/fabric"
)

type Context struct {
	Codec    codec.ProtoCodecMarshaler
	Config   *fabric.ChainConfig
	ClientID string
	Chain    *Chain
}
