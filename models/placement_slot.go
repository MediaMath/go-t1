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
	"time"
)

// PlacementSlot represents a placement_slot object
type PlacementSlot struct {
	AdSlot            int       `json:"ad_slot"`
	AllowRemnant      bool      `json:"allow_remnant"`
	AuctionType       string    `json:"auction_type"`
	Budget            float32   `json:"budget"`
	BuyPrice          float32   `json:"buy_price"`
	BuyPriceType      string    `json:"buy_price_type"`
	CreatedOn         time.Time `json:"created_on"`
	Description       string    `json:"description"`
	EndDate           time.Time `json:"end_date"`
	EstVolume         float32   `json:"est_volume"`
	FrequencyAmount   int       `json:"frequency_amount"`
	FrequencyInterval string    `json:"frequency_interval"`
	FrequencyType     string    `json:"frequency_type"`
	Height            int       `json:"height"`
	ID                int       `json:"id,omitempty,readonly"`
	Name              string    `json:"name"`
	PRMPubCeiling     float32   `json:"prm_pub_ceiling"`
	PRMPubMarkup      float32   `json:"prm_pub_markup"`
	SellPrice         float32   `json:"sell_price"`
	SellPriceType     string    `json:"sell_price_type"`
	SitePlacementID   int       `json:"site_placement_id"`
	StartDate         time.Time `json:"start_date"`
	UpdatedOn         time.Time `json:"updated_on"`
	Version           int       `json:"version"`
	VolumeUnit        string    `json:"volume_unit"`
	Width             int       `json:"width"`
	EntityType        string    `json:"entity_type"`
}
