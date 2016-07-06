package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// PublisherSite represents a publisher_site object
type PublisherSite struct {
	CreatedOn   time.Time `json:"created_on"`
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	PublisherID int       `json:"publisher_id"`
	UpdatedOn   time.Time `json:"updated_on"`
	Version     int       `json:"version"`
	EntityType  string    `json:"entity_type"`
}
