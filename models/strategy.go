package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// Strategy represents a strategy object
type Strategy struct {
	AudienceSegmentExcludeOp       string    `json:"audience_segment_exclude_op"`
	AudienceSegmentIncludeOp       string    `json:"audience_segment_include_op"`
	BidAggressiveness              float32   `json:"bid_aggresiveness"`
	BidPriceIsMediaOnly            bool      `json:"bid_price_is_media_only"`
	Budget                         float32   `json:"budget"`
	CampaignID                     int       `json:"campaign_id"`
	CreatedOn                      time.Time `json:"created_on"`
	CurrencyCode                   string    `json:"currency_code"`
	Description                    string    `json:"description"`
	EffectiveGoalValue             float32   `json:"effective_goal_value"`
	EndDate                        time.Time `json:"end_date"`
	FeatureCompatibility           string    `json:"feature_compatibility"`
	FrequencyAmount                int       `json:"frequency_amount"`
	FrequencyInterval              string    `json:"frequency_interval"`
	FrequencyType                  string    `json:"frequency_type"`
	GoalType                       string    `json:"goal_type"`
	ID                             int       `json:"id,omitempty,readonly"`
	ImpressionCap                  int       `json:"impression_cap"`
	MediaType                      string    `json:"media_type"`
	Name                           string    `json:"name"`
	PacingInterval                 string    `json:"pacing_interval"`
	PacingType                     string    `json:"pacing_type"`
	PixelTargetExpr                string    `json:"pixel_target_expr"` // TODO custom type?
	RunOnAllExchanges              bool      `json:"run_on_all_exchanges"`
	RunOnAllPMP                    bool      `json:"run_on_all_pmp"`
	RunOnDisplay                   bool      `json:"run_on_display"`
	RunOnMobile                    bool      `json:"run_on_mobile"`
	RunOnStreaming                 bool      `json:"run_on_streaming"`
	SiteRestrictionTransparentURLs bool      `json:"site_restriction_transparent_urls"`
	SiteSelectiveness              string    `json:"site_selectiveness"`
	StartDate                      time.Time `json:"start_date"`
	Status                         bool      `json:"status"`
	SupplyType                     string    `json:"supply_type"`
	Type                           string    `json:"type"`
	UpdatedOn                      time.Time `json:"updated_on"`
	UseCampaignEnd                 bool      `json:"use_campaign_end"`
	UseCampaignStart               bool      `json:"use_campaign_start"`
	UseMMFreq                      bool      `json:"use_mm_freq"`
	UseOptimization                bool      `json:"use_optimization"`
	Version                        int       `json:"version"`
	EntityType                     string    `json:"entity_type"`
}
