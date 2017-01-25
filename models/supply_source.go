package models

// Copyright 2016-2017 MediaMath
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"github.com/MediaMath/go-t1/time"
)

// SupplySource represents a supply_source object
type SupplySource struct {
	BidderExchangeIdentifier int           `json:"bidder_exchange_identifier"`
	Code                     string        `json:"code"`
	CreatedOn                t1time.T1Time `json:"created_on"`
	DefaultSeatIdentifier    string        `json:"default_seat_identifier"`
	Distribute               bool          `json:"distribute"`
	HasDisplay               bool          `json:"has_display"`
	HasMobileDisplay         bool          `json:"has_mobile_display"`
	HasMobileVideo           bool          `json:"has_mobile_video"`
	HasVideo                 bool          `json:"has_video"`
	ID                       int           `json:"id,omitempty,readonly"`
	IsProservice             bool          `json:"is_proservice"`
	MMSafe                   bool          `json:"mm_safe"`
	Name                     string        `json:"name"`
	ParentSupplyID           int           `json:"parent_supply_id"`
	PixelTag                 string        `json:"pixel_tag"`
	PMPEnabled               bool          `json:"pmp_enabled"`
	RTBEnabled               bool          `json:"rtb_enabled"`
	RTBType                  string        `json:"rtb_type"`
	SeatEnabled              bool          `json:"seat_enabled"`
	Status                   bool          `json:"status"`
	SupplyType               string        `json:"supply_type"`
	UpdatedOn                t1time.T1Time `json:"updated_on"`
	UsePool                  bool          `json:"use_pool"`
	Version                  int           `json:"version"`
	EntityType               string        `json:"entity_type"`
}
