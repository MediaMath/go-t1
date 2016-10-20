package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// Agency represents an agency object
type Agency struct {
	AllowXAdvOptimization bool      `json:"allow_x_adv_optimization"`
	AllowXAdvPixels       bool      `json:"allow_x_adv_pixels"`
	CreatedOn             time.Time `json:"created_on"`
	ID                    int       `json:"id,omitempty,readonly"`
	Name                  string    `json:"name"`
	OrganizationID        int       `json:"organization_id"`
	Status                bool      `json:"status"`
	UpdatedOn             time.Time `json:"updated_on"`
	Version               int       `json:"version"`
	EntityType            string    `json:"entity_type"`
}
