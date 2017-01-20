package cookie

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
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"
)

var (
	prod, _ = url.Parse("https://api.mediamath.com")
)

func setup() {
	os.Setenv("T1_API_USERNAME", "user")
	os.Setenv("T1_API_PASSWORD", "password")
	os.Setenv("T1_API_KEY", "apikey")
}

func setupServer(statusCode int, filename, cType string) *httptest.Server {
	hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if cType == "" {
			w.Header().Set("Content-Type", mediaTypeJSON)
		} else {
			w.Header().Set("Content-Type", cType)
		}
		w.WriteHeader(statusCode)
		io.Copy(w, f)
	})

	return httptest.NewServer(hf)
}

func TestCredentialsFromEnv(t *testing.T) {
	setup()
	c := GetCredentialsFromEnv()
	if exp := "user"; c.Username != exp {
		t.Errorf("env username: want %v, got %v", exp, c.Username)
	}
	if exp := "password"; c.Password != exp {
		t.Errorf("env password: want %v, got %v", exp, c.Password)
	}
	if exp := "apikey"; c.APIKey != exp {
		t.Errorf("env api key: want %v, got %v", exp, c.APIKey)
	}
}

func TestConfigEncode(t *testing.T) {
	setup()
	vals := GetCredentialsFromEnv().Encode()
	exp := url.Values{
		"user":     []string{"user"},
		"password": []string{"password"},
		"api_key":  []string{"apikey"},
	}
	if !reflect.DeepEqual(exp, vals) {
		t.Errorf("config encode: want %v, got %v", exp, vals)
	}
}

func TestNewClient(t *testing.T) {
	c, err := New(Config{}, nil)
	if err != nil {
		t.Errorf("new: %v", err)
	}
	if c.Jar == nil {
		t.Error("new: expected cookie jar, got none attached")
	}
	if exp := 300 * time.Second; c.Timeout != exp {
		t.Errorf("new timeout: want %v, got %v", exp, c.Timeout)
	}
}

func TestSetSession(t *testing.T) {
	c, _ := New(Config{}, nil)
	err := SetSession(c, "mysessionid", prod)
	if err != nil {
		t.Errorf("set session: %v", err)
	}

	cooks := c.Jar.Cookies(prod)
	if len(cooks) == 0 {
		t.Fatal("set session: no cookies set")
	}
	cook := cooks[0]
	if want, got := "adama_session", cook.Name; want != got {
		t.Errorf("cookie name: want %v, got %v", want, got)
	}
	if want, got := "mysessionid", cook.Value; want != got {
		t.Errorf("cookie value: want %v, got %v", want, got)
	}
}

func TestValidLogin(t *testing.T) {
	setup()
	conf := GetCredentialsFromEnv()
	c, _ := New(conf, nil)

	s := setupServer(200, "testdata/valid_login.json", "")
	defer s.Close()

	u, _ := url.Parse(s.URL)
	err := Login(c, u, conf)
	if err != nil {
		t.Errorf("valid login: %v", err)
	}
}

func TestDeveloperInactive(t *testing.T) {
	setup()
	conf := GetCredentialsFromEnv()
	c, _ := New(conf, nil)

	s := setupServer(403, "testdata/invalid_developerinactive.html", "text/xml")
	defer s.Close()

	u, _ := url.Parse(s.URL)
	err := Login(c, u, conf)
	if err == nil {
		t.Error("dev inactive: expected an error, got none")
	} else if exp, e := "login: <h1>Developer Inactive</h1>\n", err.Error(); e != exp {
		t.Errorf("dev inactive: want %v, got %v", exp, e)
	}
}

func TestAuthError(t *testing.T) {
	setup()
	conf := GetCredentialsFromEnv()
	c, _ := New(conf, nil)

	s := setupServer(401, "testdata/invalid_autherror.json", "")
	defer s.Close()

	u, _ := url.Parse(s.URL)
	err := Login(c, u, conf)
	if err == nil {
		t.Error("auth error: expected an error, got none")
	} else if exp, e := "login: Authentication error", err.Error(); e != exp {
		t.Errorf("dev inactive: want %v, got %v", exp, e)
	}
}

func TestGetSessionOK(t *testing.T) {
	setup()
	conf := GetCredentialsFromEnv()
	c, _ := New(conf, nil)

	s := setupServer(200, "test/valid_login.json", "")
	defer s.Close()

	u, _ := url.Parse(s.URL)
	data, err := GetSession(c, u)
	if err != nil {
		t.Errorf("session OK: %v", err)
	}

	if want, got := 1, data.UserID; want != got {
		t.Errorf("user ID: want %v, got %v", want, got)
	}

	if want, got := "user", data.UserName; want != got {
		t.Errorf("user name: want %v, got %v", want, got)
	}

	if want, got := "9d1f77a99dd2b805c425b575028e1a11b314cbe0", data.SessionID; want != got {
		t.Errorf("session ID: want %v, got %v", want, got)
	}

	if want, got := time.Date(2016, time.June, 27, 16, 21, 48, 0, time.UTC), data.ServerTime; !want.Equal(got) {
		t.Errorf("server time: want %s, got %s", want, got)
	}

	if want, got := time.Date(2016, time.June, 28, 16, 21, 48, 0, time.UTC), data.SessionExpires; !want.Equal(got) {
		t.Errorf("server time: want %s, got %s", want, got)
	}
}

func TestGetSessionAuthRequired(t *testing.T) {
	setup()
	conf := GetCredentialsFromEnv()
	c, _ := New(conf, nil)

	s := setupServer(401, "test/auth_required_session.json", "")
	defer s.Close()

	u, _ := url.Parse(s.URL)
	data, err := GetSession(c, u)

	if err == nil {
		t.Error("session auth required: expected an error but got none")
	} else if exp, e := "get session: authentication required", err.Error(); e != exp {
		t.Errorf("session auth required: want %v, got %v", exp, e)
	}

	if data != nil {
		t.Errorf("session auth required: got data instead of nil: %v", data)
	}
}
