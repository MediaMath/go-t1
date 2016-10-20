package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// VendorDomain represents a vendor_domain object
type VendorDomain struct {
	AllowSubdomainMatch bool      `json:"allow_subdomain_match"`
	CreatedOn           time.Time `json:"created_on"`
	Domain              string    `json:"domain"`
	ID                  int       `json:"id,omitempty,readonly"`
	Name                string    `json:"name"`
	UpdatedOn           time.Time `json:"updated_on"`
	VendorID            int       `json:"vendor_id"`
	Version             int       `json:"version"`
	EntityType          string    `json:"entity_type"`
}
