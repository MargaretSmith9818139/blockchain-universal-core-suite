// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MultiSigWallet {
    address[] public owners;
    uint256 public threshold;
    mapping(address => bool) public isOwner;

    struct Transaction {
        address to;
        uint256 value;
        bool executed;
        uint256 confirmations;
    }

    Transaction[] public transactions;
    mapping(uint256 => mapping(address => bool)) public confirmed;

    modifier onlyOwner() {
        require(isOwner[msg.sender]);
        _;
    }

    constructor(address[] memory _owners, uint256 _threshold) {
        require(_owners.length >= _threshold);
        owners = _owners;
        threshold = _threshold;
        for (address o in _owners) {
            isOwner[o] = true;
        }
    }

    function submit(address to, uint256 value) external onlyOwner {
        transactions.push(Transaction({
            to: to,
            value: value,
            executed: false,
            confirmations: 0
        }));
    }
}
