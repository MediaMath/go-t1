package models

// Copyright 2016 MediaMath <http://www.mediamath.com>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

type Meta struct {
	CalledOn   time.Time `json:"called_on"`
	Count      int       `json:"count"`
	ETag       string    `json:"etag"`
	NextPage   string    `json:"next_page"`
	Offset     int       `json:"offset"`
	Status     string    `json:"status"`
	TotalCount int       `json:"total_count"`
}
