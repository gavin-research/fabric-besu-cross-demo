#!/bin/bash
set -eu pipefail

# import utils
. ./scripts/setEnv.sh

OREDERER_ENDPOINT=localhost:7050

createChannel() {
    CHANNEL_NAME=$1
    ORG_NAME=$2

    setGlobals "${ORG_NAME}"
    echo "### Creating channel ${CHANNEL_NAME}"
    peer channel create -c ${CHANNEL_NAME} -f ./artifacts/${CHANNEL_NAME}.tx -o ${OREDERER_ENDPOINT}
    mv ./${CHANNEL_NAME}.block ./artifacts/${CHANNEL_NAME}.block
    echo "### Creating channel ${CHANNEL_NAME} Done"
}

joinChannel() {
    CHANNEL_NAME=$1
    ORG_NAME=$2

    setGlobals "${ORG_NAME}"
    echo "### Join channel ${CHANNEL_NAME} Org=${ORG_NAME}"
    peer channel join -b ./artifacts/${CHANNEL_NAME}.block -o ${OREDERER_ENDPOINT}
    echo "### Join channel ${CHANNEL_NAME} Org=${ORG_NAME} Done"
}

updateAnchorPeers() {
    CHANNEL_NAME=$1
    ORG_NAME=$2

    setGlobals "${ORG_NAME}"
    peer channel update -o ${OREDERER_ENDPOINT} -c ${CHANNEL_NAME} -f ./artifacts/"${CHANNEL_NAME}"-"${ORG_NAME}"Anchors.tx
}

createChannel erc20-channel Platformer

joinChannel erc20-channel Platformer
updateAnchorPeers erc20-channel Platformer

joinChannel erc20-channel OrgA
updateAnchorPeers erc20-channel OrgA

joinChannel erc20-channel OrgB
updateAnchorPeers erc20-channel OrgB

exit 0
