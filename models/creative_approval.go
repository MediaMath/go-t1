package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// CreativeApproval represents a creative_approval object
type CreativeApproval struct {
	AdditionalDetail     int       `json:"additional_detail"`
	ApprovalStatus       int       `json:"approval_status"`
	AtomicCreativeID     int       `json:"atomic_creative_id"`
	CreatedOn            time.Time `json:"created_on"`
	CreativeImportFileID int       `json:"creative_import_file_id"`
	ExternalIdentifier   string    `json:"external_identifier"`
	ID                   int       `json:"id,omitempty,readonly"`
	Name                 string    `json:"name"`
	RejectedReason       string    `json:"rejected_reason"`
	SupplySourceID       int       `json:"supply_source_id"`
	UpdatedOn            time.Time `json:"updated_on"`
	Version              int       `json:"version"`
	EntityType           string    `json:"entity_type"`
}
