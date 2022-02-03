//SPDX-License-Identifier: Unlicense
pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import "./utils/B2A.sol";
import "@datachainlab/cross-solidity/contracts/core/IContractModule.sol";
import "solidity-rlp/contracts/RLPReader.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "./ERC721Token.sol";

contract ERC721ContractModule is IContractModule, AccessControl {
    constructor() public {
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    using RLPReader for RLPReader.RLPItem;
    using RLPReader for RLPReader.Iterator;
    using RLPReader for bytes;

    bytes32 public constant CALLABLE_ROLE = keccak256("CALLABLE_ROLE");
    bytes32 public constant TYPE_URL = keccak256(abi.encodePacked("/extension.types.BesuAuthExtension"));

    struct CallInfo {
        ERC721Token erc721;
        string functionName;
        RLPReader.RLPItem args;
    }

    modifier checkRlpFormat(bytes memory raw) {
        RLPReader.RLPItem[] memory items = raw.toRlpItem().toList();
        require(items.length >= 2, "items must be more than two");
        _;
    }

    modifier onlyCaller() {
        require(hasRole(CALLABLE_ROLE, msg.sender), "msg.sender must be the caller");
        _;
    }

    function setCaller(address caller) external {
        grantRole(CALLABLE_ROLE, caller);
    }

    function onContractCall(CrossContext calldata context, bytes calldata callInfo)
        external
        override
        onlyCaller
        returns (bytes memory)
    {
        require(context.signers.length == 1, "signers length must be 1");
        require(context.signers[0].auth_type.mode == AuthType.AuthMode.AUTH_MODE_EXTENSION, "auth mode must be EXTENSION");
        require(keccak256(abi.encodePacked(context.signers[0].auth_type.option.type_url)) == TYPE_URL, "extension must be BesuAuthExtension");

        CallInfo memory converted = convert(callInfo);
        if (keccak256(bytes(converted.functionName)) == keccak256(bytes("transferFrom"))) {
            RLPReader.RLPItem[] memory args = converted.args.toList();
            require(args.length == 2, "invalid length of args");
            address from = B2A.convert(context.signers[0].id);
            address to = args[0].toAddress();
            uint256 amount = args[1].toUint();
            converted.erc721.transferFrom(from, to, amount);
            return bytes("call succeed");
        } else {
            revert("invalid callInfo");
        }
    }

    function convert(bytes memory raw) private checkRlpFormat(raw) returns (CallInfo memory callInfo) {
        RLPReader.Iterator memory iter = raw.toRlpItem().iterator();
        callInfo.erc721 = ERC721Token(iter.next().toAddress());
        callInfo.functionName = string(iter.next().toBytes());
        if (iter.hasNext()) {
            callInfo.args = iter.next();
        }
    }
}
