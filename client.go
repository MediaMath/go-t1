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
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	mediaTypeJSON       = "application/vnd.mediamath.v1+json" // Accept for JSON
	mediaTypeURLEncoded = "application/x-www-form-urlencoded" // POST Content-Type
)

// Standard base URLs
var (
	T1ProductionURL, _    = url.Parse("https://api.mediamath.com")
	Auth0ProductionURL, _ = url.Parse("https://auth.mediamath.com")
	Auth0DevURL, _        = url.Parse("https://mediamath-dev.auth0.com")
)

var (
	nullTime = time.Time{}
)

// A Client manages communication with the MediaMath APIs
type Client struct {
	// HTTP client used to make requests. This client should know how to handle authentication
	client *http.Client

	// Access token to be used. This will be set as the Authorization
	// bearer token in all requests made.
	AccessToken string

	// Base URL for API requests. Defaults to the production API endpoint,
	// but can be set to a specific domain for Sandbox or other similar
	// environments. Should be protocol and domain name without trailing slash
	BaseURL        *url.URL
	rateMu         sync.Mutex
	RateLimitReset time.Time

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
// provided by the authenticators/oauth2 library).
func NewClient(httpClient *http.Client, accessToken string, baseURL *url.URL) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 300 * time.Second}
	}
	if baseURL == nil {
		baseURL = T1ProductionURL
	}

	c := &Client{
		client:      httpClient,
		AccessToken: accessToken,
		BaseURL:     baseURL,
		userAgent:   generateUserAgentString(),
		rateMu:      sync.Mutex{},
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

	var req *http.Request
	if body != nil {
		buf := strings.NewReader(body.Encode())
		req, err = http.NewRequest(method, u.String(), buf)
		req.Header.Add("Content-Type", mediaTypeURLEncoded)
	} else {
		req, err = http.NewRequest(method, u.String(), nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", mediaTypeJSON)
	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("Authorization", "Bearer "+c.AccessToken)
	return req, nil
}

// parseRateLimit parses the rate limit headers when a request comes
// back over the rate limit
func parseRateLimit(r *http.Response) (t time.Time) {

	date, retryStr := r.Header.Get("Date"), r.Header.Get("Retry-After")
	if date == "" || retryStr == "" {
		return
	}

	t, err := time.Parse(time.RFC1123, date)
	if err != nil {
		return
	}

	retry, err := strconv.Atoi(retryStr)
	if err != nil {
		return
	}
	t = t.Add(time.Duration(retry) * time.Second)
	return
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it. If rate limit is exceeded and reset time is in the future,
// Do returns *RateLimitError immediately without making a network API call.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	// If we've hit rate limit, don't make further requests before Reset time.
	if err := c.checkRateLimitBeforeDo(req); err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		if e, ok := err.(*RateLimitError); ok {
			c.rateMu.Lock()
			c.RateLimitReset = e.RetryAt
			c.rateMu.Unlock()
		}
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return resp, err
	}

	if v != nil {
		defer func() {
			// Let the Transport reuse the connection
			// cf. https://github.com/google/go-github/pull/317
			io.CopyN(ioutil.Discard, resp.Body, 512)
			resp.Body.Close()
		}()
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return resp, err
}

// checkRateLimitBeforeDo does not make any network calls, but uses existing knowledge from
// current client state in order to quickly check if *RateLimitError can be immediately returned
// from Client.Do, and if so, returns it so that Client.Do can skip making a network API call unneccessarily.
// Otherwise it returns nil, and Client.Do should proceed normally.
func (c *Client) checkRateLimitBeforeDo(req *http.Request) error {
	c.rateMu.Lock()
	rate := c.RateLimitReset
	c.rateMu.Unlock()
	if !rate.IsZero() && time.Now().Before(rate) {
		// Create a fake response.
		resp := &http.Response{
			Status:     http.StatusText(http.StatusForbidden),
			StatusCode: http.StatusForbidden,
			Request:    req,
			Header:     make(http.Header),
			Body:       ioutil.NopCloser(strings.NewReader("")),
		}
		return &RateLimitError{
			RetryAt:  rate,
			Response: resp,
			Message:  fmt.Sprintf("API rate limit still exceeded until %v, not making remote request.", rate),
		}
	} else if !rate.IsZero() {
		c.rateMu.Lock()
		c.RateLimitReset = nullTime
		c.rateMu.Unlock()
	}

	return nil
}

// generateUserAgentString generates the user agent for the client to use.
// It is to be called once when generating the client and should not be used
// otherwise.
func generateUserAgentString() string {
	version := strings.Replace(runtime.Version(), "go", "go-http-package/", 1)
	return fmt.Sprintf("go-t1/%d.%d.%d %s", versionMajor, versionMinor,
		versionPatch, version)
}

// Session retrieves the session information of the client.
func (c *Client) Session() (Session, error) {
	var s Session
	u := entityPath + "session"
	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return s, err
	}

	var resp EntityResponse
	_, err = c.Do(req, &resp)
	if err != nil {
		return s, err
	}

	if resp.Data == nil {
		return s, fmt.Errorf("Session: %v", resp.Meta.Status)
	}

	err = json.Unmarshal(resp.Data, &s)
	if err != nil && err != io.EOF {
		return s, err
	}

	return s, nil
}
