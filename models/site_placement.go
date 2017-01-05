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

// SitePlacement represents a site_placement object
type SitePlacement struct {
	BillMediaToClient bool      `json:"bill_media_to_client"`
	CreatedOn         time.Time `json:"created_on"`
	DealSource        string    `json:"deal_source"`
	DisplayText       string    `json:"display_text"`
	ID                int       `json:"id,omitempty,readonly"`
	MediaType         string    `json:"media_type"`
	Name              string    `json:"name"`
	PMPType           string    `json:"pmp_type"`
	PublisherSiteID   int       `json:"publisher_site_id"`
	UpdatedOn         time.Time `json:"updated_on"`
	Version           int       `json:"version"`
	EntityType        string    `json:"entity_type"`
}
