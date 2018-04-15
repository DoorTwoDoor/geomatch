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

// HandlePanic handles panics recovered from HTTP handlers.
func HandlePanic(
	responseWriter http.ResponseWriter,
	request *http.Request,
	value interface{},
) {
	// HTTP response header field names and values.
	const (
		acceptEncodingKey   = "Accept-Encoding"
		acceptEncodingValue = "gzip"
	)
	contentEncoding := request.Header.Get(acceptEncodingKey)
	shouldGzip := strings.Contains(contentEncoding, acceptEncodingValue)

	// Error code and message.
	const (
		code    = 500
		message = "An unexpected internal error has occurred."
	)
	utilities.WriteErrorResponse(responseWriter, code, message, shouldGzip)
}
