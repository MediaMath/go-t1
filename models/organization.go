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
	t1time "github.com/MediaMath/go-t1/time"
)

// Organization represents an organization object
type Organization struct {
	Address1                           string        `json:"address_1,omitempty"`
	Address2                           string        `json:"address_2,omitempty"`
	AdxSeatAccountID                   int64         `json:"adx_seat_account_id,omitempty"`
	AllowBYOPrice                      bool          `json:"allow_byo_price,omitempty"`
	AllowXAgencyPixels                 bool          `json:"allow_x_agency_pixels,omitempty"`
	BillingCountryCode                 string        `json:"billing_country_code,omitempty"`
	City                               string        `json:"city,omitempty"`
	ConnectedIDType                    string        `json:"connected_id_type"`
	ContactName                        string        `json:"contact_name,omitempty"`
	Country                            string        `json:"country,omitempty"`
	CreatedOn                          t1time.T1Time `json:"created_on,readonly"`
	CurrencyCode                       string        `json:"currency_code,omitempty"`
	DmpEnabled                         string        `json:"dmp_enabled,omitempty"`
	EligibleForDataSharing             bool          `json:"eligible_for_data_sharing"`
	Facebook                           string        `json:"facebook,omitempty"`
	ID                                 int           `json:"id,omitempty,readonly"`
	LinkedIn                           string        `json:"linkedin,omitempty"`
	MMContactName                      string        `json:"mm_contact_name,omitempty"`
	Name                               string        `json:"name,omitempty"`
	OptOutConnectedID                  bool          `json:"opt_out_connected_id,omitempty"`
	OptOutConnectedIDMathID            bool          `json:"opt_out_connected_id_mathid,omitempty"`
	OrgType                            []string      `json:"org_type,omitempty"`
	OverrideSuspiciousTrafficFilter    bool          `json:"override_suspicious_traffic_filter"`
	Phone                              string        `json:"phone,omitempty"`
	RestrictTargetingToDeterministicID bool          `json:"restrict_targeting_to_deterministic_id,omitempty"`
	RestrictTargetingToSameDeviceID    bool          `json:"restrict_targeting_to_same_device_id,omitempty"`
	State                              string        `json:"state,omitempty"`
	Status                             bool          `json:"status,omitempty"`
	SuspiciousTrafficFilterLevel       int           `json:"suspicious_traffic_filter_level,omitempty"`
	TagRuleset                         string        `json:"tag_ruleset,omitempty"`
	Terminated                         bool          `json:"terminated"`
	Twitter                            string        `json:"twitter,omitempty"`
	UpdatedOn                          t1time.T1Time `json:"updated_on,readonly,omitempty"`
	UseAdsTxt                          bool          `json:"use_ads_txt"`
	UseEvidonOptout                    bool          `json:"use_evidon_optout,omitempty"`
	Version                            int           `json:"version,omitempty"`
	Website                            string        `json:"website,omitempty"`
	Zip                                string        `json:"zip,omitempty"`
	EntityType                         string        `json:"entity_type,readonly"`
}
