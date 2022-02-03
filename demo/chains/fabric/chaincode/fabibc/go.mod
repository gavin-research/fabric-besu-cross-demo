module github.com/datachainlab/fabric-besu-cross-demo/demo/chains/chaincode/fabibc

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.43.0-beta1
	github.com/cosmos/ibc-go v1.0.0-beta1
	github.com/datachainlab/cross v0.2.2
	github.com/datachainlab/cross-cdt v0.0.0-20211216051311-b41689652356
	github.com/datachainlab/fabric-besu-cross-demo/contracts/erc20 v0.0.0
	github.com/datachainlab/ibc-mock-client v0.0.0-20210801020446-f9dbe7320294
	github.com/datachainlab/ibc-proxy v0.0.0-20211221015432-a8c24b85f7ba
	github.com/datachainlab/ibc-proxy-solidity v0.0.0-20211221024153-892f76a1e9f5
	github.com/datachainlab/ibc-trusted-ethereum-client v0.0.0-20211216095521-4254ba045d2d
	github.com/hyperledger-labs/yui-fabric-ibc v0.2.1-0.20211115073030-02f4f392c023
	github.com/hyperledger/fabric-chaincode-go v0.0.0-20200511190512-bcfeb58dd83a
	github.com/hyperledger/fabric-contract-api-go v1.0.0
	github.com/tendermint/tendermint v0.34.10
	github.com/tendermint/tm-db v0.6.4
)

replace (
	github.com/cosmos/ibc-go => github.com/datachainlab/ibc-go v0.0.0-20210903082226-209e4c3c13d4
	github.com/datachainlab/fabric-besu-cross-demo/contracts/erc20 => ./../../../../../contracts/erc20
	github.com/datachainlab/fabric-besu-cross-demo/contracts/extension => ../../../../../contracts/extension
	github.com/go-kit/kit => github.com/go-kit/kit v0.8.0
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
)
