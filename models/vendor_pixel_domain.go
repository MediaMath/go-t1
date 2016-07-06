package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// VendorPixelDomain represents a vendor_pixel_domain object
type VendorPixelDomain struct {
	CreatedOn      time.Time `json:"created_on"`
	Domain         string    `json:"domain"`
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	UpdatedOn      time.Time `json:"updated_on"`
	VendorDomainID int       `json:"vendor_domain_id"`
	VendorPixelID  int       `json:"vendor_pixel_id"`
	Version        int       `json:"version"`
	EntityType     string    `json:"entity_type"`
}
