ARG BESU_VERSION
FROM hyperledger/besu:${BESU_VERSION}

USER root

WORKDIR /tmp/besu
ADD ./networkFiles /tmp/besu/networkFiles/
ADD ./node/node_start.sh /tmp/besu/

RUN apt update && apt upgrade -y && apt install -y curl

EXPOSE $RPC_PORT $P2P_PORT

ENTRYPOINT ./node_start.sh
