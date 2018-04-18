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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetErrorMessage(t *testing.T) {
	t.Run("Bad Request", func(t *testing.T) {
		t.Parallel()
		const expectedResult = "The resource submitted could not be parsed."

		actualResult := getErrorMessage(http.StatusBadRequest)

		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("Not Found", func(t *testing.T) {
		t.Parallel()
		const expectedResult = "The requested resource could not be found."

		actualResult := getErrorMessage(http.StatusNotFound)

		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("Method Not Allowed", func(t *testing.T) {
		t.Parallel()
		const expectedResult = "The requested method and resource are not compatible."

		actualResult := getErrorMessage(http.StatusMethodNotAllowed)

		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("Unprocessable Entity", func(t *testing.T) {
		t.Parallel()
		const expectedResult = "The fields submitted with this resource are invalid."

		actualResult := getErrorMessage(http.StatusUnprocessableEntity)

		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		t.Parallel()
		const expectedResult = "An unexpected internal error has occurred."

		actualResult := getErrorMessage(http.StatusInternalServerError)

		assert.Equal(t, expectedResult, actualResult)
	})
}
