package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

type SitePlacement struct {
	BillMediaToClient bool      `json:"bill_media_to_client"`
	CreatedOn         time.Time `json:"created_on"`
	DealSource        string    `json:"deal_source"`
	DisplayText       string    `json:"display_text"`
	ID                int       `json:"id"`
	MediaType         string    `json:"media_type"`
	Name              string    `json:"name"`
	PMPType           string    `json:"pmp_type"`
	PublisherSiteID   int       `json:"publisher_site_id"`
	UpdatedOn         time.Time `json:"updated_on"`
	Version           int       `json:"version"`
	EntityType        string    `json:"entity_type"`
}
