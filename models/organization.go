package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// Organization represents a organization object
type Organization struct {
	Address1           string    `json:"address_1"`
	Address2           string    `json:"address_2"`
	AdXSeatAccountID   int       `json:"adx_seat_account_id"`
	AllowBYOPrice      bool      `json:"allow_byo_price"`
	AllowXAgencyPixels bool      `json:"allow_x_agency_pixels"`
	City               string    `json:"city"`
	ContactName        string    `json:"contact_name"`
	Country            string    `json:"country"`
	CreatedOn          time.Time `json:"created_on"`
	CurrencyCode       string    `json:"currency_code"`
	ID                 int       `json:"id"`
	MMContactName      string    `json:"mm_contact_name"`
	Name               string    `json:"name"`
	OrgType            []string  `json:"org_type"`
	Phone              string    `json:"phone"`
	State              string    `json:"state"`
	Status             bool      `json:"status"`
	TagRuleset         string    `json:"tag_ruleset"`
	UpdatedOn          time.Time `json:"updated_on"`
	UseEvidonOptout    bool      `json:"use_evidon_optout"`
	Version            int       `json:"version"`
	Zip                string    `json:"zip"`
	EntityType         string    `json:"entity_type"`
}
