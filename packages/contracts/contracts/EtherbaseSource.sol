// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "./RoleControl.sol";
import "./EtherDatabaseLib.sol";
error EventNameAlreadyRegistered();
error EventNameNotRegistered();
error TooManyIndexedargs();
error TooManyTopics();
error IncorrectNumberOfTopics();

struct Argument {
    string name;
    string argType;
    bool isIndexed;
}

struct EventSchema {
    string name;
    string id;
    Argument[] args;
    bytes32 eventTopic;
    uint8 numIndexedArgs;
}

contract EtherbaseSource is RoleControl, EtherDatabaseLib {
    mapping(string => EventSchema) public eventSchemas;
    mapping(string => string) public idToName;
    string[] private registeredEventNames;

    event SchemaRegistered(string indexed name, string id);

    constructor(address _owner, address validator) {
        initialize(_owner);
        init(validator);
    }

    function registerEventSchema(
        string calldata name,
        string calldata id,
        bytes32 eventTopic,
        Argument[] calldata args
    ) external auth(Role.REGISTRAR) {
        if (bytes(eventSchemas[name].name).length != 0) {
            revert EventNameAlreadyRegistered();
        }
        uint8 numIndexedArgs = 0;
        for (uint256 i = 0; i < args.length; i++) {
            if (args[i].isIndexed) {
                numIndexedArgs++;
            }
        }
        if (numIndexedArgs > 3) {
            revert TooManyIndexedargs();
        }
        EventSchema storage schema = eventSchemas[name];
        schema.name = name;
        schema.id = id;
        schema.eventTopic = eventTopic;
        for (uint256 i = 0; i < args.length; i++) {
            schema.args.push(args[i]);
        }
        schema.numIndexedArgs = numIndexedArgs;
        registeredEventNames.push(name);
        idToName[id] = name;
        emit SchemaRegistered(name, id);
    }

    function getRegisteredEventNames() external view returns (string[] memory) {
        return registeredEventNames;
    }

    function getEventSchemaFromName(
        string calldata name
    ) external view returns (EventSchema memory) {
        if (bytes(eventSchemas[name].name).length == 0) {
            revert EventNameNotRegistered();
        }
        return eventSchemas[name];
    }

    function getEventSchemaFromId(
        string calldata id
    ) external view returns (EventSchema memory) {
        string memory name = idToName[id];
        if (bytes(name).length == 0) {
            revert EventNameNotRegistered();
        }
        return eventSchemas[name];
    }

    function _emitEvent(
        string calldata name,
        bytes32[] calldata ArgumentTopics,
        bytes memory data
    ) internal returns (bool) {
        if (bytes(eventSchemas[name].name).length == 0) {
            revert EventNameNotRegistered();
        }

        EventSchema storage schema = eventSchemas[name];
        bytes32 eventTopic = schema.eventTopic;

        if (ArgumentTopics.length > 3) {
            revert TooManyTopics();
        }

        if (ArgumentTopics.length != schema.numIndexedArgs) {
            revert IncorrectNumberOfTopics();
        }

        uint256 dataLength = data.length;

        assembly {
            let dataPtr := add(data, 32)

            switch ArgumentTopics.length 
            case 0 {
                log1(dataPtr, dataLength, eventTopic)
            }
            case 1 {
                log2(
                    dataPtr,
                    dataLength,
                    eventTopic,
                    calldataload(ArgumentTopics.offset)
                )
            }
            case 2 {
                log3(
                    dataPtr,
                    dataLength,
                    eventTopic,
                    calldataload(ArgumentTopics.offset),
                    calldataload(add(ArgumentTopics.offset, 0x20))
                )
            }
            case 3 {
                log4(
                    dataPtr,
                    dataLength,
                    eventTopic,
                    calldataload(ArgumentTopics.offset),
                    calldataload(add(ArgumentTopics.offset, 0x20)),
                    calldataload(add(ArgumentTopics.offset, 0x40))
                )
            }
        }

        return true;
    }

    function emitEvent(
        string calldata name,
        bytes32[] calldata ArgumentTopics,
        bytes memory data
    ) external auth(Role.EMITTER) returns (bool success) {
        return _emitEvent(name, ArgumentTopics, data);
    }

    struct BatchEmitEvent {
        string name;
        bytes32[] argumentTopics;
        bytes data;
    }

    /**
     * @notice Emit multiple events in a single transaction
     * @param events Array of BatchEmitEvent structs containing event details
     * @return success True if all events were emitted successfully
     */
    function emitEvents(
        BatchEmitEvent[] calldata events
    ) external auth(Role.EMITTER) returns (bool success) {
        for (uint i = 0; i < events.length; i++) {
            BatchEmitEvent calldata evt = events[i];
            _emitEvent(evt.name, evt.argumentTopics, evt.data);
        }
        return true;
    }
}
