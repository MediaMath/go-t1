package models

// Copyright 2016 MediaMath <http://www.mediamath.com>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

type SiteList struct {
	CreatedOn        time.Time `json:"created_on"`
	Filename         string    `json:"filename"`
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	OrganizationID   int       `json:"organization_id"`
	Restriction      string    `json:"restriction"`
	SitesCount       int       `json:"sites_count"`
	SitesCountDomain int       `json:"sites_count_domain"`
	Status           bool      `json:"status"`
	UpdatedOn        time.Time `json:"updated_on"`
	Version          int       `json:"version"`
	EntityType       string    `json:"entity_type"`
}
