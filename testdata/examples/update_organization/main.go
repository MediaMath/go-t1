package main

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

	// Model object gets passed in to the various service methods
	var org models.Organization

	meta, err := t1Client.Organizations.Get(100048, &org)
	if err != nil {
		log.Fatalf("get org error: %v", err)
	}
	fmt.Printf("Meta:\t%#v\nOrg:\t%#v\n", meta, org)

	oName := org.Name

	org.Name = oName + " updated"
	meta, err = t1Client.Organizations.Save(&org)
	if err != nil {
		log.Fatalf("update org error: %v", err)
	}
	fmt.Printf("Meta:\t%#v\nNew Name:\t%#v\n", meta, org.Name)

	org.Name = oName
	meta, err = t1Client.Organizations.Save(&org)
	if err != nil {
		log.Fatalf("update org error: %v", err)
	}
	fmt.Printf("Meta:\t%#v\nNew Name:\t%#v\n", meta, org.Name)
}
