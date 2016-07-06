package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// StrategyConcept represents a strategy_concept object
type StrategyConcept struct {
	ConceptID  int       `json:"concept_id"`
	CreatedOn  time.Time `json:"created_on"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Status     bool      `json:"status"`
	StrategyID int       `json:"strategy_id"`
	UpdatedOn  time.Time `json:"updated_on"`
	Version    int       `json:"version"`
	Weighting  string    `json:"weighting"`
	EntityType string    `json:"entity_type"`
}
