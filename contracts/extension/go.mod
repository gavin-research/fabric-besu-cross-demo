module github.com/datachainlab/fabric-besu-cross-demo/contracts/extension

go 1.16

require (
	github.com/VividCortex/gohistogram v1.0.0 // indirect
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/cosmos/cosmos-sdk v0.43.0-beta1
	github.com/datachainlab/cross v0.2.2
	github.com/ethereum/go-ethereum v1.9.25
	github.com/gogo/protobuf v1.3.3
	github.com/stretchr/testify v1.7.0
	google.golang.org/genproto v0.0.0-20210114201628-6edceaf6022f
)

replace (
	// github.com/cosmos/cosmos-sdk => github.com/datachainlab/cosmos-sdk v0.34.4-0.20210312160443-1cdc372f98d0
	github.com/go-kit/kit => github.com/go-kit/kit v0.8.0
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
)
