package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// SupplySource represents a supply_source object
type SupplySource struct {
	BidderExchangeIdentifier int       `json:"bidder_exchange_identifier"`
	Code                     string    `json:"code"`
	CreatedOn                time.Time `json:"created_on"`
	DefaultSeatIdentifier    string    `json:"default_seat_identifier"`
	Distribute               bool      `json:"distribute"`
	HasDisplay               bool      `json:"has_display"`
	HasMobileDisplay         bool      `json:"has_mobile_display"`
	HasMobileVideo           bool      `json:"has_mobile_video"`
	HasVideo                 bool      `json:"has_video"`
	ID                       int       `json:"id"`
	IsProservice             bool      `json:"is_proservice"`
	MMSafe                   bool      `json:"mm_safe"`
	Name                     string    `json:"name"`
	ParentSupplyID           int       `json:"parent_supply_id"`
	PixelTag                 string    `json:"pixel_tag"`
	PMPEnabled               bool      `json:"pmp_enabled"`
	RTBEnabled               bool      `json:"rtb_enabled"`
	RTBType                  string    `json:"rtb_type"`
	SeatEnabled              bool      `json:"seat_enabled"`
	Status                   bool      `json:"status"`
	SupplyType               string    `json:"supply_type"`
	UpdatedOn                time.Time `json:"updated_on"`
	UsePool                  bool      `json:"use_pool"`
	Version                  int       `json:"version"`
	EntityType               string    `json:"entity_type"`
}
