package t1

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	entityPath = "/api/v2.0/"
)

// EntityService is a generalized service object that helps work with entities.
// Designed to be instantiated for each entity type.
type EntityService struct {
	client     *Client
	entityType string
}

// EntityResponse is the response returned by the Execution and Management API.
type EntityResponse struct {
	Meta Meta
	Data json.RawMessage
}

// Get fetches an entity by ID and loads it into the data object passed in.
// The object should be provided so that EntityService does not need to do
// the job of figuring out what kind of entity it is. Like with encoding/json,
// it is the user's responsiblity to provide the correct type of object.
func (s *EntityService) Get(id int, data interface{}) (Meta, error) {
	buf := bufferPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		bufferPool.Put(buf)
	}()

	buf.WriteString(entityPath)
	buf.WriteString(s.entityType)
	buf.WriteByte('/')
	buf.WriteString(strconv.Itoa(id))

	vals := valuesPool.Get().(url.Values)
	defer func() {
		vals.Del("api_key")
		valuesPool.Put(vals)
	}()

	// The only user param that matters when retrieving a single entity
	// is with. If with isn't supported at all then this block isn't needed.
	// If we support including with-relations, then this block should be
	// put back in, and the method signature should change to:
	// func (s *EntityService) Get(id int, params *UserParams, data interface{}) (Meta, error) {
	// if params != nil && len(params.With) > 0 {
	// 	structToMapGivenValues(params, vals)
	// }

	vals.Set("api_key", s.client.APIKey)

	req, err := s.client.NewRequest("GET", buf.String(), vals)
	if err != nil {
		return Meta{}, err
	}

	return execute(req, s.client, data)
}

func execute(req *http.Request, c *Client, data interface{}) (Meta, error) {
	r, err := c.Do(req, nil)
	if err != nil {
		return Meta{}, err
	}
	defer func() {
		// Let the Transport reuse the connection
		// cf. https://github.com/google/go-github/pull/317
		io.CopyN(ioutil.Discard, r.Body, 512)
		r.Body.Close()
	}()

	var resp EntityResponse
	err = json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		return resp.Meta, err
	}

	err = json.Unmarshal(resp.Data, data)
	if err != nil && err != io.EOF {
		return resp.Meta, err
	}

	return resp.Meta, nil
}

// List fetches a list of entities according to the given user params and loads
// it into the data object passed in. data should be a slice of whatever entity
// the service represents.
func (s *EntityService) List(params *UserParams, data interface{}) (Meta, error) {
	buf, vals := bufferPool.Get().(*bytes.Buffer), valuesPool.Get().(url.Values)
	defer func() {
		buf.Reset()
		bufferPool.Put(buf)
		for key := range vals {
			vals.Del(key)
		}
		valuesPool.Put(vals)
	}()

	buf.WriteString(entityPath)
	buf.WriteString(s.entityType)
	if params != nil {
		structToMapGivenValues(params, vals)
	}

	vals.Set("api_key", s.client.APIKey)
	buf.WriteByte('?')
	buf.WriteString(vals.Encode())
	req, err := s.client.NewRequest("GET", buf.String(), nil)
	if err != nil {
		return Meta{}, err
	}

	return execute(req, s.client, data)
}
