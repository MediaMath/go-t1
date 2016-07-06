package t1

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	// Accept header value to get JSON output
	mediaTypeJSON = "application/vnd.mediamath.v1+json"
)

// Standard base URLs
var (
	ProductionURL, _ = url.Parse("https://api.mediamath.com")
	SandboxURL, _    = url.Parse("https://t1sandbox.mediamath.com")
)

// A Client manages communication with the MediaMath APIs
type Client struct {
	// HTTP client used to make requests. This client should know how to handle authentication
	client *http.Client

	// API Key to be used. This will be included as a query string
	// parameter in all requests made. API key typically will need
	// to be included in the construction of the HTTP client as well.
	// Independent of that, it must be included here.
	APIKey string

	// Base URL for API requests. Defaults to the production API endpoint,
	// but can be set to a specific domain for Sandbox or other similar
	// environments. Should be protocol and domain name without trailing slash
	BaseURL   *url.URL
	rateMu    sync.Mutex
	rateLimit Rate

	// User Agent will include library's version number
	userAgent string

	Advertisers        *EntityService
	AtomicCreatives    *EntityService
	Agencies           *EntityService
	AdServers          *EntityService
	Campaigns          *EntityService
	Concepts           *EntityService
	Creatives          *EntityService
	Deals              *EntityService
	Organizations      *EntityService
	Pixels             *EntityService
	PixelBundles       *EntityService
	PixelProviders     *EntityService
	PlacementSlots     *EntityService
	Publishers         *EntityService
	PublisherSites     *EntityService
	SitePlacements     *EntityService
	SiteLists          *EntityService
	Strategies         *EntityService
	SupplySources      *EntityService
	Users              *EntityService
	Vendors            *EntityService
	VendorContracts    *EntityService
	VendorDomains      *EntityService
	VendorPixels       *EntityService
	VendorPixelDomains *EntityService
	Verticals          *EntityService
}

// NewClient returns a new MediaMath API client. If a nil httpClient is
// provided, an http.Client with sufficient timeout will be used. To use methods
// which require authentication (all methods at this point), provide an
// http.Client that will perform the authentication for you (such as that
// provided by the authenticators/oauth2 or authenticators/cookie libraries).
func NewClient(httpClient *http.Client, apiKey string, baseURL *url.URL) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 300 * time.Second}
	}
	if baseURL == nil {
		baseURL = ProductionURL
	}

	c := &Client{
		client:    httpClient,
		APIKey:    apiKey,
		BaseURL:   baseURL,
		userAgent: generateUserAgentString(),
		rateMu:    sync.Mutex{},
	}

	c.Advertisers = &EntityService{client: c, entityType: "advertisers"}
	c.AtomicCreatives = &EntityService{client: c, entityType: "atomic_creatives"}
	c.Agencies = &EntityService{client: c, entityType: "agencies"}
	c.AdServers = &EntityService{client: c, entityType: "ad_servers"}
	c.Campaigns = &EntityService{client: c, entityType: "campaigns"}
	c.Concepts = &EntityService{client: c, entityType: "concepts"}
	c.Creatives = &EntityService{client: c, entityType: "creatives"}
	c.Deals = &EntityService{client: c, entityType: "deals"}
	c.Organizations = &EntityService{client: c, entityType: "organizations"}
	c.Pixels = &EntityService{client: c, entityType: "pixels"}
	c.PixelBundles = &EntityService{client: c, entityType: "pixel_bundles"}
	c.PixelProviders = &EntityService{client: c, entityType: "pixel_providers"}
	c.PlacementSlots = &EntityService{client: c, entityType: "placement_slots"}
	c.Publishers = &EntityService{client: c, entityType: "publishers"}
	c.PublisherSites = &EntityService{client: c, entityType: "publisher_sites"}
	c.SitePlacements = &EntityService{client: c, entityType: "site_placements"}
	c.SiteLists = &EntityService{client: c, entityType: "site_lists"}
	c.Strategies = &EntityService{client: c, entityType: "strategies"}
	c.SupplySources = &EntityService{client: c, entityType: "supply_sources"}
	c.Users = &EntityService{client: c, entityType: "users"}
	c.Vendors = &EntityService{client: c, entityType: "vendors"}
	c.VendorContracts = &EntityService{client: c, entityType: "vendor_contracts"}
	c.VendorDomains = &EntityService{client: c, entityType: "vendor_domains"}
	c.VendorPixels = &EntityService{client: c, entityType: "vendor_pixels"}
	c.VendorPixelDomains = &EntityService{client: c, entityType: "vendor_pixel_domains"}
	c.Verticals = &EntityService{client: c, entityType: "verticals"}

	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified with a preceding slash. If
// specified, the value pointed to by body is URL-encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body Encoder) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.Reader
	if body != nil {
		buf = strings.NewReader(body.Encode())
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", mediaTypeJSON)
	req.Header.Add("User-Agent", c.userAgent)
	return req, nil
}

// generateUserAgentString generates the user agent for the client to use.
// It is to be called once when generating the client and should not be used
// otherwise.
func generateUserAgentString() string {
	version := strings.Replace(runtime.Version(), "go", "go-http-package/", 1)
	return fmt.Sprintf("go-t1/%d.%d.%d %s", versionMajor, versionMinor,
		versionPatch, version)
}
