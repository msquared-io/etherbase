export const EtherbaseSourceAbi = [
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_owner",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "validator",
        "type": "address"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "constructor"
  },
  {
    "inputs": [],
    "name": "EventNameAlreadyRegistered",
    "type": "error"
  },
  {
    "inputs": [],
    "name": "EventNameNotRegistered",
    "type": "error"
  },
  {
    "inputs": [],
    "name": "IdentityAlreadyExists",
    "type": "error"
  },
  {
    "inputs": [],
    "name": "IdentityDoesNotExist",
    "type": "error"
  },
  {
    "inputs": [],
    "name": "IncorrectNumberOfTopics",
    "type": "error"
  },
  {
    "inputs": [],
    "name": "TooManyIndexedargs",
    "type": "error"
  },
  {
    "inputs": [],
    "name": "Unauthorized",
    "type": "error"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "wallet",
        "type": "address"
      }
    ],
    "name": "IdentityCreated",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "string",
        "name": "segment",
        "type": "string"
      },
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "id",
        "type": "uint256"
      }
    ],
    "name": "NewSegment",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "wallet",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "enum RoleControl.Role",
        "name": "role",
        "type": "uint8"
      }
    ],
    "name": "RoleGranted",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "string",
        "name": "name",
        "type": "string"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "id",
        "type": "string"
      }
    ],
    "name": "SchemaRegistered",
    "type": "event"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "walletAddress",
        "type": "address"
      },
      {
        "internalType": "enum RoleControl.Role[]",
        "name": "initialRoles",
        "type": "uint8[]"
      }
    ],
    "name": "createWalletIdentity",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "wallet",
        "type": "address"
      }
    ],
    "name": "deleteIdentity",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string",
        "name": "name",
        "type": "string"
      },
      {
        "internalType": "bytes32[]",
        "name": "ArgumentTopics",
        "type": "bytes32[]"
      },
      {
        "internalType": "bytes",
        "name": "data",
        "type": "bytes"
      }
    ],
    "name": "emitEvent",
    "outputs": [
      {
        "internalType": "bool",
        "name": "success",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "components": [
          {
            "internalType": "string",
            "name": "name",
            "type": "string"
          },
          {
            "internalType": "bytes32[]",
            "name": "argumentTopics",
            "type": "bytes32[]"
          },
          {
            "internalType": "bytes",
            "name": "data",
            "type": "bytes"
          }
        ],
        "internalType": "struct EtherbaseSource.BatchEmitEvent[]",
        "name": "events",
        "type": "tuple[]"
      }
    ],
    "name": "emitEvents",
    "outputs": [
      {
        "internalType": "bool",
        "name": "success",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string[]",
        "name": "segments",
        "type": "string[]"
      }
    ],
    "name": "encodePath",
    "outputs": [
      {
        "internalType": "bytes",
        "name": "",
        "type": "bytes"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string",
        "name": "",
        "type": "string"
      }
    ],
    "name": "eventSchemas",
    "outputs": [
      {
        "internalType": "string",
        "name": "name",
        "type": "string"
      },
      {
        "internalType": "string",
        "name": "id",
        "type": "string"
      },
      {
        "internalType": "bytes32",
        "name": "eventTopic",
        "type": "bytes32"
      },
      {
        "internalType": "uint8",
        "name": "numIndexedArgs",
        "type": "uint8"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getAllIdentities",
    "outputs": [
      {
        "components": [
          {
            "internalType": "address",
            "name": "walletAddress",
            "type": "address"
          },
          {
            "internalType": "enum RoleControl.Role[]",
            "name": "roles",
            "type": "uint8[]"
          }
        ],
        "internalType": "struct RoleControl.IdentityView[]",
        "name": "",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getAllSegments",
    "outputs": [
      {
        "internalType": "string[]",
        "name": "",
        "type": "string[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getAllWallets",
    "outputs": [
      {
        "internalType": "address[]",
        "name": "",
        "type": "address[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string[]",
        "name": "segments",
        "type": "string[]"
      }
    ],
    "name": "getBool",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string[]",
        "name": "segments",
        "type": "string[]"
      }
    ],
    "name": "getBytes",
    "outputs": [
      {
        "internalType": "bytes",
        "name": "",
        "type": "bytes"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getEntries",
    "outputs": [
      {
        "components": [
          {
            "internalType": "bytes",
            "name": "path",
            "type": "bytes"
          },
          {
            "internalType": "enum EtherDatabaseLib.DataType",
            "name": "dataType",
            "type": "uint8"
          },
          {
            "internalType": "bytes",
            "name": "data",
            "type": "bytes"
          },
          {
            "internalType": "bool",
            "name": "exists",
            "type": "bool"
          }
        ],
        "internalType": "struct EtherDatabaseLib.Node[]",
        "name": "",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string",
        "name": "id",
        "type": "string"
      }
    ],
    "name": "getEventSchemaFromId",
    "outputs": [
      {
        "components": [
          {
            "internalType": "string",
            "name": "name",
            "type": "string"
          },
          {
            "internalType": "string",
            "name": "id",
            "type": "string"
          },
          {
            "components": [
              {
                "internalType": "string",
                "name": "name",
                "type": "string"
              },
              {
                "internalType": "string",
                "name": "argType",
                "type": "string"
              },
              {
                "internalType": "bool",
                "name": "isIndexed",
                "type": "bool"
              }
            ],
            "internalType": "struct Argument[]",
            "name": "args",
            "type": "tuple[]"
          },
          {
            "internalType": "bytes32",
            "name": "eventTopic",
            "type": "bytes32"
          },
          {
            "internalType": "uint8",
            "name": "numIndexedArgs",
            "type": "uint8"
          }
        ],
        "internalType": "struct EventSchema",
        "name": "",
        "type": "tuple"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string",
        "name": "name",
        "type": "string"
      }
    ],
    "name": "getEventSchemaFromName",
    "outputs": [
      {
        "components": [
          {
            "internalType": "string",
            "name": "name",
            "type": "string"
          },
          {
            "internalType": "string",
            "name": "id",
            "type": "string"
          },
          {
            "components": [
              {
                "internalType": "string",
                "name": "name",
                "type": "string"
              },
              {
                "internalType": "string",
                "name": "argType",
                "type": "string"
              },
              {
                "internalType": "bool",
                "name": "isIndexed",
                "type": "bool"
              }
            ],
            "internalType": "struct Argument[]",
            "name": "args",
            "type": "tuple[]"
          },
          {
            "internalType": "bytes32",
            "name": "eventTopic",
            "type": "bytes32"
          },
          {
            "internalType": "uint8",
            "name": "numIndexedArgs",
            "type": "uint8"
          }
        ],
        "internalType": "struct EventSchema",
        "name": "",
        "type": "tuple"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string[]",
        "name": "segments",
        "type": "string[]"
      }
    ],
    "name": "getInt",
    "outputs": [
      {
        "internalType": "int256",
        "name": "",
        "type": "int256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string",
        "name": "segment",
        "type": "string"
      }
    ],
    "name": "getOrCreateSegmentId",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getRegisteredEventNames",
    "outputs": [
      {
        "internalType": "string[]",
        "name": "",
        "type": "string[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string",
        "name": "segment",
        "type": "string"
      }
    ],
    "name": "getSegmentId",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string[]",
        "name": "segments",
        "type": "string[]"
      }
    ],
    "name": "getString",
    "outputs": [
      {
        "internalType": "string",
        "name": "",
        "type": "string"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string[]",
        "name": "segments",
        "type": "string[]"
      }
    ],
    "name": "getUint",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string[]",
        "name": "segments",
        "type": "string[]"
      }
    ],
    "name": "getValue",
    "outputs": [
      {
        "internalType": "enum EtherDatabaseLib.DataType",
        "name": "",
        "type": "uint8"
      },
      {
        "internalType": "bytes",
        "name": "",
        "type": "bytes"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "wallet",
        "type": "address"
      },
      {
        "internalType": "enum RoleControl.Role",
        "name": "role",
        "type": "uint8"
      }
    ],
    "name": "grantRole",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string[]",
        "name": "segments",
        "type": "string[]"
      }
    ],
    "name": "hasEntry",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string",
        "name": "",
        "type": "string"
      }
    ],
    "name": "idToName",
    "outputs": [
      {
        "internalType": "string",
        "name": "",
        "type": "string"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "owner",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string",
        "name": "name",
        "type": "string"
      },
      {
        "internalType": "string",
        "name": "id",
        "type": "string"
      },
      {
        "internalType": "bytes32",
        "name": "eventTopic",
        "type": "bytes32"
      },
      {
        "components": [
          {
            "internalType": "string",
            "name": "name",
            "type": "string"
          },
          {
            "internalType": "string",
            "name": "argType",
            "type": "string"
          },
          {
            "internalType": "bool",
            "name": "isIndexed",
            "type": "bool"
          }
        ],
        "internalType": "struct Argument[]",
        "name": "args",
        "type": "tuple[]"
      }
    ],
    "name": "registerEventSchema",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string[]",
        "name": "segments",
        "type": "string[]"
      }
    ],
    "name": "removeValue",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "wallet",
        "type": "address"
      },
      {
        "internalType": "enum RoleControl.Role",
        "name": "role",
        "type": "uint8"
      }
    ],
    "name": "revokeRole",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string[]",
        "name": "segments",
        "type": "string[]"
      },
      {
        "internalType": "bool",
        "name": "value",
        "type": "bool"
      }
    ],
    "name": "setBool",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string[]",
        "name": "segments",
        "type": "string[]"
      },
      {
        "internalType": "bytes",
        "name": "data",
        "type": "bytes"
      }
    ],
    "name": "setBytes",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string[]",
        "name": "segments",
        "type": "string[]"
      },
      {
        "internalType": "int256",
        "name": "value",
        "type": "int256"
      }
    ],
    "name": "setInt",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string[]",
        "name": "segments",
        "type": "string[]"
      },
      {
        "internalType": "string",
        "name": "value",
        "type": "string"
      }
    ],
    "name": "setString",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string[]",
        "name": "segments",
        "type": "string[]"
      },
      {
        "internalType": "uint256",
        "name": "value",
        "type": "uint256"
      }
    ],
    "name": "setUint",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string[]",
        "name": "segments",
        "type": "string[]"
      },
      {
        "internalType": "enum EtherDatabaseLib.DataType",
        "name": "dataType",
        "type": "uint8"
      },
      {
        "internalType": "bytes",
        "name": "data",
        "type": "bytes"
      }
    ],
    "name": "setValue",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "components": [
          {
            "internalType": "string[]",
            "name": "segments",
            "type": "string[]"
          },
          {
            "internalType": "enum EtherDatabaseLib.DataType",
            "name": "dataType",
            "type": "uint8"
          },
          {
            "internalType": "bytes",
            "name": "data",
            "type": "bytes"
          }
        ],
        "internalType": "struct EtherDatabaseLib.BatchSetValue[]",
        "name": "values",
        "type": "tuple[]"
      }
    ],
    "name": "setValues",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
] as const;
