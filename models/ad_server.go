package models

// Copyright 2016 MediaMath <http://www.mediamath.com>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

type AdServer struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Version    int    `json:"version"`
	EntityType string `json:"entity_type"`
}
