.PHONY: protoc
protoc:
	@echo "Generating Protobuf files"
	docker run -v $(CURDIR):/workspace -v $(CURDIR)/../extension:/extension --workdir /workspace tendermintdev/sdk-proto-gen sh ./scripts/protocgen.sh
