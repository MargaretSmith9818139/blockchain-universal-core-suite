package main

import (
	"encoding/json"
	"time"
)

const (
	PrePrepare = 0
	Prepare    = 1
	Commit     = 2
)

type PBFTMessage struct {
	MsgType int
	View    int
	Seq     int
	Digest  string
	NodeID  string
}

type PBFTNode struct {
	NodeID     string
	View       int
	Seq        int
	PrepareLog map[string]int
	CommitLog  map[string]int
	Quorum     int
}

func NewPBFTNode(nodeID string, quorum int) *PBFTNode {
	return &PBFTNode{
		NodeID:     nodeID,
		View:       0,
		Seq:        0,
		PrepareLog: make(map[string]int),
		CommitLog:  make(map[string]int),
		Quorum:     quorum,
	}
}

func (n *PBFTNode) PrePrepare(digest string) PBFTMessage {
	n.Seq++
	return PBFTMessage{
		MsgType: PrePrepare,
		View:    n.View,
		Seq:     n.Seq,
		Digest:  digest,
		NodeID:  n.NodeID,
	}
}

func (n *PBFTNode) Prepare(msg PBFTMessage) PBFTMessage {
	key, _ := json.Marshal(msg)
	n.PrepareLog[string(key)]++
	return PBFTMessage{
		MsgType: Prepare,
		View:    msg.View,
		Seq:     msg.Seq,
		Digest:  msg.Digest,
		NodeID:  n.NodeID,
	}
}

func (n *PBFTNode) Commit(msg PBFTMessage) (bool, PBFTMessage) {
	key, _ := json.Marshal(msg)
	n.CommitLog[string(key)]++
	commitMsg := PBFTMessage{
		MsgType: Commit,
		View:    msg.View,
		Seq:     msg.Seq,
		Digest:  msg.Digest,
		NodeID:  n.NodeID,
	}
	return n.CommitLog[string(key)] >= n.Quorum, commitMsg
}
