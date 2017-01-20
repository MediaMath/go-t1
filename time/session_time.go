package t1time

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
	"time"
)

const (
	// sessionFmt is the time layout used by the /session endpoint.
	// It is an ISO8601-compatible format without a time zone specifier.
	sessionFmt = "\"2006-01-02T15:04:05\""
)

// SessionTime is a time.Time type with a different JSON-parsing format.
// SessionTime implements the json.Unmarshaler interface,
// but *not* json.Marshaler. We should never be marshaling a session time:
// it only exists for the /session endpoint, which is GET-only
type SessionTime time.Time

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format,
// except without a time zone specifier.
func (t *SessionTime) UnmarshalJSON(data []byte) error {
	// Fractional seconds are handled implicitly by Parse
	val, err := time.Parse(sessionFmt, string(data))
	*t = SessionTime(val)
	return err
}

// String fulfills the Stringer interface.
func (t SessionTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05.999999999 -0700 MST")
}
