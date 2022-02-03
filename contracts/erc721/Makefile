GETH_VERSION ?= v1.10.4
ABIGEN ?= docker run -it --rm -v $(CURDIR)/abigenBindings:/abigenBindings -v $(CURDIR)/../../cmds/erc721/erc721/contract:/contract ethereum/client-go:alltools-$(GETH_VERSION) abigen

.PHONYL: install
install:
	npm install

.PHONY: truffle-compile
truffle-compile:
	npm run compile

.PHONY: truffle-gen-abi
truffle-gen-abi: truffle-compile
	npm run abigen

.PHONY: truffle-migrate
truffle-migrate: truffle-compile
	npm run migrate

.PHONY: truffle-migrate-demo
truffle-migrate-demo: truffle-compile
	npm run migrate:demo

.PHONY: gen-go-bindings
gen-go-bindings: truffle-gen-abi
	$(ABIGEN) --bin=/abigenBindings/bin/ERC721ContractModule.bin --abi=/abigenBindings/abi/ERC721ContractModule.abi --type ERC721ContractModule --pkg=contract --out=/contract/ERC721ContractModule.go
	$(ABIGEN) --bin=/abigenBindings/bin/ERC721Token.bin --abi=/abigenBindings/abi/ERC721Token.abi --type ERC721Token --pkg=contract --out=/contract/ERC721Token.go

.PHONY: ganache-cli
ganache-cli:
	npm run ganache-cli

.PHONY: test
test:
	npm run test

.PHONY: lint
lint:
	npm run lint
