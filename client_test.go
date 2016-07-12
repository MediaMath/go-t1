package t1

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"runtime"
	"strings"
	"testing"
)

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil, "myapikey", nil)

	if got, want := c.BaseURL.String(), ProductionURL.String(); got != want {
		t.Errorf("NewClient BaseURL: got %v, want %v", got, want)
	}

	userAgent := fmt.Sprintf("go-t1/%d.%d.%d %s",
		versionMajor, versionMinor, versionPatch,
		strings.Replace(runtime.Version(), "go", "go-http-package/", -1))
	if got, want := c.userAgent, userAgent; got != want {
		t.Errorf("NewClient UserAgent: got %v, want %v", got, want)
	}

	if got, want := c.APIKey, "myapikey"; got != want {
		t.Errorf("NewClient UserAgent: got %v, want %v", got, want)
	}
}

func TestNewClientServices(t *testing.T) {
	c := NewClient(nil, "", nil)
	serviceNames := []struct {
		service *EntityService
		name    string
	}{
		{c.Advertisers, "Advertisers"},
		{c.AtomicCreatives, "AtomicCreatives"},
		{c.Agencies, "Agencies"},
		{c.AdServers, "AdServers"},
		{c.Campaigns, "Campaigns"},
		{c.Concepts, "Concepts"},
		{c.Creatives, "Creatives"},
		{c.Deals, "Deals"},
		{c.Organizations, "Organizations"},
		{c.Pixels, "Pixels"},
		{c.PixelBundles, "PixelBundles"},
		{c.PixelProviders, "PixelProviders"},
		{c.PlacementSlots, "PlacementSlots"},
		{c.Publishers, "Publishers"},
		{c.PublisherSites, "PublisherSites"},
		{c.SitePlacements, "SitePlacements"},
		{c.SiteLists, "SiteLists"},
		{c.Strategies, "Strategies"},
		{c.SupplySources, "SupplySources"},
		{c.Users, "Users"},
		{c.Vendors, "Vendors"},
		{c.VendorContracts, "VendorContracts"},
		{c.VendorDomains, "VendorDomains"},
		{c.VendorPixels, "VendorPixels"},
		{c.VendorPixelDomains, "VendorPixelDomains"},
		{c.Verticals, "Verticals"},
	}
	for _, test := range serviceNames {
		if test.service == nil {
			t.Errorf("%s service: got nil, expected non-nil", test.name)
		}
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil, "", nil)

	inURL, outURL := "/foo", "https://api.mediamath.com/foo"
	inBody, outBody := url.Values{
		"name":    []string{"Test Name"},
		"version": []string{"1"},
	}, "name=Test+Name&version=1"

	req, err := c.NewRequest("GET", inURL, inBody)

	if err != nil {
		t.Fatal(err)
	}

	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL: got %v, want %v", inURL, got, want)
	}

	// test that body was URL-encoded
	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%q) Body: got %v, want %v", inBody, got, want)
	}

	// test that we send the right Content-Type header
	if got, want := req.Header.Get("Content-Type"), mediaTypeURLEncoded; got != want {
		t.Errorf("NewRequest POST Content-Type: got %v, want %v", got, want)
	}

	// test that default user-agent is attached to the request
	if got, want := req.Header.Get("User-Agent"), c.userAgent; got != want {
		t.Errorf("NewRequest() User-Agent: got %v, want %v", got, want)
	}
}

func TestNewRequestBadURL(t *testing.T) {
	c := NewClient(nil, "", nil)
	_, err := c.NewRequest("GET", ":", nil)
	testURLParseError(t, err)
}

// If a nil body is passed to client.NewRequest, make sure that nil is also
// passed to http.NewRequest. In most cases, passing an io.Reader that returns
// no content is fine, since there is no difference between an HTTP request
// body that is an empty string versus one that is not set at all. However, in
// certain cases, intermediate systems may treat these differently resulting in
// subtle errors.
func TestNewRequestEmptyBody(t *testing.T) {
	c := NewClient(nil, "", nil)
	req, err := c.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("NewRequest returned unexpected error: %v", err)
	}
	if req.Body != nil {
		t.Error("constructed request contains a non-nil Body")
	}
}
