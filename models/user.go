package models

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"github.com/MediaMath/go-t1/time"
)

// User represents a user object
type User struct {
	AccessInternalFees        bool          `json:"access_internal_fees"`
	Active                    bool          `json:"active"`
	CreatedOn                 t1time.T1Time `json:"created_on"`
	CreatorID                 int           `json:"creator_id"`
	EditCampaigns             bool          `json:"edit_campaigns"`
	EditDataDefinition        bool          `json:"edit_data_definition"`
	EditMarginsAndPerformance bool          `json:"edit_margins_and_performance"`
	Email                     string        `json:"email,omitempty,readonly"`
	Fax                       string        `json:"fax"`
	FirstName                 string        `json:"first_name"`
	ID                        int           `json:"id,omitempty,readonly"`
	LabsEnableRMX             bool          `json:"labs_enable_rmx"`
	LastLoginOn               t1time.T1Time `json:"last_login_on"`
	LastName                  string        `json:"last_name"`
	LinkLDAP                  bool          `json:"link_ldap"`
	Mobile                    string        `json:"mobile"`
	Name                      string        `json:"name"`
	Password                  string        `json:"password,omitempty"`
	PasswordResetSent         t1time.T1Time `json:"password_reset_sent"`
	PasswordResetToken        string        `json:"password_reset_token"`
	Phone                     string        `json:"phone"`
	Role                      string        `json:"role"`
	Scope                     string        `json:"scope"`
	SSOAuthSent               t1time.T1Time `json:"sso_auth_sent"`
	SSOAuthToken              string        `json:"sso_auth_token"`
	Title                     string        `json:"title"`
	Type                      string        `json:"type"`
	UpdatedOn                 t1time.T1Time `json:"updated_on"`
	Username                  string        `json:"username"`
	Version                   int           `json:"version"`
	ViewDataDefinition        bool          `json:"view_data_definition"`
	ViewDMPReports            bool          `json:"view_dmp_reports"`
	ViewOrganizations         bool          `json:"view_organizations"`
	ViewSegments              bool          `json:"view_segments"`
	EntityType                string        `json:"entity_type"`
}
