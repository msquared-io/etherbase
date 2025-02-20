// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "./EtherDatabaseLib.sol";

contract EtherDatabaseImpl is EtherDatabaseLib {
    constructor(address validator) {
        init(validator);
    }
}