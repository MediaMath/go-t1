package t1

// Copyright 2016-2017 MediaMath
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
		f, err := os.Open("testdata/fixtures/error_not_found.json")
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
		f, err := os.Open("testdata/fixtures/error_bad_request_field_error.json")
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
