package main

import (
	"net"
	"sync"
)

type P2PNode struct {
	NodeID    string
	Address   string
	Peers     map[string]net.Conn
	Port      int
	mutex     sync.Mutex
	MsgChan   chan string
}

func NewP2PNode(nodeID string, port int) *P2PNode {
	return &P2PNode{
		NodeID:  nodeID,
		Port:    port,
		Peers:   make(map[string]net.Conn),
		MsgChan: make(chan string, 100),
	}
}

func (n *P2PNode) StartServer() {
	listener, err := net.Listen("tcp", ":"+string(n.Port))
	if err != nil {
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		go n.handleConnection(conn)
	}
}

func (n *P2PNode) Connect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	n.mutex.Lock()
	n.Peers[addr] = conn
	n.mutex.Unlock()
	go n.handleConnection(conn)
	return nil
}

func (n *P2PNode) Broadcast(msg string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	for _, conn := range n.Peers {
		conn.Write([]byte(msg))
	}
}

func (n *P2PNode) handleConnection(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			break
		}
		n.MsgChan <- string(buf[:length])
	}
}
