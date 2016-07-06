package t1

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Encoder represents an object that can be encoded.
type Encoder interface {
	Encode() string
}
