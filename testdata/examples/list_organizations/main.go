package main

import (
	"fmt"
	"github.com/MediaMath/go-t1"
	"github.com/MediaMath/go-t1/authenticators/cookie"
	"github.com/MediaMath/go-t1/models"
	"log"
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
	var orgs []*models.Organization
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
		org = *orgs[0]
		fmt.Printf("org: %#v\n", org)
	}
	t1Client.Organizations.Save(&org)
}
