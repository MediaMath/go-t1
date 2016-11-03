package main

import (
	"fmt"
	"github.com/MediaMath/go-t1"
	"github.com/MediaMath/go-t1/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/mediamath"
	"log"
	"os"
	"time"
)

func main() {
	// Set up configuration from envvars
	conf := oauth2.Config{
		ClientID:     os.Getenv("T1_API_CLIENT_ID"),
		ClientSecret: os.Getenv("T1_API_CLIENT_SECRET"),
		Endpoint:     mediamath.Endpoint,
		RedirectURL:  "https://www.mediamath.com/",
	}

	// Redirect user to consent page to ask for permission
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v\n", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	fmt.Print("Paste code here: ")
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}

	c := conf.Client(oauth2.NoContext, tok)

	// Construct new t1 client
	t1Client := t1.NewClient(c, conf.ClientID, t1.ProductionURL)

	// Model object gets passed in to the various service methods
	var org models.Organization

	meta, err := t1Client.Organizations.Get(100048, &org)
	if err != nil {
		log.Fatalf("get org error: %v", err)
	}
	fmt.Printf("Meta:\t%#v\nOrg:\t%#v\n", meta, org)

	fmt.Printf("Org created time: %v\n", time.Time(org.CreatedOn))
}
