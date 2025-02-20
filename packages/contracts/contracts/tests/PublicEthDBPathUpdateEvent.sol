// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "../EthDBUpdater.sol";
import "../EtherDatabaseLib.sol";
contract PublicVariableEthDBPathUpdateUintEvent is EthDBUpdater {
    uint256 public valueA = 1;
    uint256 public valueB = 10;

    function increment() public {
        valueA += 1;
        valueB += 2;

        _emitUpsertUint256("valueA", valueA);
        _emitUpsertUint256("valueB", valueB);
    }
}

contract PublicVariableEthDBUpdateIntEvent is EthDBUpdater {
    int256 public valueA = -100;
    int256 public valueB = 10;

    function increment() public {
        valueA += 1;
        valueB += 2;
    
        _emitUpsertInt256("valueA", valueA);
        _emitUpsertInt256("valueB", valueB);
    }
}

contract PublicVariableEthDBUpdateBoolEvent is EthDBUpdater {
    bool public valueA = true;
    bool public valueB = false;

    function toggle() public {
        valueA = !valueA;
        valueB = !valueB;

        _emitUpsertBool("valueA", valueA);
        _emitUpsertBool("valueB", valueB);
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

        _emitUpsertString("valueA", valueA);
        _emitUpsertString("valueB", valueB);
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

        _emitUpsertBytes("valueA", valueA);
        _emitUpsertBytes("valueB", valueB);
    }
}

contract PublicVariableEthDBPathUpdateMultipleEvent is EthDBUpdater {
    uint256 public valueA = 1;
    uint256 public valueB = 10;

    function increment() public {
        valueA += 1;
        valueB += 2;

        _emitUpsertUint256("valueA", valueA);
        _emitUpsertUint256("valueB", valueB);

        valueA += 1;
        valueB += 2;

        _emitUpsertUint256("valueA", valueA);
        _emitUpsertUint256("valueB", valueB);
    }
}

contract PublicMappingEthDBUpdateEvent is EthDBUpdater {
    mapping(uint256 => uint256) public mappingA;
    mapping(uint256 => uint256) public mappingB;

    function increment() public {
        mappingA[1] = mappingA[1] + 1;
        mappingA[2] = mappingA[2] + 2;
        mappingB[1] = mappingB[1] + 3;
        mappingB[2] = mappingB[2] + 4;

        _emitUpsertUint256("mappingA[1]", mappingA[1]);
        _emitUpsertUint256("mappingA[2]", mappingA[2]);
        _emitUpsertUint256("mappingB[1]", mappingB[1]);
        _emitUpsertUint256("mappingB[2]", mappingB[2]);
    }
}

// dont support granular array updating rn
// contract PublicArrayEthDBUpdateEvent is EthDBUpdater {
//     uint256[] public arrayA;
//     uint256[] public arrayB;

//     function addToArray() public {
//         arrayA.push(arrayA.length + 1);
//         arrayB.push(arrayB.length + 2);

//         _emitUpsertUint("arrayA[0]", arrayA[0]);
//         _emitUpsertUint("arrayB[0]", arrayB[0]);
//     }
// }

contract PublicStructEthDBUpdateEvent is EthDBUpdater {
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

        _emitUpsertUint256("structA.valueA", structA.valueA);
        _emitUpsertUint256("structA.valueB", structA.valueB);
        _emitUpsertUint256("structB.valueA", structB.valueA);
        _emitUpsertUint256("structB.valueB", structB.valueB);
    }
}
