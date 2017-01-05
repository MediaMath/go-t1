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
	"time"
)

// StrategyConcept represents a strategy_concept object
type StrategyConcept struct {
	ConceptID  int       `json:"concept_id"`
	CreatedOn  time.Time `json:"created_on"`
	ID         int       `json:"id,omitempty,readonly"`
	Name       string    `json:"name"`
	Status     bool      `json:"status"`
	StrategyID int       `json:"strategy_id"`
	UpdatedOn  time.Time `json:"updated_on"`
	Version    int       `json:"version"`
	Weighting  string    `json:"weighting"`
	EntityType string    `json:"entity_type"`
}
