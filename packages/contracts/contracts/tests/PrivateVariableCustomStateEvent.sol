// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract PrivateVariableCustomStateEvent {
    uint256 private value;

    event ValueChanged(uint256 newValue);

    function increment() public {
        value += 1;
        emit ValueChanged(value);
    }

    // todo allow specifying this function in the interface
    function getValue() public view returns (uint256) {
        return value;
    }
}
