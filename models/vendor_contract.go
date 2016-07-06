package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// VendorContract represents a vendor_contract object
type VendorContract struct {
	CampaignID    int       `json:"campaign_id"`
	CreatedOn     time.Time `json:"created_on"`
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Price         float32   `json:"price"`
	RateCardType  string    `json:"rate_card_type"`
	UpdatedOn     time.Time `json:"updated_on"`
	UseMMContract bool      `json:"use_mm_contract"`
	VendorID      int       `json:"vendor_id"`
	Version       int       `json:"version"`
	EntityType    string    `json:"entity_type"`
}
