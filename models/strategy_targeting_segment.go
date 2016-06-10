package models

// Copyright 2016 MediaMath <http://www.mediamath.com>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

type StrategyTargetingSegment struct {
	CreatedOn          time.Time `json:"created_on"`
	GroupIdentifier    string    `json:"group_identifier"`
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Operator           string    `json:"operator"`
	Restriction        string    `json:"restriction"`
	StrategyID         int       `json:"strategy_id"`
	TargetingSegmentID int       `json:"targeting_segment_id"`
	UpdatedOn          time.Time `json:"updated_on"`
	UserCPM            float32   `json:"user_cpm"`
	Version            int       `json:"version"`
	EntityType         string    `json:"entity_type"`
}
