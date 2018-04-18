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
	"net/http"
	"strings"
)

// DecodeQueryParameters decodes query parameters using decoding rules.
func DecodeQueryParameters(
	request *http.Request,
	rules map[string]string,
) (map[string]interface{}, error) {
	const (
		stringToFloat = "stof"
		stringToInt   = "stoi"
	)
	var err error
	encodedQueryParameters := request.URL.Query()
	decodedQueryParameters := make(map[string]interface{})

	for key, value := range rules {
		var number interface{}
		encodedQueryParameter := encodedQueryParameters.Get(key)

		switch value {
		case stringToFloat:
			number, err = ParseFloat(encodedQueryParameter, 64)

		case stringToInt:
			number, err = Atoi(encodedQueryParameter)
		}

		if err != nil {
			break
		}

		decodedQueryParameters[key] = number
	}

	return decodedQueryParameters, err
}

// ShouldGzipResponse returns whether the response should be gzipped.
func ShouldGzipResponse(request *http.Request) bool {
	// HTTP request header field names and values.
	const (
		acceptEncodingKey   = "Accept-Encoding"
		acceptEncodingValue = "gzip"
	)
	contentEncoding := request.Header.Get(acceptEncodingKey)

	return strings.Contains(contentEncoding, acceptEncodingValue)
}
