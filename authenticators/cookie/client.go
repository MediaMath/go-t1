package cookie

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"encoding/json"
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

type Config struct {
	Username string
	Password string
	APIKey   string
}

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

func GetCredentialsFromEnv() Config {
	return Config{
		Username: os.Getenv("T1_API_USERNAME"),
		Password: os.Getenv("T1_API_PASSWORD"),
		APIKey:   os.Getenv("T1_API_KEY"),
	}
}

func (c Config) Encode() url.Values {
	return url.Values{
		"user":     []string{c.Username},
		"password": []string{c.Password},
		"api_key":  []string{c.APIKey},
	}
}

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
	if val, ok := data["errors"]; ok {
		switch val.(type) {
		case []interface{}:
			ob1 := val.([]interface{})[0].(map[string]interface{})
			return fmt.Errorf("login: %v", ob1["message"])
		}
	}
	return fmt.Errorf("login: unknown error")
}
