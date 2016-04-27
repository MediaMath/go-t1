/*
Package t1 is a go (golang) library for MediaMath's APIs.

Construct a new T1 client, then use the various services on the client to
access different parts of the MediaMath API modules.

For example:

	package main

	import (
		"github.com/MediaMath/go-t1"
		"github.com/MediaMath/go-t1/authenticators/cookie"
	)

	auth := cookieauth.New("myusername", "mypassword", "myapikey")
	service := t1.New(auth)
	orgs, err := client.Organizations.List(nil)

For API documentation, please visit https://developer.mediamath.com
*/
package t1
