package main

import (
	"log"

	ethereum "github.com/hyperledger-labs/yui-relayer/chains/ethereum/module"
	fabric "github.com/hyperledger-labs/yui-relayer/chains/fabric/module"
	tendermint "github.com/hyperledger-labs/yui-relayer/chains/tendermint/module"
	"github.com/hyperledger-labs/yui-relayer/cmd"

	proxymodule "github.com/datachainlab/ibc-proxy-relay/pkg/proxy/module"
	proxytmmodule "github.com/datachainlab/ibc-proxy-relay/pkg/proxy/tendermint/module"
	"github.com/datachainlab/ibc-proxy-solidity/modules/relay/ethmultisig"
	ethProver "github.com/datachainlab/ibc-trusted-ethereum-client/pkg/ethereum"
)

func main() {
	if err := cmd.Execute(
		fabric.Module{},
		ethereum.Module{},
		ethProver.Module{},
		ethmultisig.Module{},
		tendermint.Module{},
		proxymodule.Module{},
		proxytmmodule.Module{},
	); err != nil {
		log.Fatal(err)
	}
}
