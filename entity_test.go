package t1

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/MediaMath/go-t1/models"
)

func TestExecuteRequestError(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:234567", nil)
	c := &http.Client{Timeout: 1 * time.Second}
	cl := NewClient(c, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	meta, err := execute(req, cl, nil)
	if err == nil {
		t.Error("Expected err, got none")
	}
	if want := (Meta{}); !reflect.DeepEqual(meta, want) {
		t.Errorf("request error meta: got %v, want %v", meta, want)
	}
}

func TestList(t *testing.T) {
	s := setupServer(200, entityPath+"advertisers", "testdata/fixtures/advertisers.json")
	defer s.Close()
	u, _ := url.Parse(s.URL)
	cl := NewClient(nil, "", u)
	var as []*models.Advertiser
	_, err := cl.Advertisers.List(nil, &as)
	if err != nil {
		t.Errorf("EntityService list: expected no error, got %v", err)
	}
	if got, want := len(as), 100; got != want {
		t.Errorf("EntitySservice list: got %v entities, want %v entities", got, want)
	}
}

func TestGetSingleEntity(t *testing.T) {
	s := setupServer(200, entityPath+"advertisers/1", "testdata/fixtures/advertiser.json")
	defer s.Close()
	u, _ := url.Parse(s.URL)
	cl := NewClient(nil, "", u)
	var a models.Advertiser
	_, err := cl.Advertisers.Get(1, &a)
	if err != nil {
		t.Errorf("EntityService Get: expected no error, got %v", err)
	}
	if got, want := a.Status, true; got != want {
		t.Errorf("Advertiser status: got %v, want %v", got, want)
	}
	if got, want := a.DMPEnabled, "disabled"; got != want {
		t.Errorf("Advertiser dmp_enabled: got %v, want %v", got, want)
	}
	if got, want := a.AgencyID, 300; got != want {
		t.Errorf("Advertiser agency_id: got %v, want %v", got, want)
	}
	cTime := time.Date(2016, 11, 7, 9, 7, 57, 0, time.UTC)
	if got := time.Time(a.CreatedOn); !got.Equal(cTime) {
		t.Errorf("Advertiser created_on: got %v, want %v", got, cTime)
	}
}

// TODO more tests :/

func setupServer(statusCode int, path, filename string) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		w.Header().Set("Content-Type", mediaTypeJSON)
		w.WriteHeader(statusCode)
		io.Copy(w, f)
	})

	return httptest.NewServer(mux)
}

func TestGetIDWithoutIDFieldFalses(t *testing.T) {
	type a struct {
		NotID int
	}
	b := &a{}
	id, valid := getIDOfObject(b)
	if valid {
		t.Errorf("getID with invalid object: want %v, got %v", false, valid)
	}
	if want, got := 0, id; want != got {
		t.Errorf("getID zero field: want %d, got %d", want, got)
	}
}

func TestGetIDWithIDFieldTrue(t *testing.T) {
	type a struct {
		ID int
	}
	b := &a{}
	id, valid := getIDOfObject(b)
	if !valid {
		t.Errorf("getID with valid object: want %v, got %v", true, valid)
	}
	if want, got := 0, id; want != got {
		t.Errorf("getID zero field: want %d, got %d", want, got)
	}
}

func TestGetIDWithIDFieldNonZero(t *testing.T) {
	type a struct {
		ID int
	}
	b := &a{ID: 1}
	id, valid := getIDOfObject(b)
	if !valid {
		t.Errorf("getID with valid object: want %v, got %v", true, valid)
	}
	if want, got := 1, id; want != got {
		t.Errorf("getID non-zero field: want %d, got %d", want, got)
	}
}
