package t1

// Copyright 2016 MediaMath <http://www.mediamath.com>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime"
)

const (
	// Accept header value to get JSON output
	mediaTypeJSON = "application/vnd.mediamath.v1+json"
)

// Standard base URLs
var (
	Production, _ = url.Parse("https://api.mediamath.com/")
	Sandbox, _    = url.Parse("https://t1sandbox.mediamath.com/")
)

// A Client manages communication with the MediaMath APIs
type Client struct {
	// HTTP client used to make requests. This client should know how to handle authentication
	client *http.Client

	// Bsae URL for API requests. Defaults to the production API endpoint,
	// but can be set to a specific domain for Sandbox or other similar
	// environments. Should just be protocol and domain name with trailing slash
	BaseURL *url.URL

	// User Agent will include library's version number
	userAgent string
}

// generateUserAgentString generates the user agent for the client to use.
// It is to be called once when generating the client and should not be used
// otherwise.
func generateUserAgentString() string {
	return fmt.Sprintf("go-t1/%d.%d.%d %s", versionMajor, versionMinor,
		versionPatch, runtime.Version())
}

// NewClient returns a new MediaMath API client.  If a nil httpClient is
// provided, http.DefaultClient will be used.  To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client, baseURL *url.URL) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		userAgent: generateUserAgentString(),
	}

	// Additional services should be attached here

	return c
}
