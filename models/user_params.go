package models

// Copyright 2016 MediaMath <http://www.mediamath.com>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

type UserParams struct {
	Full       []string       `json:"full,omitempty"`
	Limit      map[string]int `json:"limit,omitempty"`
	PageLimit  int            `json:"page_limit,omitempty"`
	PageOffset int            `json:"page_offset,omitempty"`
	Q          string         `json:"q,omitempty"`
	SortBy     string         `json:"sort_by,omitempty"`
	With       []string       `json:"with,omitempty"`
}
