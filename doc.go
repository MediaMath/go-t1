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

	func main() {
		conf := cookie.Config{"myusername", "mypassword", "myapikey"}
		auth := cookie.New(conf, t1.ProductionURL)
		client := t1.NewClient(auth, conf.APIKey, t1.ProductionURL)
		var orgs []models.Organization
		meta, err := client.Organizations.List(nil, &orgs)
	}

The client here takes any `*http.Client` that can authenticate properly. As
such, you can also use a client from the OAuth2 package:

	package main

	import (
		"github.com/MediaMath/go-t1"
		"github.com/MediaMath/go-t1/models"
		"golang.org/x/oauth2"
		"golang.org/x/oauth2/mediamath"
	)

	func main() {
		conf := oauth2.Config{
			ClientID:     os.Getenv("T1_API_CLIENT_ID"),
			ClientSecret: os.Getenv("T1_API_CLIENT_SECRET"),
			Endpoint:     mediamath.Endpoint,
			RedirectURL:  "https://www.mediamath.com/",
		}
		// get the token
		c := conf.Client(oauth2.NoContext, tok)
		t1Client := t1.NewClient(c, conf.ClientID, t1.ProductionURL)
		var org models.Organization
		meta, err := t1Client.Organizations.Get(100048, &org)
	}

For API documentation, please visit https://developer.mediamath.com
*/
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

const (
	versionMajor = 0
	versionMinor = 1
	versionPatch = 0
)
