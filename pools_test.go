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
	"bytes"
	"net/url"
	"testing"
)

func TestBufferPoolReturnsBuffer(t *testing.T) {
	bufVal := bufferPool.New()
	_, ok := bufVal.(*bytes.Buffer)
	if !ok {
		t.Errorf("buffer pool: want *bytes.Buffer, got %T", bufVal)
	}
}

func TestValuesPoolReturnsValues(t *testing.T) {
	val := valuesPool.New()
	_, ok := val.(url.Values)
	if !ok {
		t.Errorf("buffer pool: want url.Values, got %T", val)
	}
}
