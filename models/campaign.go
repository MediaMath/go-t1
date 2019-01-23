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

type CurrencyValue struct {
	CurrencyCode string `json:"currency_code"`
	Value        int    `json:"value"`
}

type (
	TotalBudget    CurrencyValue
	GoalValue      CurrencyValue
	SpendCapAmount CurrencyValue
)

// Campaign represents a campaign object
type Campaign struct {
	UseMmFreq                          bool             `json:"use_mm_freq"`
	SpendCapType                       string           `json:"spend_cap_type"`
	SourceCampaignID                   int              `json:"source_campaign_id"`
	IsProgrammaticGuaranteed           bool             `json:"is_programmatic_guaranteed"`
	ZoneName                           string           `json:"zone_name"`
	FrequencyInterval                  string           `json:"frequency_interval"`
	UpdatedOn                          t1time.T1Time    `json:"updated_on"`
	UseDefaultAdServer                 bool             `json:"use_default_ad_server"`
	CostPassthroughPercentEnabled      bool             `json:"cost_passthrough_percent_enabled"`
	InitialStartDate                   t1time.T1Time    `json:"initial_start_date"`
	RestrictTargetingToDeterministicID bool             `json:"restrict_targeting_to_deterministic_id"`
	CreatedOn                          t1time.T1Time    `json:"created_on"`
	ID                                 int              `json:"id"`
	TotalBudget                        []TotalBudget    `json:"total_budget"`
	ServiceType                        string           `json:"service_type"`
	CurrencyCode                       string           `json:"currency_code"`
	HasCustomAttribution               bool             `json:"has_custom_attribution"`
	Name                               string           `json:"name"`
	AdServerID                         int              `json:"ad_server_id"`
	CostPassthroughCPMEnabled          bool             `json:"cost_passthrough_cpm_enabled"`
	FrequencyAmount                    int              `json:"frequency_amount"`
	RestrictTargetingToSameDeviceID    bool             `json:"restrict_targeting_to_same_device_id"`
	SuspiciousTrafficFilterLevel       int              `json:"suspicious_traffic_filter_level"`
	EndDate                            t1time.T1Time    `json:"end_date"`
	FrequencyOptimization              bool             `json:"frequency_optimization"`
	AdvertiserID                       int              `json:"advertiser_id"`
	ConversionVariableMinutes          int              `json:"conversion_variable_minutes"`
	MinimizeMultiAds                   bool             `json:"minimize_multi_ads"`
	Status                             bool             `json:"status"`
	GoalType                           string           `json:"goal_type"`
	FrequencyType                      string           `json:"frequency_type"`
	MarginPCT                          int              `json:"margin_pct"`
	DCSDataIsCampaignLevel             bool             `json:"dcs_data_is_campaign_level"`
	PvWindowMinutes                    int              `json:"pv_window_minutes"`
	ImpressionCapType                  string           `json:"impression_cap_type"`
	GoalValue                          []GoalValue      `json:"goal_value"`
	ImpressionCapAmount                int              `json:"impression_cap_amount"`
	SpendCapAmount                     []SpendCapAmount `json:"spend_cap_amount"`
	EntityType                         string           `json:"entity_type"`
	PcWindowMinutes                    int              `json:"pc_window_minutes"`
	IOName                             string           `json:"io_name"`
	StartDate                          t1time.T1Time    `json:"start_date"`
	OverrideSuspiciousTrafficFilter    bool             `json:"override_suspicious_traffic_filter"`
	IOReferenceNum                     string           `json:"io_reference_num"`
	Version                            int              `json:"version"`
	AgencyFeePCT                       int              `json:"agency_fee_pct"`
	ImpressionCapAutomatic             bool             `json:"impression_cap_automatic"`
	SpendCapAutomatic                  bool             `json:"spend_cap_automatic"`
	ConversionType                     string           `json:"conversion_type"`
	TotalImpressionBudget              int              `json:"total_impression_budget"`
	MeritPixelID                       int              `json:"merit_pixel_id"`
	PVPCT                              int              `json:"pv_pct"`
}
