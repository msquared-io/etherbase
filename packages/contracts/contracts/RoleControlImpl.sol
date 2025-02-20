// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "./RoleControl.sol";

contract RoleControlImpl is RoleControl {
    constructor(address _owner) {
        initialize(_owner);
    }
}
