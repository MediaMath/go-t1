package t1

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func TestNotFound(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("testdata/error_not_found.json")
		if err != nil {
			t.Fatalf("not found open fixture: %v", err)
		}
		defer f.Close()
		w.Header().Add("Content-Type", mediaTypeJSON)
		w.WriteHeader(404)
		io.Copy(w, f)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(req, nil)

	if err == nil {
		t.Fatal("Expected error to be returned.")
	}
	er, ok := err.(*ErrorResponse)
	if !ok {
		t.Fatalf("Not found error type: got %T, want *ErrorResponse. Error: %v", err, err)
	}
	if got, want := er.Meta.Status, "not_found"; got != want {
		t.Errorf("Meta status: got %v, want %v", got, want)
	}
	if er.Errors == nil || len(er.Errors) != 1 {
		t.Fatalf("Expected one error, got %v", er.Errors)
	}
	if got, want := er.Message, "Not found"; got != want {
		t.Errorf("Error message: got %v, want %v", got, want)
	}
}

func TestBadRequest_fieldError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("testdata/error_bad_request_field_error.json")
		if err != nil {
			t.Fatalf("not found open fixture: %v", err)
		}
		defer f.Close()
		w.Header().Add("Content-Type", mediaTypeJSON)
		w.WriteHeader(400)
		io.Copy(w, f)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(req, nil)

	if err == nil {
		t.Fatal("Expected error to be returned.")
	}
	er, ok := err.(*ErrorResponse)
	if !ok {
		t.Fatalf("Field error type: got %T, want *ErrorResponse. Error: %v", err, err)
	}
	if got, want := er.Meta.Status, "invalid"; got != want {
		t.Errorf("Meta status: got %v, want %v", got, want)
	}
	if er.Errors == nil || len(er.Errors) != 1 {
		t.Fatalf("Expected one error, got %v", er.Errors)
	}
	if got, want := er.Message, "agency_id is not writable"; got != want {
		t.Errorf("Error message: got %v, want %v", got, want)
	}
	if got, want := er.Errors[0].Type, "field-error"; got != want {
		t.Errorf("Error type: got %v, want %v", got, want)
	}
	if got, want := er.Errors[0].Field, "agency_id"; got != want {
		t.Errorf("Error field: got %v, want %v", got, want)
	}
}

// TODO test conflict, multiple validation error, access denied
