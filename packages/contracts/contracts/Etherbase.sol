// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "./EtherbaseSource.sol";

contract Etherbase {
    address public validator;
    
    struct SourceInfo {
        address sourceAddress;
        address owner;
    }
    
    SourceInfo[] public sources;
    
    struct CustomContract {
        address contractAddress;
        string name;
        string[] eventNames;
        string contractABI;
        mapping(string => EventSchema) eventSchemas;
    }

    mapping(address => CustomContract) public customContracts;
    address[] public customContractAddresses;

    event SourceCreated(address indexed sourceAddress, address indexed owner);
    event CustomContractAdded(address indexed contractAddress, string name);
    event CustomContractSchemaAdded(address indexed contractAddress, string eventName);

    // Add custom errors at the contract level
    error ContractAlreadyExists();
    error EmptyContractABI();
    error ContractNotFound();

    constructor(address _validator) {
        validator = _validator;
    }

    function createSource() public returns (address) {
        EtherbaseSource newSource = new EtherbaseSource(msg.sender, validator);
        sources.push(SourceInfo({
            sourceAddress: address(newSource),
            owner: msg.sender
        }));
        emit SourceCreated(address(newSource), msg.sender);
        return address(newSource);
    }

    function addCustomContract(
        address contractAddress, 
        string memory name,
        string memory contractABI
    ) public {
        if (customContracts[contractAddress].contractAddress != address(0)) {
            revert ContractAlreadyExists();
        }
        if (bytes(contractABI).length == 0) {
            revert EmptyContractABI();
        }
        
        CustomContract storage newContract = customContracts[contractAddress];
        newContract.contractAddress = contractAddress;
        newContract.name = name;
        newContract.contractABI = contractABI;
        newContract.eventNames = new string[](0);
        customContractAddresses.push(contractAddress);
        
        emit CustomContractAdded(contractAddress, name);
    }

    function addCustomContractSchema(
        address contractAddress,
        string memory name,
        string memory id,
        bytes32 eventTopic,
        Argument[] memory args
    ) public {
        require(customContracts[contractAddress].contractAddress != address(0), "Contract not found");

        CustomContract storage contractObject = customContracts[contractAddress];
        
        // Check if event name already exists
        require(bytes(contractObject.eventSchemas[name].name).length == 0, "Event name already registered");
        
        // Count indexed arguments
        uint8 numIndexedArgs = 0;
        for (uint256 i = 0; i < args.length; i++) {
            if (args[i].isIndexed) {
                numIndexedArgs++;
            }
        }
        require(numIndexedArgs <= 3, "Too many indexed arguments");

        // Store the schema
        EventSchema storage newSchema = contractObject.eventSchemas[name];
        newSchema.name = name;
        newSchema.id = id;
        newSchema.eventTopic = eventTopic;
        newSchema.numIndexedArgs = numIndexedArgs;
        
        // Store arguments
        for (uint256 i = 0; i < args.length; i++) {
            newSchema.args.push(args[i]);
        }

        contractObject.eventNames.push(name);

        emit CustomContractSchemaAdded(contractAddress, name);
    }

    function getCustomContracts() public view returns (address[] memory) {
        return customContractAddresses;
    }

    function getCustomContractEventNames(address contractAddress) public view returns (string[] memory) {
        return customContracts[contractAddress].eventNames;
    }

    function getCustomContractSchema(address contractAddress, string memory eventName) 
        public view returns (EventSchema memory) {
        return customContracts[contractAddress].eventSchemas[eventName];
    }

    function getCustomContractABI(address contractAddress) 
        public view returns (string memory) {
        if (customContracts[contractAddress].contractAddress == address(0)) {
            revert ContractNotFound();
        }
        return customContracts[contractAddress].contractABI;
    }

    function getSources() public view returns (SourceInfo[] memory) {
        SourceInfo[] memory _sources = new SourceInfo[](sources.length);
        for (uint256 i = 0; i < sources.length; i++) {
            _sources[i] = sources[i];
        }
        return _sources;
    }

    // Global schemas
    function getGlobalSchemas() public pure returns (EventSchema[] memory) {
        EventSchema[] memory schemas = new EventSchema[](5);
        
        // ERC20 Transfer
        Argument[] memory transferArgs = new Argument[](3);
        transferArgs[0] = Argument("from", "address", true);
        transferArgs[1] = Argument("to", "address", true);
        transferArgs[2] = Argument("value", "uint256", false);
        schemas[0] = EventSchema({
            name: "Transfer",
            args: transferArgs,
            eventTopic: keccak256("Transfer(address,address,uint256)"),
            numIndexedArgs: 2,
            id: "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef:2"
        });

        // // ERC20 Approval
        // Argument[] memory approvalArgs = new Argument[](3);
        // approvalArgs[0] = Argument("owner", "address", true);
        // approvalArgs[1] = Argument("spender", "address", true);
        // approvalArgs[2] = Argument("value", "uint256", false);
    
        // schemas[1] = EventSchema({
        //     name: "Approval",
        //     args: approvalArgs,
        //     eventTopic: keccak256("Approval(address,address,uint256)"),
        //     numIndexedArgs: 2,
        //     id: "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925:2"
        // });

        // ERC721 Transfer
        Argument[] memory nftTransferArgs = new Argument[](3);
        nftTransferArgs[0] = Argument("from", "address", true);
        nftTransferArgs[1] = Argument("to", "address", true);
        nftTransferArgs[2] = Argument("tokenId", "uint256", true);
        schemas[2] = EventSchema({
            name: "Transfer",
            args: nftTransferArgs,
            eventTopic: keccak256("Transfer(address,address,uint256)"),
            numIndexedArgs: 3,
            id: "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef:3"
        });

        // // ERC721 Approval
        // Argument[] memory nftApprovalArgs = new Argument[](3);
        // nftApprovalArgs[0] = Argument("owner", "address", true);
        // nftApprovalArgs[1] = Argument("approved", "address", true);
        // nftApprovalArgs[2] = Argument("tokenId", "uint256", true);
        // schemas[3] = EventSchema({
        //     name: "Approval",
        //     args: nftApprovalArgs,
        //     eventTopic: keccak256("Approval(address,address,uint256)"),
        //     numIndexedArgs: 3,
        //     id: "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925:3"
        // });

        // // ERC721 ApprovalForAll
        // Argument[] memory approvalForAllArgs = new Argument[](3);
        // approvalForAllArgs[0] = Argument("owner", "address", true);
        // approvalForAllArgs[1] = Argument("operator", "address", true);
        // approvalForAllArgs[2] = Argument("approved", "bool", false);
        // schemas[4] = EventSchema({
        //     name: "ApprovalForAll",
        //     args: approvalForAllArgs,
        //     eventTopic: keccak256("ApprovalForAll(address,address,bool)"),
        //     numIndexedArgs: 2,
        //     id: "0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31:2"
        // });

        return schemas;
    }
}
