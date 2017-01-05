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
	"encoding/json"
	"testing"
	"time"
)

var notJSONEncodableTimes = []struct {
	time T1Time
	want string
}{
	{T1Time(time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC)), "Time.MarshalJSON: year outside of range [0,9999]"},
	{T1Time(time.Date(-1, 1, 1, 0, 0, 0, 0, time.UTC)), "Time.MarshalJSON: year outside of range [0,9999]"},
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
		Time T1Time
	}
	var a A
	if err := json.Unmarshal([]byte(data), &a); err != nil {
		t.Error(err)
	}
}
