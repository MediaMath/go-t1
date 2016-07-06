package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// Creative represents a creative object
type Creative struct {
	AtomicCreativeID int       `json:"atomic_creative_id"`
	CreatedOn        time.Time `json:"created_on"`
	ID               int       `json:"id"`
	LastModified     time.Time `json:"last_modified"`
	Name             string    `json:"name"`
	Tag              string    `json:"tag"`
	TagType          string    `json:"tag_type"`
	Version          int       `json:"version"`
	EntityType       string    `json:"entity_type"`
}
