module github.com/datachainlab/fabric-besu-cross-demo/relayer

go 1.16

require (
	github.com/datachainlab/ibc-proxy-relay v0.0.0-20211221021543-b747411ec852
	github.com/datachainlab/ibc-proxy-solidity v0.0.0-20211221024153-892f76a1e9f5
	github.com/datachainlab/ibc-trusted-ethereum-client v0.0.0-20211216095521-4254ba045d2d
	github.com/hyperledger-labs/yui-relayer v0.1.1-0.20211201082514-122526148f85
)

replace (
	github.com/cosmos/ibc-go => github.com/datachainlab/ibc-go v0.0.0-20210623043207-6582d8c965f8
	github.com/go-kit/kit => github.com/go-kit/kit v0.8.0
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/hyperledger/fabric-sdk-go => github.com/datachainlab/fabric-sdk-go v0.0.0-20200730074927-282a61dcd92e
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
)
