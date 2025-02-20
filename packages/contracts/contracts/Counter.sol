// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract Counter {
    uint256 public value;

    event ValueChanged(uint256 newValue);

    function increment() public {
        value += 1;
        emit ValueChanged(value);
    }

    // function getValue() public view returns (uint256) {
    //     return value;
    // }
}
