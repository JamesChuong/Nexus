package types

type Player struct {
	PlayerId   string `json:"playerId"`
	PlayerName string `json:"playerName"`
	Status     string `json:"status"`
	IPAddress  string `json:"ipAddress"`
	LastPing   int64  `json:"lastPing"`
	SessionId  string `json:"sessionId"`
}
