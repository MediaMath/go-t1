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
	"github.com/MediaMath/go-t1/time"
)

// Meta is a struct of the metadata returned by some of the APIs.
type Meta struct {
	CalledOn   t1time.T1Time `json:"called_on"`
	Count      int           `json:"count"`
	ETag       string        `json:"etag"`
	NextPage   string        `json:"next_page"`
	Offset     int           `json:"offset"`
	Status     string        `json:"status"`
	TotalCount int           `json:"total_count"`
}
