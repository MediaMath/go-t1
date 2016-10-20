package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// VendorPixel represents a vendor_pixel object
type VendorPixel struct {
	CreatedOn  time.Time `json:"created_on"`
	CreativeID int       `json:"creative_id"`
	ID         int       `json:"id,omitempty,readonly"`
	Name       string    `json:"name"`
	SetBy      string    `json:"set_by"`
	Tag        string    `json:"tag"`
	TagType    string    `json:"tag_type"`
	UpdatedOn  time.Time `json:"updated_on"`
	Version    int       `json:"version"`
	EntityType string    `json:"entity_type"`
}
