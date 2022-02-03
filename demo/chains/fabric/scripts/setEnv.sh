#!/usr/bin/env bash

set -eu

setGlobals() {
  USING_ORG=$1
  echo "Using organization ${USING_ORG}"
  if [[ ${USING_ORG} = "Platformer" ]]; then
    export CORE_PEER_ID=peer0.platformer.fabric-besu-cross-demo.com
    export CORE_PEER_LOCALMSPID=PlatformerMSP
    export CORE_PEER_ADDRESS=localhost:7051
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/platformer.fabric-besu-cross-demo.com/users/Admin@platformer.fabric-besu-cross-demo.com/msp
  elif [[ ${USING_ORG} = "OrgA" ]]; then
    export CORE_PEER_ID=peer0.orga.fabric-besu-cross-demo.com
    export CORE_PEER_LOCALMSPID=OrgAMSP
    export CORE_PEER_ADDRESS=localhost:8051
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/orga.fabric-besu-cross-demo.com/users/Admin@orga.fabric-besu-cross-demo.com/msp
  elif [[ ${USING_ORG} = "OrgB" ]]; then
    export CORE_PEER_ID=peer0.orgb.fabric-besu-cross-demo.com
    export CORE_PEER_LOCALMSPID=OrgBMSP
    export CORE_PEER_ADDRESS=localhost:9051
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/orgb.fabric-besu-cross-demo.com/users/Admin@orgb.fabric-besu-cross-demo.com/msp
  else
    echo "================== ERROR !!! ORG Unknown =================="
    exit 1
  fi
}

