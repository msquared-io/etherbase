export const EtherbaseAbi = [
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_validator",
        "type": "address"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "constructor"
  },
  {
    "inputs": [],
    "name": "ContractAlreadyExists",
    "type": "error"
  },
  {
    "inputs": [],
    "name": "ContractNotFound",
    "type": "error"
  },
  {
    "inputs": [],
    "name": "EmptyContractABI",
    "type": "error"
  },
  {
    "inputs": [],
    "name": "NotAuthorized",
    "type": "error"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "contractAddress",
        "type": "address"
      }
    ],
    "name": "CustomContractAdded",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "contractAddress",
        "type": "address"
      }
    ],
    "name": "CustomContractDeleted",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "contractAddress",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "eventName",
        "type": "string"
      }
    ],
    "name": "CustomContractSchemaAdded",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "sourceAddress",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "owner",
        "type": "address"
      }
    ],
    "name": "SourceCreated",
    "type": "event"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "contractAddress",
        "type": "address"
      },
      {
        "internalType": "string",
        "name": "contractABI",
        "type": "string"
      }
    ],
    "name": "addCustomContract",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "createSource",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "name": "customContractAddresses",
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
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "name": "customContracts",
    "outputs": [
      {
        "internalType": "address",
        "name": "contractAddress",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "addedBy",
        "type": "address"
      },
      {
        "internalType": "string",
        "name": "contractABI",
        "type": "string"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "contractAddress",
        "type": "address"
      }
    ],
    "name": "deleteCustomContract",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "contractAddress",
        "type": "address"
      }
    ],
    "name": "getCustomContract",
    "outputs": [
      {
        "components": [
          {
            "internalType": "address",
            "name": "contractAddress",
            "type": "address"
          },
          {
            "internalType": "address",
            "name": "addedBy",
            "type": "address"
          },
          {
            "internalType": "string",
            "name": "contractABI",
            "type": "string"
          }
        ],
        "internalType": "struct Etherbase.CustomContract",
        "name": "",
        "type": "tuple"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getCustomContracts",
    "outputs": [
      {
        "components": [
          {
            "internalType": "address",
            "name": "contractAddress",
            "type": "address"
          },
          {
            "internalType": "address",
            "name": "addedBy",
            "type": "address"
          },
          {
            "internalType": "string",
            "name": "contractABI",
            "type": "string"
          }
        ],
        "internalType": "struct Etherbase.CustomContract[]",
        "name": "",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getGlobalSchemas",
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
        "internalType": "struct EventSchema[]",
        "name": "",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "pure",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getSources",
    "outputs": [
      {
        "components": [
          {
            "internalType": "address",
            "name": "sourceAddress",
            "type": "address"
          },
          {
            "internalType": "address",
            "name": "owner",
            "type": "address"
          }
        ],
        "internalType": "struct Etherbase.SourceInfo[]",
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
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "name": "sources",
    "outputs": [
      {
        "internalType": "address",
        "name": "sourceAddress",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "owner",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "validator",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
] as const;
