// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

error PathNotFound();
error InvalidDataType(EtherDatabaseLib.DataType expected, EtherDatabaseLib.DataType actual);
error InvalidDataEncoding(EtherDatabaseLib.DataType dataType);

contract DataValidator {
    function validateDataEncoding(bytes memory data, EtherDatabaseLib.DataType dataType) external pure {
        if (dataType == EtherDatabaseLib.DataType.STRING) {
            abi.decode(data, (string));
        } else if (dataType == EtherDatabaseLib.DataType.UINT256) {
            abi.decode(data, (uint256));
        } else if (dataType == EtherDatabaseLib.DataType.INT256) {
            abi.decode(data, (int256));
        } else if (dataType == EtherDatabaseLib.DataType.BOOL) {
            abi.decode(data, (bool));
        }
    }
}

abstract contract EtherDatabaseLib {
    // -------------------------------------------------------------------------
    // TYPES, STRUCTS and ENUMS
    // -------------------------------------------------------------------------

    enum DataType {
        NONE,
        STRING,
        BOOL,
        UINT256,
        INT256,
        BYTES
    }

    struct Node {
        bytes path;
        DataType dataType;
        bytes data;
        bool exists;
    }

    // State variables
    address public validator;
    uint256 private nextSegmentId;
    
    // Maps a segment string to its unique ID
    mapping(string => uint256) private segmentId;
    
    string[] private segmentNames;
    
    // Indicates whether a segment ID is valid
    mapping(uint256 => bool) private segmentExists;

    // Store nodes in array with index mapping
    Node[] private store;
    mapping(bytes => uint256) private storeIndex;
    mapping(bytes => bool) private storeExists;
    
    // -------------------------------------------------------------------------
    // INITIALIZATION
    // -------------------------------------------------------------------------
    /**
     * @notice Initializes the database with validator address
     */
    function init(address validatorAddress) internal {
        nextSegmentId = 1;
        validator = validatorAddress;
    }

    // -------------------------------------------------------------------------
    // EXTERNAL INTERFACE FUNCTIONS
    // -------------------------------------------------------------------------

    function encodePath(string[] calldata segments) internal returns (bytes memory) {
        return _encodePath(segments);
    }

    /**
     * @notice Store or update a value in the DB, specifying the path as an array of strings.
     */
    function setValue(
        string[] calldata segments,
        DataType dataType,
        bytes calldata data
    ) external {
        _setValueInternal(segments, dataType, data);
    }

    // Internal helper for setting values
    function _setValueInternal(
        string[] calldata segments,
        DataType dataType,
        bytes memory encodedData
    ) internal {
        bytes memory path = _encodePath(segments);
        if (dataType == DataType.NONE) {
            // If type is NONE, treat it as a removal
            if (storeExists[path]) {
                uint256 index = storeIndex[path];
                store[index].exists = false;
                delete storeIndex[path];
                delete storeExists[path];
            }
            emit EthDBPathUpdate(segments, bytes(""), uint8(DataType.NONE));
            return;
        }

        if (storeExists[path]) {
            // Update existing entry
            uint256 index = storeIndex[path];
            Node storage node = store[index];
            node.data = encodedData;
            node.dataType = dataType;
        } else {
            // New entry: push and set mapping
            Node memory newNode = Node({
                path: path,
                dataType: dataType,
                data: encodedData,
                exists: true
            });
            store.push(newNode);
            storeIndex[path] = store.length - 1;
            storeExists[path] = true;
        }

        emit EthDBPathUpdate(segments, encodedData, uint8(dataType));
    }


    /**
     * @notice Retrieve an entry by specifying its string segments.
     * @return The dataType and raw data bytes. Returns (NONE, "") if not found.
     */
    function getValue(
        string[] calldata segments
    ) external view returns (DataType, bytes memory) {
        bytes memory path = _encodePathView(segments);
        if (!storeExists[path]) {
            return (DataType.NONE, "");
        }
        // Pass DataType.NONE to skip type checking
        bytes memory data = _getValueInternal(segments, DataType.NONE);
        return (store[storeIndex[path]].dataType, data);
    }

    /**
     * @notice Removes the entry at the specified path.
     */
    function removeValue(
        string[] calldata segments
    ) external {
        bytes memory path = _encodePathView(segments);
        if (!storeExists[path]) {
            revert PathNotFound();
        }
        uint256 index = storeIndex[path];
        
        // Mark as deleted by setting exists to false
        store[index].exists = false;
        delete storeIndex[path];
        delete storeExists[path];

        emit EthDBPathUpdate(segments, bytes(""), uint8(DataType.NONE));
    }

    /**
     * @notice Get the segment ID for a given segment string.
     */
    function getSegmentId(
        string calldata segment
    ) external view returns (uint256) {
        return segmentId[segment];
    }

    /**
     * @notice Get all stored segments.
     */
    function getAllSegments() external view returns (string[] memory) {
        return segmentNames;
    }
    
    /**
     * @notice Get all stored Nodes.
     */
    function getEntries() external view returns (Node[] memory) {
        uint256 validCount = 0;
        
        // Count valid entries first
        for (uint i = 0; i < store.length; i++) {
            if (store[i].exists) {
                validCount++;
            }
        }

        // Create result array with correct size
        Node[] memory result = new Node[](validCount);
        uint256 resultIndex = 0;
        
        // Copy valid entries
        for (uint i = 0; i < store.length; i++) {
            if (store[i].exists) {
                result[resultIndex] = store[i];
                resultIndex++;
            }
        }
        
        return result;
    }

    function hasEntry(string[] calldata segments) external view returns (bool) {
        bytes memory path = _encodePathView(segments);
        return storeExists[path];
    }

    // -------------------------------------------------------------------------
    // INTERNAL HELPER FUNCTIONS
    // -------------------------------------------------------------------------

    /**
     * @dev For each string segment, get or create its segmentId then varint-encode it.
     */
    function _encodePath(
        string[] calldata segments
    ) internal returns (bytes memory) {
        bytes memory buf = new bytes(0);
        for (uint i = 0; i < segments.length; i++) {
            uint id = _getOrCreateSegmentId(segments[i]);
            buf = abi.encodePacked(buf, _encodeVarint(id));
        }
        // Append 0 terminator
        buf = abi.encodePacked(buf, bytes1(0x00));
        return buf;
    }

    /**
     * @dev Same as _encodePath but in view mode (no segment creation).
     */
    function _encodePathView(
        string[] calldata segments
    ) internal view returns (bytes memory) {
        bytes memory buf = new bytes(0);
        for (uint i = 0; i < segments.length; i++) {
            uint id = segmentId[segments[i]];
            buf = abi.encodePacked(buf, _encodeVarint(id));
        }
        // Append 0 terminator
        buf = abi.encodePacked(buf, bytes1(0x00));
        return buf;
    }

    /**
     * @dev Varint-encodes a positive integer.
     */
    function _encodeVarint(uint x) internal pure returns (bytes memory) {
        if (x == 0) {
            return abi.encodePacked(bytes1(0x00));
        }
        bytes memory out;
        uint temp = x;
        while (temp >= 0x80) {
            out = abi.encodePacked(out, bytes1(uint8((temp & 0x7F) | 0x80)));
            temp >>= 7;
        }
        out = abi.encodePacked(out, bytes1(uint8(temp & 0x7F)));
        return out;
    }

    function getOrCreateSegmentId(string calldata segment) external returns (uint256) {
        return _getOrCreateSegmentId(segment);
    }

    /**
     * @dev Look up the ID of a segment string; if not present, create a new one.
     */
    function _getOrCreateSegmentId(
        string calldata segment
    ) internal returns (uint256) {
        uint256 id = segmentId[segment];
        if (segmentExists[id]) {
            return id;
        } else {
            uint256 newId = nextSegmentId;
            nextSegmentId++;

            segmentId[segment] = newId;
            segmentNames.push(segment);
            segmentExists[newId] = true;

            emit NewSegment(segment, newId);

            return newId;
        }
    }

    event NewSegment(string segment, uint256 indexed id);
    
    event EthDBPathUpdate(
        string[] path,
        bytes data,
        uint8 dataType
    );

    // Internal helper for getting values
    function _getValueInternal(
        string[] calldata segments,
        DataType expectedType
    ) internal view returns (bytes memory) {
        bytes memory path = _encodePathView(segments);
        if (!storeExists[path]) {
            revert PathNotFound();
        }
        
        Node storage node = store[storeIndex[path]];
        if (!node.exists) {
            revert PathNotFound();
        }
        
        if (expectedType != DataType.NONE && node.dataType != expectedType) {
            revert InvalidDataType(expectedType, node.dataType);
        }

        // Add validation for data encoding based on type - we require the validator to be a different contract
        // as we need the try/catch to avoid getting unhelpful errors if you try to decode the wrong type or incorrectly encoded data
        if (expectedType != DataType.NONE && expectedType != DataType.BYTES) {
            try DataValidator(validator).validateDataEncoding(node.data, expectedType) {} catch {
                revert InvalidDataEncoding(expectedType);
            }
        }
        
        return node.data;
    }

    struct BatchSetValue {
        string[] segments;
        DataType dataType;
        bytes data;
    }

    /**
     * @notice Store or update multiple values in the DB in a single operation
     * @param values Array of BatchSetValue structs containing the path segments, data type, and data
     */
    function setValues(
        BatchSetValue[] calldata values
    ) external {
        for (uint i = 0; i < values.length; i++) {
            _setValueInternal(
                values[i].segments,
                values[i].dataType,
                values[i].data
            );
        }
    }
}
