//SPDX-License-Identifier: Unlicense
pragma solidity ^0.6.8;

library B2A {
    /* solhint-disable */
    function convert(bytes memory b) internal pure returns (address addr) {
        assembly {
            addr := mload(add(b,20))
        }
    }
    /* solhint-enable */
}
