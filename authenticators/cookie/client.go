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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	mediaTypeJSON = "application/vnd.mediamath.v1+json"
)

// Config represents a configuration for cookie authentication. It should
// be directly instantiated with the username, password, and API key for
// authentication.
type Config struct {
	Username string
	Password string
	APIKey   string
}

// New returns an HTTP client configured to use the provided Config and base URL.
// If base is provided, the client will log in automatically and will be usable
// upon return (or will return an auth error). If base is nil, the client will
// not authenticate. This is useful
func New(conf Config, base *url.URL) (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Jar:     jar,
		Timeout: 300 * time.Second,
	}
	if base != nil {
		err = Login(client, base, conf)
		return client, err
	}
	return client, nil
}

// GetCredentialsFromEnv constructs a new Config object from the environment.
// It expecs the following environment variables: T1_API_USERNAME,
// T1_API_PASSWORD, T1_API_KEY
func GetCredentialsFromEnv() Config {
	return Config{
		Username: os.Getenv("T1_API_USERNAME"),
		Password: os.Getenv("T1_API_PASSWORD"),
		APIKey:   os.Getenv("T1_API_KEY"),
	}
}

// Encode constructs a url.Values object from a Config object
func (c Config) Encode() url.Values {
	return url.Values{
		"user":     []string{c.Username},
		"password": []string{c.Password},
		"api_key":  []string{c.APIKey},
	}
}

// Login makes a login request using the supplied HTTP client and Config to
// the supplied base. The HTTP client must have a cookie jar attached, like
// what is provided by New.
func Login(client *http.Client, base *url.URL, conf Config) error {
	body := conf.Encode()
	base.Path = "/api/v2.0/login"
	req, err := http.NewRequest("POST", base.String(), strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Accept", mediaTypeJSON)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return nil
	}
	if cType := resp.Header.Get("Content-Type"); cType != mediaTypeJSON {
		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("login: %v", string(response))
	}
	data := make(map[string]interface{})
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}

	var (
		errs interface{}
		ok   bool
	)
	if errs, ok = data["errors"]; !ok {
		return errors.New("login: unknown error")
	}

	switch ob := errs.(type) {
	case string:
		return fmt.Errorf("login: %v", ob)
	case []interface{}:
		switch obj := ob[0].(type) {
		case map[string]interface{}:
			return fmt.Errorf("login: %v", obj["message"])
		case string:
			return fmt.Errorf("login: %v", obj)
		default:
			return errors.New("login: unknown error")
		}
	default:
		return errors.New("login: unknown error")
	}
}

// SetSession sets an existing adama_session cookie to a given client.
// This is useful for apps where the consumer arrives with a session already
// intact (such as apps).
func SetSession(client *http.Client, sessionID string, baseURL *url.URL) error {
	if client.Jar == nil {
		return errors.New("can't set cookie on nil cookie jar")
	}
	cookie := &http.Cookie{
		Name:     "adama_session",
		Value:    sessionID,
		Path:     "/",
		Domain:   baseURL.Host,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}
	client.Jar.SetCookies(baseURL, []*http.Cookie{cookie})
	return nil
}
