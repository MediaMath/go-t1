package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

type StrategyDayPart struct {
	CreatedOn  time.Time `json:"created_on"`
	Days       string    `json:"days"`
	EndHour    int       `json:"end_hour"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	StartHour  int       `json:"start_hour"`
	Status     bool      `json:"status"`
	StrategyID int       `json:"strategy_id"`
	UpdatedOn  time.Time `json:"updated_on"`
	UserTime   bool      `json:"user_time"`
	Version    int       `json:"version"`
	EntityType string    `json:"entity_type"`
}
