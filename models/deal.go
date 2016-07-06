package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// Deal represents a deal object
type Deal struct {
	AdvertiserID   int       `json:"advertiser_id"`
	CreatedOn      time.Time `json:"created_on"`
	CurrencyCode   string    `json:"currency_code"`
	DealIdentifier string    `json:"deal_identifier"`
	DealSource     string    `json:"deal_source"`
	Description    string    `json:"description"`
	EndDatetime    time.Time `json:"end_datetime"`
	ID             int       `json:"id"`
	MediaType      string    `json:"media_type"`
	Name           string    `json:"name"`
	PartnerSourced bool      `json:"partner_sourced"`
	PriceMethod    string    `json:"price_method"`
	PriceType      string    `json:"price_type"`
	PublisherID    int       `json:"publisher_id"`
	StartDatetime  time.Time `json:"start_datetime"`
	Status         bool      `json:"status"`
	SupplySourceID int       `json:"supply_source_id"`
	UpdatedOn      time.Time `json:"updated_on"`
	ZoneName       string    `json:"zone_name"`
	Version        int       `json:"version"`
	EntityType     string    `json:"entity_type"`
}
