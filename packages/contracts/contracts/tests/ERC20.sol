// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract ERC20 {
    uint256 public totalSupply;
    mapping(address => uint256) public balances;
    address public owner;

    event Transfer(address indexed from, address indexed to, uint256 value);

    constructor(uint256 _initialSupply) {
        owner = msg.sender;
        totalSupply = _initialSupply;
        balances[owner] = _initialSupply;
    }

    function transfer(address _to, uint256 _value) public returns (bool) {
        require(balances[msg.sender] >= _value, "Insufficient balance");
        balances[msg.sender] -= _value;
        balances[_to] += _value;
        emit Transfer(msg.sender, _to, _value);
        return true;
    }

    function mint(address _to, uint256 _value) public {
        require(msg.sender == owner, "Only owner can mint");
        totalSupply += _value;
        balances[_to] += _value;
        emit Transfer(address(0), _to, _value);
    }
}
