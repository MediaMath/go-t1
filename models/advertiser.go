package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"github.com/MediaMath/go-t1/time"
)

// Advertiser represents an advertiser object
type Advertiser struct {
	AdServerID              int            `json:"ad_server_id"`
	AgencyID                int            `json:"agency_id"`
	AllowXStratOptimization bool           `json:"allow_x_strat_optimization"`
	CreatedOn               time.AdamaTime `json:"created_on"`
	DMPEnabled              string         `json:"dmp_enabled"`
	Domain                  string         `json:"domain"`
	FrequencyAmount         int            `json:"frequency_amount"`
	FrequencyInterval       string         `json:"frequency_interval"`
	FrequencyType           string         `json:"frequency_type"`
	ID                      int            `json:"id"`
	MinimizeMultiAds        bool           `json:"minimize_multi_ads"`
	Name                    string         `json:"name"`
	Status                  bool           `json:"status"`
	UpdatedOn               time.AdamaTime `json:"updated_on"`
	Version                 int            `json:"version"`
	VerticalID              int            `json:"vertical_id"`
	EntityType              string         `json:"entity_type"`
}
