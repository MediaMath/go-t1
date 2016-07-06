package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// PixelBundle represents a pixel_bundle object (what you think of as a pixel)
type PixelBundle struct {
	AdvertiserID         int       `json:"advertiser_id"`
	AgencyID             int       `json:"agency_id"`
	CostCPM              float32   `json:"cost_cpm"`
	CostCPTS             float32   `json:"cost_cpts"`
	CostPctCPM           float32   `json:"cost_pct_cpm"`
	CreatedOn            time.Time `json:"created_on"`
	Eligible             bool      `json:"eligible"`
	ExternalIdentifier   string    `json:"external_identifier"`
	ID                   int       `json:"id"`
	Keywords             string    `json:"keywords"`
	Name                 string    `json:"name"`
	PixelType            string    `json:"pixel_type"`
	Pricing              string    `json:"pricing"`
	ProviderID           int       `json:"provider_id"`
	RMXConversionMinutes int       `json:"rmx_conversion_minutes"`
	RMXConversionType    string    `json:"rmx_conversion_type"`
	RMXFriendly          bool      `json:"rmx_friendly"`
	RMXMerit             bool      `json:"rmx_merit"`
	RMXPCWindowMinutes   int       `json:"rmx_pc_window_minutes"`
	RMXPVWindowMinutes   int       `json:"rmx_pv_window_minutes"`
	SegmentOp            string    `json:"segment_op"`
	Status               bool      `json:"status"`
	Tags                 string    `json:"tags"`
	TagType              string    `json:"tag_type"`
	Type                 string    `json:"type"`
	UpdatedOn            time.Time `json:"updated_on"`
	Version              int       `json:"version"`
	EntityType           string    `json:"entity_type"`
}
