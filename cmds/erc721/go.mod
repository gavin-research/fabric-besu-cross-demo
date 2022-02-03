module github.com/datachainlab/fabric-besu-cross-demo/cmds/erc721

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.43.0-beta1
	github.com/datachainlab/cross v0.2.2
	github.com/datachainlab/fabric-besu-cross-demo/contracts/extension v0.0.0
	github.com/ethereum/go-ethereum v1.10.4
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
)

replace (
	github.com/datachainlab/fabric-besu-cross-demo/contracts/extension => ../../contracts/extension
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
)
