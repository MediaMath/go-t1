package t1

import (
	"time"
)

const (
	AdamaFmt = "\"2006-01-02T15:04:05Z0700\""
)

// type AdamaTime is a time.Time type with a different JSON-parsing format.
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

func (t AdamaTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05.999999999 -0700 MST")
}

func (t AdamaTime) MarshalJSON() ([]byte, error) {
	return time.Time(t).MarshalJSON()
}
