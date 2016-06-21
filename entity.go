package t1

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	entityPath = "/api/v2.0/"
)

type EntityService struct {
	client     *Client
	entityType string
}

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

	// The only user param that matters when retrieving a single entity
	// is with. If with isn't supported at all then this block isn't needed.
	// If we support including with-relations, then this block should be
	// put back in, and the method signature should change to:
	// func (s *EntityService) Get(id int, params *UserParams, data interface{}) (Meta, error) {
	// if params != nil && len(params.With) > 0 {
	// 	buf.WriteByte('?')
	// 	buf.WriteString(params.Encode())
	// }

	req, err := s.client.NewRequest("GET", buf.String(), nil)
	if err != nil {
		return Meta{}, err
	}

	return execute(req, s.client.client, data)
}

func execute(req *http.Request, c *http.Client, data interface{}) (Meta, error) {
	r, err := c.Do(req)
	if err != nil {
		return Meta{}, err
	}
	defer r.Body.Close()

	var resp EntityResponse
	err = json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		return resp.Meta, err
	}

	err = json.Unmarshal(resp.Data, data)
	if err != nil {
		return resp.Meta, err
	}

	return resp.Meta, nil
}

// List fetches a list of entities according to the given user params and loads
// it into the data object passed in. data should be a slice of whatever entity
// the service represents.
func (s *EntityService) List(params *UserParams, data interface{}) (Meta, error) {
	buf := bufferPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		bufferPool.Put(buf)
	}()

	buf.WriteString(entityPath)
	buf.WriteString(s.entityType)
	if params != nil {
		buf.WriteByte('?')
		buf.WriteString(params.Encode())
	}
	req, err := s.client.NewRequest("GET", buf.String(), nil)
	if err != nil {
		return Meta{}, err
	}

	return execute(req, s.client.client, data)
}
