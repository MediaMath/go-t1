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

// Deal represents a deal object
type Deal struct {
	AdvertiserID   int           `json:"advertiser_id"`
	CreatedOn      t1time.T1Time `json:"created_on"`
	CurrencyCode   string        `json:"currency_code"`
	DealIdentifier string        `json:"deal_identifier"`
	DealSource     string        `json:"deal_source"`
	Description    string        `json:"description"`
	EndDatetime    t1time.T1Time `json:"end_datetime"`
	ID             int           `json:"id,omitempty,readonly"`
	MediaType      string        `json:"media_type"`
	Name           string        `json:"name"`
	PartnerSourced bool          `json:"partner_sourced"`
	PriceMethod    string        `json:"price_method"`
	PriceType      string        `json:"price_type"`
	PublisherID    int           `json:"publisher_id"`
	StartDatetime  t1time.T1Time `json:"start_datetime"`
	Status         bool          `json:"status"`
	SupplySourceID int           `json:"supply_source_id"`
	UpdatedOn      t1time.T1Time `json:"updated_on"`
	ZoneName       string        `json:"zone_name"`
	Version        int           `json:"version"`
	EntityType     string        `json:"entity_type"`
}
