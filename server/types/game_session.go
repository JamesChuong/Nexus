package types

type GameSession struct {
	HostId     string `json:"hostId"`
	Status     string `json:"status"`
	Transport  string `json:"transport"`
	SessionId  string `json:"sessionId"`
	MaxPlayers int    `json:"maxPlayers"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}
