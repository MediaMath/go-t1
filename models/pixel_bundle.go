package models

// Copyright 2016-2017 MediaMath
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	ID                   int       `json:"id,omitempty,readonly"`
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
