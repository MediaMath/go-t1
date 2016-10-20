package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// AudienceSegment represents an audience_segment object
type AudienceSegment struct {
	AudienceVendorID int       `json:"audience_vendor_id"`
	Buyable          bool      `json:"buyable"`
	ChildCount       int       `json:"child_count"`
	CreatedOn        time.Time `json:"created_on"`
	EntityType       string    `json:"entity_type"`
	FullPath         string    `json:"full_path"`
	ID               int       `json:"id,omitempty,readonly"`
	Name             string    `json:"name"`
	Type             string    `json:"type"` // TODO ???
	UpdatedOn        time.Time `json:"updated_on"`
	Version          int       `json:"version"`
}
