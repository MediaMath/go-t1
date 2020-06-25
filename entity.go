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
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
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
	buf, vals := bufferPool.Get().(*bytes.Buffer), valuesPool.Get().(url.Values)
	defer func() {
		putBufferBackInPool(buf)
		putValuesBackInPool(vals)
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
	// 	structToMapGivenValues(params, vals)
	// }

	vals.Set("api_key", s.client.APIKey)
	buf.WriteByte('?')
	buf.WriteString(vals.Encode())

	req, err := s.client.NewRequest("GET", buf.String(), nil)
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
		putBufferBackInPool(buf)
		putValuesBackInPool(vals)
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

// Save posts an entity to the API. data *must* be a pointer to an object: save
// will modify the object with what gets returned.
func (s *EntityService) Save(data interface{}) (Meta, error) {
	if data == nil {
		return Meta{}, errors.New("save: nil data provided")
	}

	buf, vals := bufferPool.Get().(*bytes.Buffer), valuesPool.Get().(url.Values)
	defer func() {
		putBufferBackInPool(buf)
		putValuesBackInPool(vals)
	}()

	buf.WriteString(entityPath)
	buf.WriteString(s.entityType)
	if id, _ := getIDOfObject(data); id != 0 {
		buf.WriteByte('/')
		buf.WriteString(strconv.Itoa(id))
	}
	buf.WriteString("?api_key=")
	buf.WriteString(s.client.APIKey)

	structToMapGivenValues(data, vals)

	req, err := s.client.NewRequest("POST", buf.String(), vals)
	if err != nil {
		return Meta{}, err
	}

	return execute(req, s.client, data)
}

// SaveWithResponse posts an entity to the API. data *must* be a pointer to an object: save
// will modify the response struct with what gets returned.
func (s *EntityService) SaveWithResponse(data interface{}, response interface{}) (Meta, error) {
	if data == nil {
		return Meta{}, errors.New("save: nil data provided")
	}

	buf, vals := bufferPool.Get().(*bytes.Buffer), valuesPool.Get().(url.Values)
	defer func() {
		putBufferBackInPool(buf)
		putValuesBackInPool(vals)
	}()

	buf.WriteString(entityPath)
	buf.WriteString(s.entityType)
	if id, _ := getIDOfObject(data); id != 0 {
		buf.WriteByte('/')
		buf.WriteString(strconv.Itoa(id))
	}
	buf.WriteString("?api_key=")
	buf.WriteString(s.client.APIKey)

	structToMapGivenValues(data, vals)

	req, err := s.client.NewRequest("POST", buf.String(), vals)
	if err != nil {
		return Meta{}, err
	}

	return execute(req, s.client, response)
}

func getIDOfObject(data interface{}) (int, bool) {
	p := reflect.ValueOf(data)
	f := reflect.Indirect(p).FieldByName("ID")
	if !f.IsValid() {
		return 0, false
	}
	return int(f.Int()), true
}
