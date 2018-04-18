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

	"github.com/doortwodoor/geomatch/models"
)

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
	statusCode int,
	shouldGzip bool,
) error {
	message := getErrorMessage(statusCode)
	errorResponse := models.ErrorResponse{
		Code:    statusCode,
		Message: message,
	}

	writeResponseHeader(responseWriter, shouldGzip, statusCode)

	if shouldGzip {
		gzipWriter := NewGzipWriter(responseWriter)
		defer gzipWriter.Close()

		return Encode(gzipWriter, &errorResponse)
	}

	return Encode(responseWriter, &errorResponse)
}

// getErrorMessage retrieves the error message corresponding to the provided
// status code.
func getErrorMessage(statusCode int) string {
	errorMessage := ""

	switch statusCode {
	case http.StatusBadRequest:
		errorMessage = "The resource submitted could not be parsed."

	case http.StatusNotFound:
		errorMessage = "The requested resource could not be found."

	case http.StatusMethodNotAllowed:
		errorMessage = "The requested method and resource are not compatible."

	case http.StatusUnprocessableEntity:
		errorMessage = "The fields submitted with this resource are invalid."

	case http.StatusInternalServerError:
		errorMessage = "An unexpected internal error has occurred."
	}

	return errorMessage
}

// writeResponseHeader sends a HTTP response header with the provided status
// code.
func writeResponseHeader(
	responseWriter http.ResponseWriter,
	shouldGzip bool,
	statusCode int,
) {
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

	if shouldGzip {
		responseWriter.Header().Set(contentEncodingKey, contentEncodingValue)
		responseWriter.Header().Set(varyKey, varyValue)
	}

	responseWriter.Header().Set(connectionKey, connectionValue)
	responseWriter.Header().Set(contentTypeKey, contentTypeValue)
	responseWriter.Header().Set(xContentTypeOptionsKey, xContentTypeOptionsValue)
	responseWriter.WriteHeader(statusCode)
}
