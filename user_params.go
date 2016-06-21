package t1

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"bytes"
	"strconv"
	"strings"
)

var (
	DefaultParams = &UserParams{}
)

type UserParams struct {
	Full       []string `json:"full,omitempty"`
	PageLimit  int      `json:"page_limit,omitempty"`
	PageOffset int      `json:"page_offset,omitempty"`
	Q          string   `json:"q,omitempty"`
	SortBy     string   `json:"sort_by,omitempty"`
	With       []string `json:"with,omitempty"`
}

func (u *UserParams) Encode() string {
	buf := bufferPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		bufferPool.Put(buf)
	}()

	buf.WriteString("full=")
	buf.WriteString(strings.Join(u.Full, ","))
	buf.WriteString("&page_limit=")
	if u.PageLimit == 0 {
		buf.WriteString("100")
	} else {
		buf.WriteString(strconv.Itoa(u.PageLimit))
	}
	buf.WriteString("&page_offset=")
	buf.WriteString(strconv.Itoa(u.PageOffset))
	buf.WriteString("&q=")
	buf.WriteString(u.Q)
	if u.SortBy != "" {
		buf.WriteString("&sort_by=")
		buf.WriteString(u.SortBy)
	}

	return buf.String()
}
