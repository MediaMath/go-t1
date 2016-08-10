package t1

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Encoder represents an object that can be encoded.
type Encoder interface {
	Encode() string
}

// Decoder represents an object that can be decoded.
type Decoder interface {
	Decode(interface{}) error
}

// Messager represents an error that has a message.
type Messager interface {
	Message() string
}
