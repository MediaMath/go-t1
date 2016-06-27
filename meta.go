package t1

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"github.com/MediaMath/go-t1/time"
)

type Meta struct {
	CalledOn   time.AdamaTime `json:"called_on"`
	Count      int            `json:"count"`
	ETag       string         `json:"etag"`
	NextPage   string         `json:"next_page"`
	Offset     int            `json:"offset"`
	Status     string         `json:"status"`
	TotalCount int            `json:"total_count"`
}
