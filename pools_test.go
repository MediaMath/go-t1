package t1

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
