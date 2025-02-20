// SPDX-License-Identifier: MIT

pragma solidity ^0.8.28;

contract Test {
    uint256 nextSegmentId = 1;
    mapping(string => uint256) segmentId;
    mapping(uint256 => bool) segmentExists;
    string[] segmentNames;

    function encodePath(string[] calldata segments) external returns (bytes memory) {
        return _encodePath(segments);
    }

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

        return newId;
        }
    }
}