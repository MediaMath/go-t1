package time

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

const (
	// AdamaFmt is the time layout used by Adama. It is most similar to
	// RFC3339, which is the standard used by encoding/json for encoding and
	// decoding of time.Time values. it differs in that it doesn't have a
	// colon in the time zone specifier, which is mandated by RFC3339, but
	// not by ISO8601 (which this does adhere to).
	AdamaFmt = "\"2006-01-02T15:04:05Z0700\""
)

// AdamaTime is a time.Time type with a different JSON-parsing format.
type AdamaTime time.Time

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format,
// except without the colon in the timezone specifier.
func (t *AdamaTime) UnmarshalJSON(data []byte) error {
	// Fractional seconds are handled implicitly by Parse
	val, err := time.Parse(AdamaFmt, string(data))
	*t = AdamaTime(val)
	return err
}

// String fulfills the Stringer interface.
func (t AdamaTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05.999999999 -0700 MST")
}

// MarshalJSON fulfills the Marshaler interface for AdamaTime.
func (t AdamaTime) MarshalJSON() ([]byte, error) {
	// Adama *accepts* times as RFC3339 (that is, with the colon in the time
	// zone specifier). So we can just delegate to the standard.
	return time.Time(t).MarshalJSON()
}
