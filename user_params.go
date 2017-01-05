package t1

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

var (
	// DefaultParams is an empty UserParams struct to use in methods when
	// no user params are set.
	DefaultParams = &UserParams{}
)

// UserParams is a struct of parameters to add to the query string of the
// URL of a request.
type UserParams struct {
	Full       []string `json:"full,omitempty"`
	PageLimit  int      `json:"page_limit,omitempty"`
	PageOffset int      `json:"page_offset,omitempty"`
	Q          string   `json:"q,omitempty"`
	SortBy     string   `json:"sort_by,omitempty"`
}
