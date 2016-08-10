package main

import (
	"fmt"
	"github.com/MediaMath/go-t1"
	"github.com/MediaMath/go-t1/authenticators/cookie"
	"github.com/MediaMath/go-t1/models"
	"log"
	"time"
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
	var org models.Organization

	meta, err := t1Client.Organizations.Get(100048, &org)
	if err != nil {
		log.Fatalf("get org error: %v", err)
	}
	fmt.Printf("Meta:\t%#v\nOrg:\t%#v\n", meta, org)

	fmt.Printf("Org created time: %v\n", time.Time(org.CreatedOn))
}
