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
	// T1Fmt is the time layout used by Adama. It is most similar to
	// RFC3339, which is the standard used by encoding/json for encoding and
	// decoding of time.Time values. it differs in that it doesn't have a
	// colon in the time zone specifier, which is mandated by RFC3339, but
	// not by ISO8601 (which this does adhere to).
	T1Fmt = "\"2006-01-02T15:04:05Z0700\""
)

// T1Time is a time.Time type with a different JSON-parsing format.
type T1Time time.Time

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format,
// except without the colon in the timezone specifier.
func (t *T1Time) UnmarshalJSON(data []byte) error {
	// Fractional seconds are handled implicitly by Parse
	val, err := time.Parse(T1Fmt, string(data))
	*t = T1Time(val)
	return err
}

// String fulfills the Stringer interface.
func (t T1Time) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05.999999999 -0700 MST")
}

// MarshalJSON fulfills the Marshaler interface for T1Time.
func (t T1Time) MarshalJSON() ([]byte, error) {
	// T1 *accepts* times as RFC3339 (that is, with the colon in the time
	// zone specifier). So we can just delegate to the standard.
	return time.Time(t).MarshalJSON()
}
