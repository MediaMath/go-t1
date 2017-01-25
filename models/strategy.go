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

// Strategy represents a strategy object
type Strategy struct {
	AudienceSegmentExcludeOp       string        `json:"audience_segment_exclude_op"`
	AudienceSegmentIncludeOp       string        `json:"audience_segment_include_op"`
	BidAggressiveness              float32       `json:"bid_aggresiveness"`
	BidPriceIsMediaOnly            bool          `json:"bid_price_is_media_only"`
	Budget                         float32       `json:"budget"`
	CampaignID                     int           `json:"campaign_id"`
	CreatedOn                      t1time.T1Time `json:"created_on"`
	CurrencyCode                   string        `json:"currency_code"`
	Description                    string        `json:"description"`
	EffectiveGoalValue             float32       `json:"effective_goal_value"`
	EndDate                        t1time.T1Time `json:"end_date"`
	FeatureCompatibility           string        `json:"feature_compatibility"`
	FrequencyAmount                int           `json:"frequency_amount"`
	FrequencyInterval              string        `json:"frequency_interval"`
	FrequencyType                  string        `json:"frequency_type"`
	GoalType                       string        `json:"goal_type"`
	ID                             int           `json:"id,omitempty,readonly"`
	ImpressionCap                  int           `json:"impression_cap"`
	MediaType                      string        `json:"media_type"`
	Name                           string        `json:"name"`
	PacingInterval                 string        `json:"pacing_interval"`
	PacingType                     string        `json:"pacing_type"`
	PixelTargetExpr                string        `json:"pixel_target_expr"` // TODO custom type?
	RunOnAllExchanges              bool          `json:"run_on_all_exchanges"`
	RunOnAllPMP                    bool          `json:"run_on_all_pmp"`
	RunOnDisplay                   bool          `json:"run_on_display"`
	RunOnMobile                    bool          `json:"run_on_mobile"`
	RunOnStreaming                 bool          `json:"run_on_streaming"`
	SiteRestrictionTransparentURLs bool          `json:"site_restriction_transparent_urls"`
	SiteSelectiveness              string        `json:"site_selectiveness"`
	StartDate                      t1time.T1Time `json:"start_date"`
	Status                         bool          `json:"status"`
	SupplyType                     string        `json:"supply_type"`
	Type                           string        `json:"type"`
	UpdatedOn                      t1time.T1Time `json:"updated_on"`
	UseCampaignEnd                 bool          `json:"use_campaign_end"`
	UseCampaignStart               bool          `json:"use_campaign_start"`
	UseMMFreq                      bool          `json:"use_mm_freq"`
	UseOptimization                bool          `json:"use_optimization"`
	Version                        int           `json:"version"`
	EntityType                     string        `json:"entity_type"`
}
