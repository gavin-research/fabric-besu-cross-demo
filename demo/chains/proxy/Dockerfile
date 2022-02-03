ARG GOLAGN_VER=1.15.8
ARG ALPINE_VER=3.13

FROM golang:${GOLAGN_VER}-alpine${ALPINE_VER} AS builder
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

WORKDIR /root

COPY ./demo/chains/proxy ./demo/chains/proxy

WORKDIR /root/demo/chains/proxy
RUN go build -v -mod=readonly -o /root/build/simd ./simapp/simd

FROM alpine:${ALPINE_VER} AS initializer

RUN apk add --no-cache jq bash libc6-compat

WORKDIR /root

COPY --from=builder /root/build/simd /usr/bin/simd
COPY ./demo/chains/proxy/scripts ./scripts

ARG CHAINID
ARG CHAINDIR=./data
ARG RPCPORT=26657
ARG P2PPORT=26656
ARG PROFPORT=6060
ARG GRPCPORT=9090

RUN ./scripts/tm-chain /usr/bin/simd $CHAINID $CHAINDIR $RPCPORT $P2PPORT $PROFPORT $GRPCPORT

FROM alpine:${ALPINE_VER}

RUN apk add --no-cache libc6-compat

WORKDIR /root

ARG CHAINID
ARG CHAINDIR=./data
ARG RPCPORT=26657
ARG P2PPORT=26656
ARG PROFPORT=6060
ARG GRPCPORT=9090

ENV CHAINID=$CHAINID
ENV CHAINDIR=$CHAINDIR
ENV GRPCPORT=$GRPCPORT

COPY --from=builder /root/build/simd /usr/bin/simd
COPY --from=initializer /root/$CHAINDIR /root/$CHAINDIR
COPY ./demo/chains/proxy/scripts/entrypoint.sh /root/entrypoint.sh

EXPOSE $RPCPORT $P2PPORT $PROFPORT $GRPCPORT

ENTRYPOINT ["./entrypoint.sh"]
