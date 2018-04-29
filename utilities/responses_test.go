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
	"net/http/httptest"
	"testing"
	"time"

	"github.com/doortwodoor/geomatch/models"
	"github.com/stretchr/testify/assert"
)

func TestWriteOKResponse(t *testing.T) {
	onlineMover := models.OnlineMover{
		Mover:     "5uls4pSbGeNvQFUYW8X74WraYcx2",
		Latitude:  43.481082,
		Longitude: -80.530143,
		CreatedAt: time.Date(2018, time.April, 8, 10, 0, 0, 0, time.UTC),
	}
	tests := []struct {
		testName   string
		testInputs []interface{}
	}{
		{"gzip", []interface{}{onlineMover, true}},
		{"without_gzip", []interface{}{onlineMover, false}},
	}

	for _, test := range tests {
		test := test

		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			responseWriter := httptest.NewRecorder()
			value := test.testInputs[0].(models.OnlineMover)
			shouldGzip := test.testInputs[1].(bool)
			result := WriteOKResponse(responseWriter, value, shouldGzip)

			assert.Nil(t, result)
		})
	}
}

func TestWriteErrorResponse(t *testing.T) {
	tests := []struct {
		testName   string
		testInputs []interface{}
	}{
		{"gzip", []interface{}{http.StatusBadRequest, true}},
		{"without_gzip", []interface{}{http.StatusNotFound, false}},
	}

	for _, test := range tests {
		test := test

		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			responseWriter := httptest.NewRecorder()
			statusCode := test.testInputs[0].(int)
			shouldGzip := test.testInputs[1].(bool)
			result := WriteErrorResponse(responseWriter, statusCode, shouldGzip)

			assert.Nil(t, result)
		})
	}
}

func TestGetErrorMessage(t *testing.T) {
	tests := []struct {
		testName       string
		testInput      int
		expectedResult string
	}{
		{
			"bad_request",
			http.StatusBadRequest,
			"The resource submitted could not be parsed.",
		},
		{
			"not_found",
			http.StatusNotFound,
			"The requested resource could not be found.",
		},
		{
			"method_not_allowed",
			http.StatusMethodNotAllowed,
			"The requested method and resource are not compatible.",
		},
		{
			"unprocessable_entity",
			http.StatusUnprocessableEntity,
			"The fields submitted with this resource are invalid.",
		},
		{
			"internal_server_error",
			http.StatusInternalServerError,
			"An unexpected internal error has occurred.",
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()
			expectedResult := test.expectedResult
			actualResult := getErrorMessage(test.testInput)

			assert.Equal(t, expectedResult, actualResult)
		})
	}
}

func TestWriteResponseHeader(t *testing.T) {
	responseWriter := httptest.NewRecorder()
	shouldGzip := true
	statusCode := http.StatusOK
	writeResponseHeader(responseWriter, shouldGzip, statusCode)
	response := responseWriter.Result()

	headerFields := map[string]string{
		"Connection":             "keep-alive",
		"Content-Encoding":       "gzip",
		"Content-Type":           "application/json; charset=utf-8",
		"X-Content-Type-Options": "nosniff",
		"Vary": "Accept-Encoding",
	}

	for key, value := range headerFields {
		expectedResult := response.Header.Get(key)
		actualResult := value

		assert.Equal(t, expectedResult, actualResult)
	}

	expectedResult := response.StatusCode
	actualResult := http.StatusOK
	assert.Equal(t, expectedResult, actualResult)
}
