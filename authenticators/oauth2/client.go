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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	mediaTypeJSON = "application/vnd.mediamath.v1+json"
)

// Config represents a configuration for oauth2 authentication.
type Config struct {
	ClientID     string
	ClientSecret string
	Audience     string
	GrantType    string
	Realm        string
	Username     string
	Password     string
	AccessToken  string
}

// New returns an HTTP client configured to use the provided Config and base URL.
// If base is provided, the client will log in automatically and will be usable
// upon return (or will return an auth error). If base is nil, the client will
// not authenticate.
func New(conf *Config, base *url.URL, auth *url.URL) (*http.Client, error) {
	client := &http.Client{
		Timeout: 300 * time.Second,
	}
	if base != nil {
		err := Login(auth, conf)
		return client, err
	}
	return client, nil
}

// GetCredentialsFromEnv constructs a new Config object from the environment.
// It expecs the following environment variables: T1_API_USERNAME,
// T1_API_PASSWORD, T1_API_KEY
func GetCredentialsFromEnv() Config {
	return Config{
		ClientID:     os.Getenv("T1_API_CLIENT_ID"),
		ClientSecret: os.Getenv("T1_API_CLIENT_SECRET"),
		Audience:     os.Getenv("T1_API_AUIDENCE"),
		GrantType:    os.Getenv("T1_API_GRANTTYPE"),
		Realm:        os.Getenv("T1_API_REALM"),
		Username:     os.Getenv("T1_API_USERNAME"),
		Password:     os.Getenv("T1_API_PASSWORD"),
	}
}

// Encode constructs a url.Values object from a Config object
func (c Config) Encode() url.Values {
	return url.Values{
		"client_id":     []string{c.ClientID},
		"client_secret": []string{c.ClientSecret},
		"audience":      []string{c.Audience},
		"grant_type":    []string{c.GrantType},
		"realm":         []string{c.Realm},
		"username":      []string{c.Username},
		"password":      []string{c.Password},
	}
}

// Login makes a login request using the supplied HTTP client and Config to
// the supplied base.
func Login(base *url.URL, conf *Config) error {
	client := &http.Client{
		Timeout: 300 * time.Second,
	}
	body := conf.Encode()
	base.Path = "oauth/token"
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

	if cType := resp.Header.Get("Content-Type"); cType != "application/vnd.mediamath.v1+json" {
		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("login: unable to read response body: %v", err)
		}
        return fmt.Errorf("login: unexpected Content-Type: %v, body: %v", cType,  string(response))
	}

	data := make(map[string]interface{})
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return fmt.Errorf("login: error decoding response body: %v", err)
	}

	if resp.StatusCode == 200 {
		if accessToken, ok := data["access_token"]; ok {
			conf.AccessToken = accessToken.(string)
			return nil
		}
		return errors.New("login: unexpected successful login result, cannot read access_token")
	}

	if err, ok := data["error"]; ok {
		errDescrip, ok := data["error_description"]
		if ok {
			err = err.(string) + ". " + errDescrip.(string)
		}
		return errors.New("login: error: " + err.(string))
	}

	return errors.New("unexpected login result")
}
