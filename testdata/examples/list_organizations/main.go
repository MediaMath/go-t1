package main

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
	"fmt"
	"log"

	"github.com/MediaMath/go-t1"
	"github.com/MediaMath/go-t1/authenticators/cookie"
	"github.com/MediaMath/go-t1/models"
)

func main() {
	// Set up configuration from envvars
	conf := cookie.GetCredentialsFromEnv()

	// Create new *http.Client with these credentials
	c, err := cookie.New(conf, t1.ProductionURL)
	if err != nil {
		log.Fatalf("initial login: %v", err)
	}

	// Construct new t1 client
	t1Client := t1.NewClient(c, conf.APIKey, t1.ProductionURL)

	// Model object gets passed in to the various service methods
	var orgs []models.Organization
	params := t1.UserParams{
		PageLimit: 5,
		SortBy:    "-id",
	}

	meta, err := t1Client.Organizations.List(&params, &orgs)
	if err != nil {
		log.Fatalf("get org error: %v", err)
	}
	fmt.Printf("Meta:\t%#v\nOrg:\t%#v\n", meta, orgs)
	var org models.Organization
	if len(orgs) > 0 {
		org = orgs[0]
		fmt.Printf("org: %#v\n", org)
	}
	t1Client.Organizations.Save(&org)
}
