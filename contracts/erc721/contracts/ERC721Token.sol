//SPDX-License-Identifier: Unlicense
pragma solidity ^0.6.8;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract ERC721Token is ERC721 {
    constructor() public ERC721("Token", "TOK") {
        /* no-op */
    }

    function mint(address receiver, uint256 itemID) public {
        _mint(receiver, itemID);
    }
}
