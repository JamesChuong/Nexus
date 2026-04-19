package types

type EnumInterface interface {
	String() string
}

type GameSessionStatus int

const (
	StatusActive GameSessionStatus = iota
	StatusEnded
)

var gameSessionStatusNames = map[GameSessionStatus]string{
	StatusActive: "active",
	StatusEnded:  "ended",
}

func (s GameSessionStatus) String() string {
	return gameSessionStatusNames[s]
}

type GameSessionTransport int

const (
	TransportP2P GameSessionTransport = iota
	TransportRelay
)

var gameSessionTransportNames = map[GameSessionTransport]string{
	TransportP2P:   "p2p",
	TransportRelay: "relay",
}

func (s GameSessionTransport) String() string {
	return gameSessionTransportNames[s]
}

type PlayerStatus int

const (
	StatusConnected PlayerStatus = iota
	StatusDisconnected
)

var playerStatusNames = map[PlayerStatus]string{
	StatusConnected:    "connected",
	StatusDisconnected: "disconnected",
}

func (s PlayerStatus) String() string {
	return playerStatusNames[s]
}
