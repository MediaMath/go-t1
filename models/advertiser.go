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

// Advertiser represents an advertiser object
type Advertiser struct {
	AdServerID              int           `json:"ad_server_id"`
	AgencyID                int           `json:"agency_id"`
	AllowXStratOptimization bool          `json:"allow_x_strat_optimization"`
	CreatedOn               t1time.T1Time `json:"created_on"`
	DMPEnabled              string        `json:"dmp_enabled"`
	Domain                  string        `json:"domain"`
	FrequencyAmount         int           `json:"frequency_amount"`
	FrequencyInterval       string        `json:"frequency_interval"`
	FrequencyType           string        `json:"frequency_type"`
	ID                      int           `json:"id,omitempty,readonly"`
	MinimizeMultiAds        bool          `json:"minimize_multi_ads"`
	Name                    string        `json:"name"`
	Status                  bool          `json:"status"`
	UpdatedOn               t1time.T1Time `json:"updated_on"`
	Version                 int           `json:"version"`
	VerticalID              int           `json:"vertical_id"`
	EntityType              string        `json:"entity_type"`
}
