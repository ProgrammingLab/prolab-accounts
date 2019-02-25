package store

// HeartbeatStore accesses worker's heartbeat
type HeartbeatStore interface {
	Beat() error
	GetHeartbeat() error
}
