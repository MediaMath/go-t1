package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// TargetingVendor represents a targeting_vendor object
type TargetingVendor struct {
	CreatedOn             time.Time `json:"created_on"`
	ID                    int       `json:"id"`
	Name                  string    `json:"name"`
	NamespaceCode         string    `json:"namespace_code"`
	OrganizationID        int       `json:"organization_id"`
	SitesCountDomain      int       `json:"sites_count_domain"`
	TargetingVendorTypeID int       `json:"targeting_vendor_type_id"`
	UpdatedOn             time.Time `json:"updated_on"`
	Version               int       `json:"version"`
	EntityType            string    `json:"entity_type"`
}
