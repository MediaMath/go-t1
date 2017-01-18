package t1

// Copyright 2016-2017 MediaMath
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
