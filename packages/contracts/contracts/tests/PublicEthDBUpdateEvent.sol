// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "../EthDBUpdater.sol";

contract PublicVariableEthDBUpdateUintEvent is EthDBUpdater {
    uint256 public valueA = 1;
    uint256 public valueB = 10;

    function increment() public {
        valueA += 1;
        valueB += 2;

        _emitUpdate();
    }
}

contract PublicVariableEthDBUpdateIntEvent is EthDBUpdater {
    int256 public valueA = -100;
    int256 public valueB = 10;

    function increment() public {
        valueA += 1;
        valueB += 2;

        _emitUpdate();
    }
}

contract PublicVariableEthDBUpdateBoolEvent is EthDBUpdater {
    bool public valueA = true;
    bool public valueB = false;

    function toggle() public {
        valueA = !valueA;
        valueB = !valueB;

        _emitUpdate();
    }
}

contract PublicVariableEthDBUpdateStringEvent is EthDBUpdater {
    string public valueA = "Hello";
    string public valueB = "World";

    function update() public {
        string memory newValueA = string.concat(valueA, "!");
        string memory newValueB = string.concat(valueB, "!");

        valueA = newValueA;
        valueB = newValueB;

        _emitUpdate();
    }
}

contract PublicVariableEthDBUpdateBytesEvent is EthDBUpdater {
    bytes public valueA = "Hello";
    bytes public valueB = "World";

    function update() public {
        bytes memory newValueA = abi.encodePacked(valueA, "!");
        bytes memory newValueB = abi.encodePacked(valueB, "!");

        valueA = newValueA;
        valueB = newValueB;

        _emitUpdate();
    }
}

contract PublicVariableEthDBUpdateMultipleEvent is EthDBUpdater {
    uint256 public valueA = 1;
    uint256 public valueB = 10;

    function increment() public {
        valueA += 1;
        valueB += 2;

        _emitUpdate();

        valueA += 1;
        valueB += 2;

        _emitUpdate();
    }
}

contract PublicVariableEthDBUpdateMappingEvent is EthDBUpdater {
    mapping(uint256 => uint256) public mappingA;
    mapping(uint256 => uint256) public mappingB;

    constructor() {
        mappingA[1] = 1;
        mappingA[2] = 2;
        mappingB[1] = 3;
        mappingB[2] = 4;
    }
    
    function increment() public {
        mappingA[1] = mappingA[1] + 1;
        mappingA[2] = mappingA[2] + 2;
        mappingB[1] = mappingB[1] + 3;
        mappingB[2] = mappingB[2] + 4;

        _emitUpdate();
    }
}

contract PublicVariableEthDBUpdateMappingInitiallyEmptyEvent is EthDBUpdater {
    mapping(uint256 => uint256) public mappingA;

    constructor() {
    }
    
    function increment() public {
        mappingA[1] = mappingA[1] + 1;
        mappingA[2] = mappingA[2] + 2;

        _emitUpdate();
    }
}

contract PublicVariableEthDBUpdateArrayEvent is EthDBUpdater {
    uint256[] public arrayA;
    uint256[] public arrayB;

    constructor() {
        arrayA.push(1);
        arrayB.push(0);
        arrayB.push(0);
    }

    function addToArray() public {
        arrayA.push(arrayA.length + 1);
        arrayB.push(arrayB.length + 2);

        _emitUpdate();
    }
}

contract PublicVariableEthDBUpdateStructEvent is EthDBUpdater {
    struct StructA {
        uint256 valueA;
        uint256 valueB;
    }

    StructA public structA;
    StructA public structB;

    function increment() public {
        structA.valueA += 1;
        structA.valueB += 2;
        structB.valueA += 3;
        structB.valueB += 4;

        _emitUpdate();
    }
}
