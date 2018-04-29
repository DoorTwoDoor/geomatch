/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package utilities provides functions to work with JSON codec, parse requests
// write responses, validate, perform Cloud Datastore operations and perform
// Redis operations.
package utilities

import (
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeQueryParameters(t *testing.T) {
	expectedResult := map[string]interface{}{
		"latitude":  43.474307,
		"longitude": -80.537230,
		"radius":    1000.0,
		"limit":     10,
	}

	request := httptest.NewRequest("GET", "/nearest-movers", nil)
	queryParameters := request.URL.Query()
	latitude := strconv.FormatFloat(43.474307, 'E', -1, 64)
	longitude := strconv.FormatFloat(-80.537230, 'E', -1, 64)
	radius := strconv.FormatFloat(1000, 'E', -1, 64)
	limit := strconv.FormatInt(10, 10)
	queryParameters.Add("latitude", latitude)
	queryParameters.Add("longitude", longitude)
	queryParameters.Add("radius", radius)
	queryParameters.Add("limit", limit)
	request.URL.RawQuery = queryParameters.Encode()

	decodingRules := map[string]string{
		"latitude":  "stof",
		"longitude": "stof",
		"radius":    "stof",
		"limit":     "stoi",
	}

	actualResult, _ := DecodeQueryParameters(request, decodingRules)

	assert.Equal(t, expectedResult, actualResult)
}

func TestShouldGzipResponse(t *testing.T) {
	expectedResult := true

	request := httptest.NewRequest("GET", "/nearest-movers", nil)
	request.Header.Set("Accept-Encoding", "application/gzip")
	request.Header.Set("Content-Type", "application/json")

	actualResult := ShouldGzipResponse(request)

	assert.Equal(t, expectedResult, actualResult)
}
