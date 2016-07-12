package t1

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"bytes"
	"net/url"
	"sync"
)

var (
	bufferPool *sync.Pool
	valuesPool *sync.Pool
)

func init() {
	bufferPool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
	valuesPool = &sync.Pool{
		New: func() interface{} {
			return make(url.Values)
		},
	}
}
