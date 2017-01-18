package t1

import (
	"github.com/MediaMath/go-t1/time"
)

// Session represents an Adama session object
type Session struct {
	nestedSession `json:"session"`
	UserName      string `json:"name"`
	UserID        int    `json:"id"`
}

// nestedSession is the session data nested within the session response.
// This is conceptually part of the same response, so there's no reason to
// expose it independently. Instead, by embedding it, we can expose its
// fields without exposing the nesting structure or jumping through hoops
// to unmarshal it.
type nestedSession struct {
	SessionID string
	ExpiresAt t1time.SessionTime `json:"expires"`
}
