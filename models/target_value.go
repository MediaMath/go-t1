package models

// Copyright 2016 MediaMath <http://www.mediamath.com>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

type TargetValue struct {
	ChildCount        int    `json:"child_count"`
	DimensionCode     string `json:"dimension_code"`
	ID                int    `json:"id"`
	IsTargetable      bool   `json:"is_targetable"`
	Name              string `json:"name"`
	TargetDimensionID int    `json:"target_dimension_id"`
	Type              string `json:"type"`
	Value             int    `json:"value"`
	Version           int    `json:"version"`
	EntityType        string `json:"entity_type"`
}
