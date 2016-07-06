package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"time"
)

// AtomicCreative represents an atomic_creative object
type AtomicCreative struct {
	AdFormat             string    `json:"ad_format"`
	AdServerType         string    `json:"ad_server_type"`
	AdvertiserID         int       `json:"advertiser_id"`
	ApprovalStatus       string    `json:"approval_status"`
	BuildDate            time.Time `json:"build_date"`
	BuildErrors          string    `json:"build_errors"`
	Built                bool      `json:"built"`
	BuiltByUserID        int       `json:"built_by_user_id"`
	ClickThroughURL      string    `json:"click_through_url"`
	ClickURL             string    `json:"click_url"`
	ConceptID            int       `json:"concept_id"`
	CreatedOn            time.Time `json:"created_on"`
	CreativeImportFileID int       `json:"creative_import_file_id"`
	EditedTag            string    `json:"edited_tag"`
	EndDate              time.Time `json:"end_date"`
	ExpansionDirection   string    `json:"expansion_direction"`
	ExpansionTrigger     string    `json:"expansion_trigger"`
	ExternalIdentifier   string    `json:"external_identifier"`
	FileType             string    `json:"file_type"`
	HasSound             bool      `json:"has_sound"`
	Height               int       `json:"height"`
	ID                   int       `json:"id"`
	IsHTTPS              bool      `json:"is_https"`
	IsMultiCreative      bool      `json:"is_multi_creative"`
	LastModified         time.Time `json:"last_modified"`
	MediaType            string    `json:"media_type"`
	Name                 string    `json:"name"`
	RejectedReason       string    `json:"rejected_reason"`
	RichMedia            bool      `json:"rich_media"`
	RichMediaProvider    string    `json:"rich_media_provider"`
	StartDate            time.Time `json:"start_date"`
	Status               bool      `json:"status"`
	T1AS                 bool      `json:"t1as"`
	Tag                  string    `json:"tag"`
	TagType              string    `json:"tag_type"`
	TPASAdTag            string    `json:"tpas_ad_tag"`
	TPASAdTagName        string    `json:"tpas_ad_tag_name"`
	UpdatedOn            time.Time `json:"updated_on"`
	Version              int       `json:"version"`
	Width                int       `json:"width"`
	EntityType           string    `json:"entity_type"`
}
