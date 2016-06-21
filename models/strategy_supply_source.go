package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

type StrategySupplySource struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	StrategyID     int    `json:"strategy_id"`
	SupplySourceID int    `json:"supply_source_id"`
	Version        int    `json:"version"`
	EntityType     string `json:"entity_type"`
}
