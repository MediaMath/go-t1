package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"github.com/MediaMath/go-t1/time"
)

// Organization represents an organization object
type Organization struct {
	Address1           string        `json:"address_1,omitempty"`
	Address2           string        `json:"address_2,omitempty"`
	AdXSeatAccountID   int           `json:"adx_seat_account_id,omitempty"`
	AllowBYOPrice      bool          `json:"allow_byo_price,omitempty"`
	AllowXAgencyPixels bool          `json:"allow_x_agency_pixels,omitempty"`
	City               string        `json:"city,omitempty"`
	ContactName        string        `json:"contact_name,omitempty"`
	Country            string        `json:"country,omitempty"`
	CreatedOn          t1time.T1Time `json:"created_on,readonly"`
	CurrencyCode       string        `json:"currency_code,omitempty"`
	ID                 int           `json:"id,omitempty,readonly"`
	MMContactName      string        `json:"mm_contact_name,omitempty"`
	Name               string        `json:"name,omitempty"`
	OrgType            []string      `json:"org_type,omitempty"`
	Phone              string        `json:"phone,omitempty"`
	State              string        `json:"state,omitempty"`
	Status             bool          `json:"status,omitempty"`
	TagRuleset         string        `json:"tag_ruleset,omitempty"`
	UpdatedOn          t1time.T1Time `json:"updated_on,readonly,omitempty"`
	UseEvidonOptout    bool          `json:"use_evidon_optout,omitempty"`
	Version            int           `json:"version,omitempty"`
	Zip                string        `json:"zip,omitempty"`
	EntityType         string        `json:"entity_type,readonly"`
}
