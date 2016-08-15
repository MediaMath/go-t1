/*
Package t1 is a Go library for MediaMath's APIs.

Construct a new T1 client, then use the various services on the client to
access different parts of the MediaMath API modules.

For example:

	package main

	import (
		"github.com/MediaMath/go-t1"
		"github.com/MediaMath/go-t1/authenticators/cookie"
		"github.com/MediaMath/go-t1/models"
	)

	conf := cookie.Config{"myusername", "mypassword", "myapikey"}
	auth := cookie.New(conf, t1.ProductionURL)
	client := t1.NewClient(auth, conf.APIKey, t1.ProductionURL)
	var orgs []models.Organization
	meta, err := client.Organizations.List(nil, &orgs)

For API documentation, please visit https://developer.mediamath.com
*/
package t1

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

const (
	versionMajor = 0
	versionMinor = 1
	versionPatch = 0
)
