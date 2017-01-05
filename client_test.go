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

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the T1 client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

// setup sets up a test HTTP server along with a github.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	url, _ := url.Parse(server.URL)
	// client configured to use test server
	client = NewClient(nil, "", url)
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

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

	if !c.RateLimitReset.IsZero() {
		t.Errorf("Client rate limit reset not initialized to zero: %v", c.RateLimitReset)
	}
}

func TestNewClient_services(t *testing.T) {
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

func TestNewRequest_badURL(t *testing.T) {
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
func TestNewRequest_emptyBody(t *testing.T) {
	c := NewClient(nil, "", nil)
	req, err := c.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("NewRequest returned unexpected error: %v", err)
	}
	if req.Body != nil {
		t.Error("constructed request contains a non-nil Body")
	}
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
		fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	body := new(foo)
	client.Do(req, body)

	want := &foo{"a"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}

func TestDo_httpError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(req, nil)

	if err == nil {
		t.Error("Expected HTTP 400 error.")
	}
}

// Test handling of an error caused by the internal http client's Do()
// function. A redirect loop is pretty unlikely to occur within the T1
// API, but does allow us to exercise the right code path.
func TestDo_redirectLoop(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(req, nil)

	if err == nil {
		t.Fatal("Expected error to be returned.")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected a URL error; got %#v.", err)
	}
}

// ensure rate limit is not parsed for other error responses
func TestDo_rateLimit_errorResponse(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Date", "Mon, 02 Jan 2016 15:04:05 GMT")
		w.Header().Add("Retry-After", "1")
		w.Header().Add(headerMasheryError, masheryErrorCodeRateLimit)
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(req, nil)

	if err == nil {
		t.Fatal("Expected error to be returned.")
	}
	if _, ok := err.(*RateLimitError); ok {
		t.Errorf("Did not expect a *RateLimitError error; got %#v.", err)
	}
	if !client.RateLimitReset.IsZero() {
		t.Errorf("Client rate limit reset changed: %v", client.RateLimitReset)
	}
}

// Ensure *RateLimitError is returned when API rate limit is exceeded.
func TestDo_rateLimit_rateLimitError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Date", "Mon, 02 Jan 2016 15:04:05 GMT")
		w.Header().Add("Retry-After", "1")
		w.Header().Add("Content-Type", mediaTypeMashery)
		w.Header().Add(headerMasheryError, masheryErrorCodeRateLimit)
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, `<h1>Developer Over Qps</h1>`)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(req, nil)

	if err == nil {
		t.Fatal("Expected error to be returned.")
	}
	rateLimitErr, ok := err.(*RateLimitError)
	if !ok {
		t.Fatalf("Expected a *RateLimitError error; got %#v.", err)
	}
	reset := time.Date(2016, 1, 2, 15, 4, 6, 0, time.UTC)
	if got := rateLimitErr.RetryAt.UTC(); got != reset {
		t.Errorf("rateLimitErr rate reset: got %v, want %v", got, reset)
	}
}

// Ensure a network call is not made when it's known that API rate limit is still exceeded.
func TestDo_rateLimit_noNetworkCall(t *testing.T) {
	setup()
	defer teardown()

	now := time.Now().UTC().Round(time.Second)

	mux.HandleFunc("/first", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Date", now.Format(time.RFC1123))
		w.Header().Add("Retry-After", "60")
		w.Header().Set("Content-Type", mediaTypeMashery)
		w.Header().Add(headerMasheryError, masheryErrorCodeRateLimit)
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, `<h1>Developer Over Qps</h1>`)
	})

	madeNetworkCall := false
	mux.HandleFunc("/second", func(w http.ResponseWriter, r *http.Request) {
		madeNetworkCall = true
	})

	// First request is made, and it makes the client aware of rate reset time being in the future.
	req, _ := client.NewRequest("GET", "/first", nil)
	reset := now.Add(time.Minute) // Rate reset is a minute from now, with 1 second precision.
	client.Do(req, nil)

	// Second request should not cause a network call to be made, since client can predict a rate limit error.
	req, _ = client.NewRequest("GET", "/second", nil)
	_, err := client.Do(req, nil)

	if madeNetworkCall {
		t.Fatal("Network call was made, even though rate limit is known to still be exceeded.")
	}

	if err == nil {
		t.Fatal("Expected error to be returned.")
	}

	rateLimitErr, ok := err.(*RateLimitError)
	if !ok {
		t.Fatalf("Expected a *RateLimitError error; got %#v.", err)
	}
	if got := rateLimitErr.RetryAt.UTC(); got != reset {
		t.Errorf("rateLimitErr rate reset: got %v, want %v", got, reset)
	}
}

func TestDo_noContent(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	var body json.RawMessage

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(req, &body)

	if err != nil {
		t.Fatalf("Do returned unexpected error: %v", err)
	}
}
