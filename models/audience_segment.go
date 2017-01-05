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
