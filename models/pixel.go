package models

// Copyright 2016 MediaMath <http://www.mediamath.com>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

type Pixel struct {
	BundleID       int       `json:"bundle_id"`
	CreatedOn      time.Time `json:"created_on"`
	Distributed    bool      `json:"distributed"`
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	PixelType      string    `json:"pixel_type"`
	SupplySourceID int       `json:"supply_source_id"`
	Tag            string    `json:"tag"`
	UpdatedOn      time.Time `json:"updated_on"`
	Version        int       `json:"version"`
	EntityType     string    `json:"entity_type"`
}
