// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract StandardToken {
    string public name;
    string public symbol;
    uint8 public decimals;
    uint256 public totalSupply;

    mapping(address => uint256) public balances;
    mapping(address => mapping(address => uint256)) public allowances;

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);

    constructor(string memory _name, string memory _symbol, uint8 _decimals, uint256 _total) {
        name = _name;
        symbol = _symbol;
        decimals = _decimals;
        totalSupply = _total;
        balances[msg.sender] = _total;
    }

    function transfer(address to, uint256 value) external returns (bool) {
        require(balances[msg.sender] >= value);
        balances[msg.sender] -= value;
        balances[to] += value;
        emit Transfer(msg.sender, to, value);
        return true;
    }
}
