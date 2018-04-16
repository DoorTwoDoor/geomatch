/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package handlers provides handler functions for the API endpoints.
package handlers

import (
	"net/http"
	"strings"

	"github.com/doortwodoor/geomatch/utilities"
)

// HandleNotFound handles requests for which no matching route is found.
func HandleNotFound(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	// HTTP request header field names and values.
	const (
		acceptEncodingKey   = "Accept-Encoding"
		acceptEncodingValue = "gzip"
	)
	contentEncoding := request.Header.Get(acceptEncodingKey)
	shouldGzip := strings.Contains(contentEncoding, acceptEncodingValue)

	utilities.WriteErrorResponse(
		responseWriter,
		http.StatusNotFound,
		shouldGzip,
	)
}
