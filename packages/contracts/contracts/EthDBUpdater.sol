// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "./EtherDatabaseLib.sol";

interface IEthDBUpdater {
    event EthDBPathUpdate(
        string[] path,
        bytes data,
        uint8 dataType
    );

    event EthDBUpdate();
}

abstract contract EthDBUpdater is IEthDBUpdater {
    // Base state change function
    function _emitStateChange(
        string[] memory path,
        bytes memory data,
        EtherDatabaseLib.DataType dataType
    ) internal virtual {
        emit EthDBPathUpdate(
            path,
            data,
            uint8(dataType)
        );
    }

    function _emitStateChange(
        string memory path,
        bytes memory data,
        EtherDatabaseLib.DataType dataType
    ) internal virtual {
        string[] memory pathArray = new string[](1);
        pathArray[0] = path;
        emit EthDBPathUpdate(
            pathArray,
            data,
            uint8(dataType)
        );
    }

    // Generic upsert helper
    function _emitUpsert(
        string[] memory path,
        bytes memory data,
        EtherDatabaseLib.DataType dataType
    ) internal virtual {
        _emitStateChange(path, data, dataType);
    }

    function _emitUpsert(
        string memory path,
        bytes memory data,
        EtherDatabaseLib.DataType dataType
    ) internal virtual {
        _emitStateChange(path, data, dataType);
    }

    // Type-specific upsert helpers
    function _emitUpsertString(
        string memory path,
        string memory value
    ) internal virtual {
        _emitStateChange(path, abi.encode(value), EtherDatabaseLib.DataType.STRING);
    }

    function _emitUpsertUint256(
        string memory path,
        uint256 value
    ) internal virtual {
        _emitStateChange(path, abi.encode(value), EtherDatabaseLib.DataType.UINT256);
    }

    function _emitUpsertInt256(
        string memory path,
        int256 value
    ) internal virtual {
        _emitStateChange(path, abi.encode(value), EtherDatabaseLib.DataType.INT256);
    }

    function _emitUpsertBool(
        string memory path,
        bool value
    ) internal virtual {
        _emitStateChange(path, abi.encode(value), EtherDatabaseLib.DataType.BOOL);
    }

    function _emitUpsertBytes(
        string memory path,
        bytes memory value
    ) internal virtual {
        _emitStateChange(path, value, EtherDatabaseLib.DataType.BYTES);
    }

    // Delete helper
    function _emitDelete(string memory path) internal virtual {
        _emitStateChange(path, "", EtherDatabaseLib.DataType.NONE);
    }

    function _emitUpdate() internal virtual {
        emit EthDBUpdate();
    }
} 