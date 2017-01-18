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
)

func TestUnmarshalSession(t *testing.T) {
	data := `{"Time": "2016-01-01T00:00:00"}`
	type A struct {
		Time SessionTime
	}
	var a A
	if err := json.Unmarshal([]byte(data), &a); err != nil {
		t.Error(err)
	}
}
