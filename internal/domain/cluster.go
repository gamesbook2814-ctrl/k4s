package domain

// ClusterInfo represents information about the connected cluster
type ClusterInfo struct {
	Name      string
	Context   string
	Namespace string
	Server    string
	Connected bool
}

// ConnectionStatus represents the current connection state
type ConnectionStatus int

const (
	StatusDisconnected ConnectionStatus = iota
	StatusConnecting
	StatusConnected
	StatusError
)

func (s ConnectionStatus) String() string {
	switch s {
	case StatusDisconnected:
		return "Disconnected"
	case StatusConnecting:
		return "Connecting..."
	case StatusConnected:
		return "Connected"
	case StatusError:
		return "Error"
	default:
		return "Unknown"
	}
}
