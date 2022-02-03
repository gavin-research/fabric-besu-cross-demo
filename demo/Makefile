
.PHONY: build
build:
	mkdir -p ./bin
	make -C ../relayer build && cp ../relayer/build/uly ./bin/
	make -C ../cmds/erc20 build && cp ../cmds/erc20/build/erc20cli ./bin/
	make -C ../cmds/erc721 build && cp ../cmds/erc721/build/erc721cli ./bin/
	make -C ../contracts/erc721 install
	make -C ../contracts/erc721 truffle-compile
	make -C chains/fabric bin
	make -C chains/fabric docker-chaincode

.PHONY: network
network:
	make -C chains/besu network
	make -C chains/fabric network
	make -C chains/proxy network

.PHONY: network-down
network-down:
	make -C chains/besu network-down
	make -C chains/fabric network-down
	make -C chains/proxy network-down
	@./scripts/relayer-stop
	make clean

.PHONY: clean
clean:
	rm -rf .erc20cli
	rm -rf .erc721cli
	rm -rf .urelayer
	rm -rf ./data
	rm -rf ./fixtures
	make -C chains/besu clean
	make -C chains/fabric clean

.PHONY: demo
demo:
	@./scripts/init-cli
	@./scripts/init-rly
	@./scripts/handshake
	@./scripts/relayer-start > /dev/null 2>&1 &
	@./scripts/scenario/sample
	@./scripts/relayer-stop
	@./scripts/scenario/concurrent
	@./scripts/scenario/concurrent-failure
