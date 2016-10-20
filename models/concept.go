package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// Concept represents a concept object
type Concept struct {
	AdvertiserID int       `json:"advertiser_id"`
	CreatedOn    time.Time `json:"created_on"`
	ID           int       `json:"id,omitempty,readonly"`
	Name         string    `json:"name"`
	Status       bool      `json:"status"`
	UpdatedOn    time.Time `json:"updated_on"`
	Version      int       `json:"version"`
	EntityType   string    `json:"entity_type"`
}
