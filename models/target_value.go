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

// TargetValue represents a target_value object
type TargetValue struct {
	ChildCount        int    `json:"child_count"`
	DimensionCode     string `json:"dimension_code"`
	ID                int    `json:"id,omitempty,readonly"`
	IsTargetable      bool   `json:"is_targetable"`
	Name              string `json:"name"`
	TargetDimensionID int    `json:"target_dimension_id"`
	Type              string `json:"type"`
	Value             int    `json:"value"`
	Version           int    `json:"version"`
	EntityType        string `json:"entity_type"`
}
