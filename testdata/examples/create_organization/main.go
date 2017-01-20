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
	c, err := cookie.New(conf, t1.SandboxURL)
	if err != nil {
		log.Fatalf("initial login: %v", err)
	}

	// Construct new t1 client
	t1Client := t1.NewClient(c, conf.APIKey, t1.SandboxURL)

	// New org we will create
	org := models.Organization{
		Address1:      "4 WTC",
		Address2:      "45th Fl",
		AllowBYOPrice: true,
		City:          "New York",
		ContactName:   "Prasanna Swaminathan",
		Country:       "US",
		MMContactName: "Prasanna Swaminathan",
		Name:          "go-t1 example",
		Phone:         "1234567890",
		State:         "NY",
		Zip:           "10007",
	}

	meta, err := t1Client.Organizations.Save(&org)
	if err != nil {
		log.Fatalf("update org error: %v", err)
	}
	fmt.Printf("Meta:\t%#v\n\nOrg:\t%#v\n", meta, org)
}
