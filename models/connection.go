package models

import "time"
import "encoding/json"

const (
	ConnectionTypeTikTokInt    int = 1
	ConnectionTypeInstagramInt int = 2
)

var ConnectionTypeStringMap = map[int]string{
	ConnectionTypeTikTokInt:    "tiktok",
	ConnectionTypeInstagramInt: "instagram",
}

type Connection struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	ConnectionType int       `json:"connection_type"`
	CreatedAt      time.Time `json:"created_at"`
}

func (c *Connection) MarshalJSON() ([]byte, error) {
	type Alias Connection
	return json.Marshal(&struct {
		*Alias
		ConnectionType string `json:"connection_type"`
	}{
		Alias:        (*Alias)(c),
		ConnectionType: ConnectionTypeStringMap[c.ConnectionType],
	})
}
