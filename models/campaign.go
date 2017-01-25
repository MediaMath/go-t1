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
	AdServerFee               float32       `json:"ad_server_fee"`
	AdServerID                int           `json:"ad_server_id"`
	AdServerPassword          string        `json:"ad_server_password"`
	AdServerUsername          string        `json:"ad_server_username"`
	AdvertiserID              int           `json:"advertiser_id"`
	AgencyFeePct              float32       `json:"agency_fee_pct"`
	ConversionType            string        `json:"conversion_type"`
	ConversionVariableMinutes int           `json:"conversion_variable_minutes"`
	CreatedOn                 t1time.T1Time `json:"created_on"`
	CurrencyCode              string        `json:"currency_code"`
	DCSDataIsCampaignLevel    bool          `json:"dcs_data_is_campaign_level"`
	EndDate                   t1time.T1Time `json:"end_date"`
	FrequencyAmount           int           `json:"frequency_amount"`
	FrequencyInterval         string        `json:"frequency_interval"`
	FrequencyType             string        `json:"frequency_type"`
	GoalCategory              string        `json:"goal_category"`
	GoalType                  string        `json:"goal_type"`
	GoalValue                 float32       `json:"has_custom_attribution"`
	ID                        int           `json:"id,omitempty,readonly"`
	InitialStartDate          t1time.T1Time `json:"initial_start_date"`
	IOName                    string        `json:"io_name"`
	IOReferenceNum            string        `json:"io_reference_num"`
	MarginPct                 float32       `json:"margin_pct"`
	MeritPixelID              int           `json:"merit_pixel_id"`
	Name                      string        `json:"name"`
	PacingAlert               float32       `json:"pacing_alert"`
	PCWindowMinutes           int           `json:"pc_window_minutes"`
	PVPct                     float32       `json:"pv_pct"`
	PVWindowMinutes           int           `json:"pv_window_minutes"`
	ServiceType               string        `json:"service_type"`
	SpendCapAmount            float32       `json:"spend_cap_amount"`
	SpendCapAutomatic         bool          `json:"spend_cap_automatic"`
	SpendCapEnabled           bool          `json:"spend_cap_enabled"`
	StartDate                 t1time.T1Time `json:"start_date"`
	Status                    bool          `json:"status"`
	UpdatedOn                 t1time.T1Time `json:"updated_on"`
	UseDefaultAdServer        bool          `json:"use_default_ad_server"`
	UseMMFreq                 bool          `json:"use_mm_freq"`
	ZoneName                  string        `json:"zone_name"`
	Version                   int           `json:"version"`
	EntityType                string        `json:"entity_type"`
}
