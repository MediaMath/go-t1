package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// Vendor represents a vendor object
type Vendor struct {
	AdxApproved            bool      `json:"adx_approved"`
	AdxDeclarationRequired bool      `json:"adx_declaration_required"`
	AdxSSLApproved         bool      `json:"adx_ssl_approved"`
	AdxVendorIdentifier    int       `json:"adx_vendor_identifier"`
	AdxVideoApproved       bool      `json:"adx_video_approved"`
	AdxVideoSSLApproved    bool      `json:"adx_video_ssl_approved"`
	CreatedOn              time.Time `json:"created_on"`
	Description            string    `json:"description"`
	ID                     int       `json:"id"`
	IsEligible             bool      `json:"is_eligible"`
	MMContractAvailable    bool      `json:"mm_contract_available"`
	Name                   string    `json:"name"`
	RateCardPrice          float32   `json:"rate_card_price"`
	RateCardType           string    `json:"rate_card_type"`
	UpdatedOn              time.Time `json:"updated_on"`
	VendorType             string    `json:"vendor_type"`
	Version                int       `json:"version"`
	WholesalePrice         float32   `json:"wholesale_price"`
	EntityType             string    `json:"entity_type"`
}
