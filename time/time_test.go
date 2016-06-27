package time

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"encoding/json"
	"testing"
	"time"
)

var notJSONEncodableTimes = []struct {
	time AdamaTime
	want string
}{
	{AdamaTime(time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC)), "Time.MarshalJSON: year outside of range [0,9999]"},
	{AdamaTime(time.Date(-1, 1, 1, 0, 0, 0, 0, time.UTC)), "Time.MarshalJSON: year outside of range [0,9999]"},
}

func TestNotJSONEncodableTime(t *testing.T) {
	for _, tt := range notJSONEncodableTimes {
		_, err := tt.time.MarshalJSON()
		if err == nil || err.Error() != tt.want {
			t.Errorf("%v MarshalJSON error = %v, want %v", tt.time, err, tt.want)
		}
	}
}

func TestUnmarshal(t *testing.T) {
	data := `{"Time": "2016-01-01T00:00:00+0000"}`
	type A struct {
		Time AdamaTime
	}
	var a A
	if err := json.Unmarshal([]byte(data), &a); err != nil {
		t.Error(err)
	}
}
