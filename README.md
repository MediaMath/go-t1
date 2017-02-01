go-t1
=====

go-t1 is a Go client for MediaMath's APIs. API Documentation is available [on the developer portal](https://developer.mediamath.com/docs/TerminalOne_API_Overview).

Reference: [![GoDoc](https://godoc.org/github.com/MediaMath/go-t1?status.svg)](https://godoc.org/github.com/MediaMath/go-t1)

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)

## Installation

```bash
$ go get github.com/MediaMath/go-t1
```

## Usage

[embedmd]:# (testdata/examples/get_organization/main.go /import/ /\)/)
```go
import (
	"fmt"
	"github.com/MediaMath/go-t1"
	"github.com/MediaMath/go-t1/authenticators/cookie"
	"github.com/MediaMath/go-t1/models"
	"log"
	"time"
)
```

To set up authentication, use an authenticator:

[embedmd]:# (testdata/examples/get_organization/main.go /.*Set up config/ /log\.Fatalf[^}]*\}/)
```go
	// Set up configuration from envvars
	conf := cookie.GetCredentialsFromEnv()

	// Create new *http.Client with these credentials
	c, err := cookie.New(conf, t1.ProductionURL)
	if err != nil {
		log.Fatalf("initial login: %v", err)
	}
```

The authenticators are just `*http.Client` objects that know how to authenticate. Cookie is provided in the package `authenticators/cookie`, and OAuth2 is supported. To use OAuth2, use Google's [OAuth2 package](https://github.com/golang/oauth2), with a [MediaMath endpoint](https://github.com/golang/oauth2/blob/master/mediamath/mediamath.go). An example is provided in the `testdata/examples` directory.

Construct a new client, then use the various services on the client to access different parts of the MediaMath API.

[embedmd]:# (testdata/examples/get_organization/main.go /.*Construct new t1/ /fmt.Printf\("Meta:.*\)/)
```go
	// Construct new t1 client
	t1Client := t1.NewClient(c, conf.APIKey, t1.ProductionURL)

	// Model object gets passed in to the various service methods
	var org models.Organization

	meta, err := t1Client.Organizations.Get(100048, &org)
	if err != nil {
		log.Fatalf("get org error: %v", err)
	}
	fmt.Printf("Meta:\t%#v\nOrg:\t%#v\n", meta, org)
```

This whole example is available in the `testdata/examples/get_organization` directory. The `testdata/examples` directory also has examples of listing, creating, and updating entities.

## Time Types

Execution and Management API currently returns times in a format conforming to ISO 8601 but not RFC 3339. As such, there is a time package `t1time` that provides a time type compatible with this. This is a `time.Time` type, so can be converted easily:

[embedmd]:# (testdata/examples/get_organization/main.go /time\.[^)]*\)/)
```go
time.Time(org.CreatedOn)
```
