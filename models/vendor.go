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
	ID                     int       `json:"id,omitempty,readonly"`
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
