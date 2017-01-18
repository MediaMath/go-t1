package t1

import (
	"time"
)

// Session represents an Adama session object
type Session struct {
	ExpiresAt time.Time
	SessionID string
	UserName  string `json:"name"`
	UserID    int    `json:"id"`
}
