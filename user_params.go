package t1

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
