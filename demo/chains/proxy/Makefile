include ../../var.mk

.PHONY: build
build:
	go build -o ./build/simd ./simapp/simd

.PHONY: network
network:
	$(DOCKER_COMPOSE) up -d --build

.PHONY: network-down
network-down:
	$(DOCKER_COMPOSE) down -v --remove-orphans
