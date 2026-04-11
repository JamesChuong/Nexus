package types

type GameSession struct {
	HostId     string `json:"hostId"`
	Status     string `json:"status"`
	Transport  string `json:"transport"`
	RelayId    string `json:"relayId"`
	MaxPlayers int    `json:"maxPlayers"`
}
