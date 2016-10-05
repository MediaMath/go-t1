package t1

import (
	"net/http"
	"reflect"
	"testing"
	"time"
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

// TODO more tests :/

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
