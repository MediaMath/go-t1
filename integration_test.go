// +build integration

package t1

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/MediaMath/go-t1/authenticators/cookie"
	"github.com/MediaMath/go-t1/models"
)

var (
	iClient *Client
)

func TestGetSave_org(t *testing.T) {
	var org models.Organization
	meta, err := iClient.Organizations.Get(100048, &org)
	if err != nil {
		t.Fatalf("get org: %v", err)
	}

	if got, want := meta.Status, "ok"; got != want {
		t.Errorf("Get org status: got %v, want %v", got, want)
	}

	v := reflect.ValueOf(org)
	for _, field := range []string{
		"Address1",
		"AdXSeatAccountID",
		"AllowBYOPrice",
		"CreatedOn",
		"ID",
		"OrgType",
		"Version",
	} {
		if isEmptyValue(v.FieldByName(field)) {
			t.Errorf("Get org %v: expected non-empty", field)
		}
	}

	meta, err = iClient.Organizations.Save(&org)
	if err != nil {
		t.Fatalf("Save org: %v", err)
	}
	if got, want := meta.Status, "ok"; got != want {
		t.Errorf("Save org status: got %v, want %v", got, want)
	}
}

func TestMeta_counts(t *testing.T) {
	var advs []models.Advertiser
	meta, err := iClient.Advertisers.List(&UserParams{PageLimit: 10}, &advs)
	if err != nil {
		t.Fatalf("list advs count: %v", err)
	}
	if got, min := meta.TotalCount, 100; got <= min {
		t.Errorf("list advs total count: got %v, want at least %v", got, min)
	}
	if got, want := meta.Count, 10; got != want {
		t.Errorf("list advs count: got %v, want %v", got, want)
	}
	if got, want := len(advs), 10; got != want {
		t.Errorf("list advs actual: got %v, want %v", got, want)
	}
}

func TestMain(m *testing.M) {
	conf := cookie.GetCredentialsFromEnv()
	var err error
	c, err := cookie.New(conf, ProductionURL)
	if err != nil {
		log.Fatalf("cookie login: %v\n", err)
	}
	iClient = NewClient(c, conf.APIKey, ProductionURL)
	os.Exit(m.Run())
}
