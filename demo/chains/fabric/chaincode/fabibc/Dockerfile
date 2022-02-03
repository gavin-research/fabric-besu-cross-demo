ARG GO_VER=1.16
ARG ALPINE_VER=3.13

FROM golang:${GO_VER}-alpine${ALPINE_VER} as builder
RUN apk add --no-cache gcc libc-dev

WORKDIR /root
COPY ./ ./

RUN go build -v -mod=vendor -o /root/build/chaincode .

FROM alpine:${ALPINE_VER}

COPY ./entrypoint.sh /entrypoint.sh
COPY --from=builder /root/build/chaincode /usr/bin/chaincode

WORKDIR /root

CMD ["/entrypoint.sh"]
