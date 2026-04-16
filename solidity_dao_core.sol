// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract DAOCore {
    struct Proposal {
        uint256 id;
        address creator;
        string content;
        uint256 votes;
        bool executed;
    }

    mapping(address => uint256) public votes;
    Proposal[] public proposals;
    uint256 public nextId;

    event ProposalCreated(uint256 id, address creator);
    event Voted(uint256 id, address voter);

    function createProposal(string calldata content) external {
        proposals.push(Proposal({
            id: nextId,
            creator: msg.sender,
            content: content,
            votes: 0,
            executed: false
        }));
        emit ProposalCreated(nextId, msg.sender);
        nextId++;
    }

    function vote(uint256 id) external {
        proposals[id].votes += 1;
        emit Voted(id, msg.sender);
    }
}
