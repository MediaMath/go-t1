package oauth2

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
	os.Setenv("T1_API_CLIENT_ID", "clientid")
	os.Setenv("T1_API_CLIENT_SECRET", "clientsecret")
	os.Setenv("T1_API_AUIDENCE", "audience")
	os.Setenv("T1_API_GRANTTYPE", "granttype")
	os.Setenv("T1_API_REALM", "realm")
	os.Setenv("T1_API_USERNAME", "username")
	os.Setenv("T1_API_PASSWORD", "password")
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

	if exp := "clientid"; c.ClientID != exp {
		t.Errorf("env client id: want %v, got %v", exp, c.ClientID)
	}
	if exp := "clientsecret"; c.ClientSecret != exp {
		t.Errorf("env client secret: want %v, got %v", exp, c.ClientSecret)
	}
	if exp := "audience"; c.Audience != exp {
		t.Errorf("env audience: want %v, got %v", exp, c.Audience)
	}
	if exp := "granttype"; c.GrantType != exp {
		t.Errorf("env grant: want %v, got %v", exp, c.GrantType)
	}
	if exp := "realm"; c.Realm != exp {
		t.Errorf("env realm: want %v, got %v", exp, c.Realm)
	}
	if exp := "username"; c.Username != exp {
		t.Errorf("env username: want %v, got %v", exp, c.Username)
	}
	if exp := "password"; c.Password != exp {
		t.Errorf("env password: want %v, got %v", exp, c.Password)
	}
}

func TestConfigEncode(t *testing.T) {
	setup()
	vals := GetCredentialsFromEnv().Encode()
	exp := url.Values{
		"client_id":     []string{"clientid"},
		"client_secret": []string{"clientsecret"},
		"audience":      []string{"audience"},
		"grant_type":    []string{"granttype"},
		"realm":         []string{"realm"},
		"username":      []string{"username"},
		"password":      []string{"password"},
	}
	if !reflect.DeepEqual(exp, vals) {
		t.Errorf("config encode: want %v, got %v", exp, vals)
	}
}

func TestNewClient(t *testing.T) {
	conf := Config{}
	c, err := New(&conf, nil, nil)
	if err != nil {
		t.Errorf("new: %v", err)
	}
	if exp := 300 * time.Second; c.Timeout != exp {
		t.Errorf("new timeout: want %v, got %v", exp, c.Timeout)
	}
}

func TestValidLogin(t *testing.T) {
	setup()
	conf := GetCredentialsFromEnv()

	s := setupServer(200, "testdata/valid_login.json", "")
	defer s.Close()

	u, _ := url.Parse(s.URL)
	err := Login(u, &conf)
	if err != nil {
		t.Errorf("valid login: %v", err)
	}
}

func TestAuthError(t *testing.T) {
	setup()
	conf := GetCredentialsFromEnv()

	s := setupServer(401, "testdata/invalid_autherror.json", "")
	defer s.Close()

	u, _ := url.Parse(s.URL)
	err := Login(u, &conf)
	if err == nil {
		t.Error("auth error: expected an error, got none")
    } else if exp, e := "login: error: invalid_grant. Wrong email or password.", err.Error(); e != exp {
		t.Errorf("dev inactive: want %v, got %v", exp, e)
	}
}
