module github.com/datachainlab/fabric-besu-cross-demo/demo/chains/proxy

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.43.0-beta1
	github.com/cosmos/ibc-go v1.0.0-beta1
	github.com/datachainlab/ibc-mock-client v0.0.0-20210801020446-f9dbe7320294
	github.com/datachainlab/ibc-proxy v0.0.0-20211221015432-a8c24b85f7ba
	github.com/datachainlab/ibc-trusted-ethereum-client v0.0.0-20211216095521-4254ba045d2d
	github.com/gorilla/mux v1.8.0
	github.com/hyperledger-labs/yui-fabric-ibc v0.2.1-0.20211115073030-02f4f392c023
	github.com/rakyll/statik v0.1.7
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.10
	github.com/tendermint/tm-db v0.6.4
)

replace (
	github.com/cosmos/ibc-go => github.com/datachainlab/ibc-go v0.0.0-20210623043207-6582d8c965f8
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
)
