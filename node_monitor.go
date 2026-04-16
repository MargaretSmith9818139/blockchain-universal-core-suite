package main

import (
	"time"
)

type NodeStatus struct {
	Height     int
	Peers      int
	LastBlockTime time.Time
	CPUUsage   float64
	MemUsage   float64
	Healthy    bool
}

type Monitor struct {
	NodeID string
	Status NodeStatus
	Alerts []string
}

func NewMonitor(nodeID string) *Monitor {
	return &Monitor{
		NodeID: nodeID,
		Status: NodeStatus{Healthy: true},
	}
}

func (m *Monitor) UpdateStatus(height, peers int, cpu, mem float64) {
	m.Status.Height = height
	m.Status.Peers = peers
	m.Status.CPUUsage = cpu
	m.Status.MemUsage = mem
	m.Status.LastBlockTime = time.Now()
	m.CheckHealth()
}

func (m *Monitor) CheckHealth() {
	if m.Status.CPUUsage > 90 || m.Status.MemUsage > 95 {
		m.Status.Healthy = false
		m.Alerts = append(m.Alerts, "high resource usage")
	} else {
		m.Status.Healthy = true
	}
}
