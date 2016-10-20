package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// PixelProvider represents a pixel_provider object
type PixelProvider struct {
	AgencyID     int       `json:"agency_id"`
	CreatedOn    time.Time `json:"created_on"`
	ExecutionBy  string    `json:"execution_by"`
	ID           int       `json:"id,omitempty,readonly"`
	Name         string    `json:"name"`
	Status       bool      `json:"status"`
	TaxonomyFile string    `json:"taxonomy_file"`
	UpdatedOn    time.Time `json:"updated_on"`
	VendorID     int       `json:"vendor_id"`
	Version      int       `json:"version"`
	EntityType   string    `json:"entity_type"`
}
