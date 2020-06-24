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

// Campaign represents a campaign object
type Campaign struct {
	AdServerFee               float32       `json:"ad_server_fee,omitempty"`
	AdServerID                int           `json:"ad_server_id,omitempty"`
	AdServerPassword          string        `json:"ad_server_password,omitempty"`
	AdServerUsername          string        `json:"ad_server_username,omitempty"`
	AdvertiserID              int           `json:"advertiser_id,omitempty"`
	AgencyFeePct              float32       `json:"agency_fee_pct,omitempty"`
	ConversionType            string        `json:"conversion_type,omitempty"`
	ConversionVariableMinutes int           `json:"conversion_variable_minutes,omitempty"`
	CreatedOn                 t1time.T1Time `json:"created_on,omitempty"`
	CurrencyCode              string        `json:"currency_code,omitempty"`
	DCSDataIsCampaignLevel    bool          `json:"dcs_data_is_campaign_level,omitempty"`
	EndDate                   t1time.T1Time `json:"end_date,omitempty"`
	FrequencyAmount           int           `json:"frequency_amount,omitempty"`
	FrequencyInterval         string        `json:"frequency_interval,omitempty"`
	FrequencyType             string        `json:"frequency_type,omitempty"`
	GoalCategory              string        `json:"goal_category,omitempty"`
	GoalType                  string        `json:"goal_type,omitempty"`
	GoalValue                 float32       `json:"goal_value,omitempty"`
	ID                        int           `json:"id,omitempty,readonly"`
	InitialStartDate          t1time.T1Time `json:"initial_start_date,omitempty"`
	IOName                    string        `json:"io_name,omitempty"`
	IOReferenceNum            string        `json:"io_reference_num,omitempty"`
	MarginPct                 float32       `json:"margin_pct,omitempty"`
	MeritPixelID              int           `json:"merit_pixel_id,omitempty"`
	Name                      string        `json:"name,omitempty"`
	PacingAlert               float32       `json:"pacing_alert,omitempty"`
	PCWindowMinutes           int           `json:"pc_window_minutes,omitempty"`
	PVPct                     float32       `json:"pv_pct,omitempty"`
	PVWindowMinutes           int           `json:"pv_window_minutes,omitempty"`
	ServiceType               string        `json:"service_type,omitempty"`
	SpendCapAmount            float32       `json:"spend_cap_amount,omitempty"`
	SpendCapAutomatic         bool          `json:"spend_cap_automatic,omitempty"`
	SpendCapEnabled           bool          `json:"spend_cap_enabled,omitempty"`
	StartDate                 t1time.T1Time `json:"start_date,omitempty"`
	Status                    bool          `json:"status,omitempty"`
	TotalBudget               float32       `json:"total_budget,omitempty"`
	UpdatedOn                 t1time.T1Time `json:"updated_on,omitempty"`
	UseDefaultAdServer        bool          `json:"use_default_ad_server,omitempty"`
	UseMMFreq                 bool          `json:"use_mm_freq,omitempty"`
	ZoneName                  string        `json:"zone_name,omitempty"`
	Version                   int           `json:"version,omitempty"`
	EntityType                string        `json:"entity_type,readonly,omitempty"`
}
