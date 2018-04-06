/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package utilities provides functions to work with JSON codec and write
// responses.
package utilities

import (
	"net/http"

	"github.com/doortwodoor/geomatch/models"
)

// HTTP response header field names and values.
const (
	connectionKey            = "Connection"
	connectionValue          = "keep-alive"
	contentEncodingKey       = "Content-Encoding"
	contentEncodingValue     = "gzip"
	contentTypeKey           = "Content-Type"
	contentTypeValue         = "application/json; charset=utf-8"
	xContentTypeOptionsKey   = "X-Content-Type-Options"
	xContentTypeOptionsValue = "nosniff"
	varyKey                  = "Vary"
	varyValue                = "Accept-Encoding"
)

// writeResponseHeader sends a HTTP response header with the provided status
// code.
func writeResponseHeader(
	responseWriter http.ResponseWriter,
	shouldGzip bool,
	statusCode int,
) {
	if shouldGzip {
		responseWriter.Header().Set(contentEncodingKey, contentEncodingValue)
		responseWriter.Header().Set(varyKey, varyValue)
	}

	responseWriter.Header().Set(connectionKey, connectionValue)
	responseWriter.Header().Set(contentTypeKey, contentTypeValue)
	responseWriter.Header().Set(xContentTypeOptionsKey, xContentTypeOptionsValue)
	responseWriter.WriteHeader(statusCode)
}

// WriteOKResponse writes the response as a standard JSON response with status
// OK.
func WriteOKResponse(
	responseWriter http.ResponseWriter,
	value interface{},
	shouldGzip bool,
) error {
	writeResponseHeader(responseWriter, shouldGzip, http.StatusOK)

	if shouldGzip {
		gzipWriter := NewGzipWriter(responseWriter)
		defer gzipWriter.Close()

		return Encode(gzipWriter, &value)
	}

	return Encode(responseWriter, &value)
}

// WriteErrorResponse writes the response as a standard JSON response with
// status code.
func WriteErrorResponse(
	responseWriter http.ResponseWriter,
	code int,
	message string,
	shouldGzip bool,
) error {
	errorResponse := models.ErrorResponse{
		Code:    code,
		Message: message,
	}

	writeResponseHeader(responseWriter, shouldGzip, code)

	if shouldGzip {
		gzipWriter := NewGzipWriter(responseWriter)
		defer gzipWriter.Close()

		return Encode(gzipWriter, &errorResponse)
	}

	return Encode(responseWriter, &errorResponse)
}
